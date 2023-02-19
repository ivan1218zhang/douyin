package constants

const (
	MySQLDefaultDSN = "root:root@tcp(1.15.222.156:3307)/douyin?charset=utf8&parseTime=True&loc=Local"
	MongoDefaultDSN = "mongodb://root:rootlc@8.130.33.1:27017"
	EtcdAddress     = "127.0.0.1:2379"
	RedisAddress    = "localhost:6379"
	// RedisAddress                    = "1.15.222.156:6379"
	VideoServiceName                     = "video_service"
	CPURateLimit                 float64 = 80.0
	DefaultLimit                         = 10
	VideoTableName                       = "video"
	UserTableName                        = "user"
	UserServiceMachineID         int64   = 1
	VideoServiceMachineID        int64   = 2
	CommentServiceMachineID      int64   = 3
	ApiServiceWithHostPorts              = ":8080"
	UserServiceWithHostPorts             = ":8081"
	VideoServiceWithHostPorts            = ":8082"
	FavoriteServiceWithHostPorts         = ":8083"
	CommentServiceWithHostPorts          = ":8084"
	RelationServiceWithHostPorts         = ":8085"
	MessageServiceWithHostPorts          = ":8086"
)
