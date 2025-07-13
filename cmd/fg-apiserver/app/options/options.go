package options

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// MYSQLOptions defines options for mysql database.
type MYSQLOptions struct {
	Addr                   string        `json:"addr,omitempty" mapstructure:"addr" `
	Username               string        `json:"username,omitempty" mapstructure:"username"`
	Password               string        `json:"-" mapstructure:"password"`
	Database               string        `json:"database" mapstructure:"database"`
	MaxIdleConnections     int           `json:"max-idle-connections,omitempty" mapstructure:"max-idle-connections,omitempty"`
	MaxOpenConnections     int           `json:"max-open-connections,omitempty" mapstructure:"max-open-connections"`
	MaxConnectionsLifeTime time.Duration `json:"max-connection-left-time,omitempty" mapstructure:"max-connection-left-time"`
}

// NewMySQLOptions crteate a `zero` value instantce.
func NewMySQLOptions() *MYSQLOptions {
	return &MYSQLOptions{
		Addr:                   "127.0.0.1:3306",
		Username:               "onex",
		Password:               "onex(#)666",
		Database:               "onex",
		MaxIdleConnections:     100,
		MaxOpenConnections:     100,
		MaxConnectionsLifeTime: time.Duration(10) * time.Second,
	}
}

type ServerOptions struct {
	MYSQLOptions *MYSQLOptions `json:"mysql" mapstructure:"mysql"`
}

// NewServerOptions 创建带有默认值的 ServerOptions 实例
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		MYSQLOptions: NewMySQLOptions(),
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

	if o.MYSQLOptions.MaxConnectionsLifeTime <= 0 {
		return fmt.Errorf("MySQL max connection lifetime must be greater than 0")
	}

	return nil
}
