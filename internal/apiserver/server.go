package apiserver

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TobyIcetea/fastgo/internal/apiserver/biz"
	"github.com/TobyIcetea/fastgo/internal/apiserver/handler"
	"github.com/TobyIcetea/fastgo/internal/apiserver/pkg/validation"
	"github.com/TobyIcetea/fastgo/internal/apiserver/store"
	"github.com/TobyIcetea/fastgo/internal/pkg/core"
	"github.com/TobyIcetea/fastgo/internal/pkg/errorsx"
	"github.com/TobyIcetea/fastgo/internal/pkg/known"
	mw "github.com/TobyIcetea/fastgo/internal/pkg/middleware"
	genericoptions "github.com/TobyIcetea/fastgo/pkg/options"
	"github.com/TobyIcetea/fastgo/pkg/token"
	"github.com/gin-gonic/gin"
)

// Config 配置结构体，用于存储应用相关的配置
// 不用 viper.Get，是因为这种方式能更加清晰的知道应用提供了哪些配置项
type Config struct {
	MySQLOptions *genericoptions.MySQLOptions
	Addr         string
	JWTKey       string
	Expiration   time.Duration
}

// Server 定义了一个服务器结构体类型
type Server struct {
	cfg *Config
	srv *http.Server
}

// NewServer 根据配置创建服务器
func (cfg *Config) NewServer() (*Server, error) {
	// 初始化 token 包的签名密钥、认证 Key 及 Token 默认过期时间
	token.Init(cfg.JWTKey, known.XUserID, cfg.Expiration)

	// 创建 Gin 引擎
	engine := gin.New()

	// gin.Recovery() 中间件，用来捕获任何 panic，并恢复
	mws := []gin.HandlerFunc{gin.Recovery(), mw.NoCache, mw.Cors, mw.RequestID()}
	engine.Use(mws...)

	// 初始化数据库连接
	db, err := cfg.MySQLOptions.NewDB()
	if err != nil {
		return nil, err
	}
	store := store.NewStore(db)

	cfg.InstallRESTAPI(engine, store)

	// 创建 HTTP Server 实例
	httpsrv := &http.Server{Addr: cfg.Addr, Handler: engine}

	return &Server{cfg: cfg, srv: httpsrv}, nil

}

// 注册 API 路由。路由的路径和 HTTP 方法，严格遵循 REST 规范
func (cfg *Config) InstallRESTAPI(engine *gin.Engine, store store.IStore) {
	// 注册 404 Handler
	engine.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errorsx.ErrNotFound.WithMessage("Page not found"), nil)
	})

	// 注册 /healthz handler.
	engine.GET("/healthz", func(c *gin.Context) {
		core.WriteResponse(c, map[string]string{"Status": "ok"}, nil)
	})

	// 创建核心业务处理器
	handler := handler.NewHandler(biz.NewBiz(store), validation.NewValidator(store))

	// 注册用户登录和令牌刷新接口。这2个接口比较简单，所以没有 API 版本
	engine.POST("/login", handler.Login)
	// 注意：认证中间件要在 handler.RefreshToken 之前加载
	engine.PUT("/refresh-token", mw.Authn(), handler.RefreshToken)

	authMiddlewares := []gin.HandlerFunc{mw.Authn()}

	// 注册 v1 版本 API 路由分组
	v1 := engine.Group("/v1")
	{
		// 用户相关路由
		userv1 := v1.Group("/users")
		{
			// 创建用户。这里要注意：创建用户是不用进行认证和授权的
			userv1.POST("", handler.CreateUser) // 创建用户
			userv1.Use(authMiddlewares...)
			userv1.PUT(":userID/change-password", handler.ChangePassword) // 修改用户密码
			userv1.PUT(":userID", handler.UpdateUser)                     // 更新用户信息
			userv1.DELETE(":userID", handler.DeleteUser)                  // 删除用户
			userv1.GET(":userID", handler.GetUser)                        // 查询用户详情
			userv1.GET("", handler.ListUser)                              // 查询用户列表.
		}

		// 博客相关路由
		postv1 := v1.Group("/posts", authMiddlewares...)
		{
			// 创建博客
			postv1.POST("", handler.CreatePost)       // 创建博客
			postv1.PUT(":postID", handler.UpdatePost) // 更新博客
			postv1.DELETE("", handler.DeletePost)     // 删除博客
			postv1.GET(":postID", handler.GetPost)    // 查询博客详情
			postv1.GET("", handler.ListPost)          // 查询博客列表
		}
	}
}

// Run 运行应用
func (s *Server) Run() error {
	// 运行 HTTP 服务器
	// 打印一条日志，用来提示 HTTP 服务已经起来，方便排错
	slog.Info("Start to listening the incoming requests on http address", "addr", s.cfg.Addr)
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error(err.Error())
			os.Exit(1)
		}
	}()

	// 创建一个 os.Signal 类型的 channel，用于接收系统信号
	quit := make(chan os.Signal, 1)
	// 当执行 kill 命令时（不带参数），默认会发送 syscall.SIGTERM 信号
	// 使用 kill -2 命令会发送 syscall.SIGINT 信号（例如按 Ctrl+C 触发）
	// 使用 kill -9 命令会发送 syscall.SIGKILL 信号，但 SIGKILL 信号无法被捕获，因此无需监听和处理
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 阻塞程序，等待从 quit channel 中接收到信号
	<-quit

	slog.Info("Shutting down server...")

	// 优雅关闭服务
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 先关闭依赖的服务，再关闭被依赖的服务
	// 10 秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过 10 秒就超时退出
	if err := s.srv.Shutdown(ctx); err != nil {
		slog.Error("Insecure Server forced to shutdown", "err", err)
		return err
	}

	slog.Info("Server exited")

	return nil
}
