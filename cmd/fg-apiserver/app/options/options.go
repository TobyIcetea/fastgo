package options

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/TobyIcetea/fastgo/internal/apiserver"
	genericoptions "github.com/TobyIcetea/fastgo/pkg/options"
)

type ServerOptions struct {
	MYSQLOptions *genericoptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`
	Addr         string                       `json:"addr" mapstructure:"addr"`
	// JWTKey 定义 JWT 密钥.
	JWTKey string `json:"jwt-key" mapstructure:"jwt-key"`
	// Expiration 定义 JWT Token 的过期时间.
	Expiration time.Duration `json:"expiration" mapstructure:"expiration"`
}

// NewServerOptions 创建带有默认值的 ServerOptions 实例
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		MYSQLOptions: genericoptions.NewMySQLOptions(),
		Addr:         "0.0.0.0:6666",
		Expiration:   2 * time.Hour,
	}
}

// Validate 校验 ServerOptions 中的选项是否合法
// 提示：Validate 方法中的具体校验逻辑可以由 GPT 自动生成。
func (o *ServerOptions) Validate() error {
	// 验证 MySQL 地址格式
	if o.MYSQLOptions.Addr == "" {
		return fmt.Errorf("MYSQL server address cannot be empty")
	}

	// 检查地址格式是否为 host:host
	host, portStr, err := net.SplitHostPort(o.MYSQLOptions.Addr)
	if err != nil {
		return fmt.Errorf("invalid MySQL address format '%s': '%w'", o.MYSQLOptions.Addr, err)
	}

	// 验证端口是否为数字
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("invalid MySQL port: %s", portStr)
	}

	// 验证主机名是否为空
	if host == "" {
		return fmt.Errorf("MYSQL hostname cannot be empty")
	}

	// 验证凭证和数据库名
	if o.MYSQLOptions.Username == "" {
		return fmt.Errorf("MySQL username cannot be empty")
	}

	if o.MYSQLOptions.Password == "" {
		return fmt.Errorf("MySQL password cannot be empty")
	}

	if o.MYSQLOptions.Database == "" {
		return fmt.Errorf("MySQL database name cannot be empty")
	}

	// 验证连接池参数
	if o.MYSQLOptions.MaxIdleConnections <= 0 {
		return fmt.Errorf("MySQL max idle connections must be greater than 0")
	}

	if o.MYSQLOptions.MaxOpenConnections <= 0 {
		return fmt.Errorf("MySQL max open connections must be greater than 0")
	}

	if o.MYSQLOptions.MaxIdleConnections > o.MYSQLOptions.MaxOpenConnections {
		return fmt.Errorf("MySQL max idle connections cannot be greater than max open connections")
	}

	if o.MYSQLOptions.MaxConnectionLifeTime <= 0 {
		return fmt.Errorf("MySQL max connection lifetime must be greater than 0")
	}

	// 验证服务器地址
	if o.Addr == "" {
		return fmt.Errorf("server address cannot be empty")
	}

	// 检查地址格式是否为 host:port
	_, portStr, err = net.SplitHostPort(o.Addr)
	if err != nil {
		return fmt.Errorf("invalid server address format '%s': '%w'", o.Addr, err)
	}

	// 验证端口是否为数字且在有效范围内
	port, err = strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("invalid server port: %s", portStr)
	}

	// 校验 JWTKey 长度
	if len(o.JWTKey) < 6 {
		return fmt.Errorf("JWTKey must be at least 6 characters long")
	}

	return nil
}

// Config 基于 ServerOptions 构建 apiserver.Config
func (o *ServerOptions) Config() (*apiserver.Config, error) {
	return &apiserver.Config{
		MySQLOptions: o.MYSQLOptions,
		Addr:         o.Addr,
		JWTKey:       o.JWTKey,
		Expiration:   o.Expiration,
	}, nil
}
