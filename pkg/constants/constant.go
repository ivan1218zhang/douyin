package constants

const (
	VideoTableName          = "video"
	UserTableName           = "user"
	MySQLDefaultDSN         = "root:root@tcp(1.15.222.156:3307)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress             = "127.0.0.1:2379"
	CPURateLimit    float64 = 80.0
	DefaultLimit            = 10
)
