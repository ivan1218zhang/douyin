package constants

const (
	MySQLDefaultDSN              = "root:root@tcp(1.15.222.156:3307)/douyin?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress                  = "127.0.0.1:2379"
	VideoServiceName             = "video_service"
	UserServiceName              = "user_service"
	CPURateLimit         float64 = 80.0
	DefaultLimit                 = 10
	VideoTableName               = "video"
	UserTableName                = "user"
	UserServiceMachineID int64   = 1
)
