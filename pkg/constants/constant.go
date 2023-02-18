package constants

const (
	MySQLDefaultDSN = "root:root@tcp(1.15.222.156:3307)/douyin?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress     = "127.0.0.1:2379"
	RedisAddress    = "localhost:6379"
	// RedisAddress                    = "1.15.222.156:6379"
	VideoServiceName                = "video_service"
	CPURateLimit            float64 = 80.0
	DefaultLimit                    = 10
	VideoTableName                  = "video"
	UserTableName                   = "user"
	FavoriteTableName            = "favorite"
	UserServiceMachineID    int64   = 1
	VideoServiceMachineID   int64   = 2
	CommentServiceMachineID int64   = 3
)

