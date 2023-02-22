package constants

const (
	MySQLDefaultDSN = "root:root@tcp(1.15.222.156:3307)/douyin?charset=utf8&parseTime=True&loc=Local"
	MongoDefaultDSN = "mongodb://root:rootlc@8.130.33.1:27017"
	EtcdAddress     = "0.0.0.0:2379"
	RedisAddress    = "1.15.222.156:6380"
	//RedisAddress                         = "1.15.222.156:6380"
	VideoServiceName                     = "video_service"
	CPURateLimit                 float64 = 80.0
	DefaultLimit                         = 10
	VideoTableName                       = "video"
	UserTableName                        = "user"
	RelationTableName                    = "relation"
	FavoriteTableName                    = "favorite"
	UserServiceMachineID         int64   = 1
	VideoServiceMachineID        int64   = 2
	CommentServiceMachineID      int64   = 3
	FavoriteServiceMachineID     int64   = 4
	ApiServiceWithHostPorts              = "0.0.0.0:8080"
	UserServiceWithHostPorts             = "0.0.0.0:8081"
	VideoServiceWithHostPorts            = "0.0.0.0:8082"
	FavoriteServiceWithHostPorts         = "0.0.0.0:8083"
	CommentServiceWithHostPorts          = "0.0.0.0:8084"
	RelationServiceWithHostPorts         = "0.0.0.0:8085"
	MessageServiceWithHostPorts          = "0.0.0.0:8086"
)
