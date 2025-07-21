package options

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLOptions defines options for mysql database.
type MySQLOptions struct {
	Addr                  string        `json:"addr,omitempty" mapstructure:"addr" `
	Username              string        `json:"username,omitempty" mapstructure:"username"`
	Password              string        `json:"-" mapstructure:"password"`
	Database              string        `json:"database" mapstructure:"database"`
	MaxIdleConnections    int           `json:"max-idle-connections,omitempty" mapstructure:"max-idle-connections,omitempty"`
	MaxOpenConnections    int           `json:"max-open-connections,omitempty" mapstructure:"max-open-connections"`
	MaxConnectionLifeTime time.Duration `json:"max-connection-left-time,omitempty" mapstructure:"max-connection-left-time"`
}

// NewMySQLOptions crteate a `zero` value instantce.
func NewMySQLOptions() *MySQLOptions {
	return &MySQLOptions{
		Addr:                  "127.0.0.1:3306",
		Username:              "onex",
		Password:              "onex(#)666",
		Database:              "onex",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: time.Duration(10) * time.Second,
	}
}

// Validate verifies flags passed to MySQLOptions.
func (o *MySQLOptions) Validate() error {
	// 验证MySQL地址格式
	if o.Addr == "" {
		return fmt.Errorf("MySQL server address cannot be empty")
	}
	// 检查地址格式是否为host:port
	host, portStr, err := net.SplitHostPort(o.Addr)
	if err != nil {
		return fmt.Errorf("invalid MySQL address format '%s': %w", o.Addr, err)
	}
	// 验证端口是否为数字
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("invalid MySQL port: %s", portStr)
	}
	// 验证主机名是否为空
	if host == "" {
		return fmt.Errorf("MySQL hostname cannot be empty")
	}

	// 验证凭据和数据库名
	if o.Username == "" {
		return fmt.Errorf("MySQL username cannot be empty")
	}

	if o.Password == "" {
		return fmt.Errorf("MySQL password cannot be empty")
	}

	if o.Database == "" {
		return fmt.Errorf("MySQL database name cannot be empty")
	}

	// 验证连接池参数
	if o.MaxIdleConnections <= 0 {
		return fmt.Errorf("MySQL max idle connections must be greater than 0")
	}

	if o.MaxOpenConnections <= 0 {
		return fmt.Errorf("MySQL max open connections must be greater than 0")
	}

	if o.MaxIdleConnections > o.MaxOpenConnections {
		return fmt.Errorf("MySQL max idle connections cannot be greater than max open connections")
	}

	if o.MaxConnectionLifeTime <= 0 {
		return fmt.Errorf("MySQL max connection lifetime must be greater than 0")
	}

	return nil
}

// DSN(Data Source Name) return DSN from MySQLOptions.
// 根据 MySQLOptions 结构体中的配置信息生成 MySQL 数据库连接字符串。
// DSN 是一个标准化的字符串格式，用于描述如何连接到数据库。它包含了连接到数据库所需的所有必要信息。
func (o *MySQLOptions) DSN() string {
	return fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		o.Username,
		o.Password,
		o.Addr,
		o.Database,
		true,
		"Local",
	)
}

// NewDB create mysql store with the given config.
// gorm 是一个 Go 语言的 ORM（对象关系映射）库，用于在 Go 应用程序中操作数据库。
func (o *MySQLOptions) NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(o.DSN()), &gorm.Config{
		// PrepareStmt 用于启用或禁用预处理语句功能。
		// 预处理语句是一种数据库优化技术，它允许数据库服务器在执行查询之前对其进行分析、优化和编译。
		// 启用预处理语句可以提高数据库的性能，尤其是在执行多次相同查询时。
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns 设置数据库连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(o.MaxIdleConnections)
	// SetMaxOpenConns 设置数据库连接池中打开的最大连接数。
	sqlDB.SetMaxOpenConns(o.MaxOpenConnections)
	// SetConnMaxLifetime 设置数据库连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(o.MaxConnectionLifeTime)

	return db, nil
}
