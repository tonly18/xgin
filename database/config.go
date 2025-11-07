package database

const (
	RoleDefault = "default"

	DbCharset      = "utf8mb4"
	DbCollation    = "utf8mb4_general_ci"
	DbMaxOpenConns = 10
	DbMaxIdleConns = 2
	DbMaxLifetime  = 3600
	DbMaxIdleTime  = 600
)

type MySqlConfig struct {
	Role         string         `json:"role"`
	Host         string         `json:"host"`
	Port         int            `json:"port"`
	Dbname       string         `json:"dbname"`
	Username     string         `json:"username"`
	Password     string         `json:"password"`
	Charset      string         `json:"charset"`
	Collation    string         `json:"collation"`
	MaxIdleConns int            `json:"max_idle_conns"`
	MaxOpenConns int            `json:"max_open_conns"`
	MaxLifetime  int            `json:"max_lifetime"`
	MaxIdleTime  int            `json:"max_idle_time"`
	Slave        []*MySqlConfig `json:"slave"`
}

type RedisClusterConfig struct {
	Host     []string `json:"host"`
	Username string   `json:"username"`
	Password string   `json:"password"`

	PoolSize     int `json:"pool_size"`      //链接池最大链接数
	MaxIdleConns int `json:"max_idle_conns"` //最大空闲链接数
	MinIdleConns int `json:"min_idle_conns"` //最小空闲链接数
	
	TLS        bool   `json:"tls"`
	ServerName string `json:"server_name"`
}
