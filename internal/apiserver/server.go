package apiserver

import (
	"errors"
	"log/slog"
	"net/http"

	genericoptions "github.com/TobyIcetea/fastgo/pkg/options"
	"github.com/gin-gonic/gin"
)

// Config 配置结构体，用于存储应用相关的配置
// 不用 viper.Get，是因为这种方式能更加清晰的知道应用提供了哪些配置项
type Config struct {
	MySQLOptions *genericoptions.MySQLOptions
	Addr         string
}

// Server 定义了一个服务器结构体类型
type Server struct {
	cfg *Config
	srv *http.Server
}

// NewServer 根据配置创建服务器
func (cfg *Config) NewServer() (*Server, error) {
	// 创建 Gin 引擎
	engine := gin.New()

	// 注册 404 Handler
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PageNotFound", "message": "Page not found."})
	})

	// 注册 /healthz handler.
	engine.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// 创建 HTTP Server 实例
	httpsrv := &http.Server{Addr: cfg.Addr, Handler: engine}

	return &Server{cfg: cfg, srv: httpsrv}, nil
}

// Run 运行应用
func (s *Server) Run() error {
	// 运行 HTTP 服务器
	// 打印一条日志，用来提示 HTTP 服务已经起来，方便排错
	slog.Info("Start to listening the incoming requests on http address", "addr", s.cfg.Addr)
	if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
