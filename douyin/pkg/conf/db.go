package conf

type db struct {
	Username string
	Password string
	Host     string
	Port     int
	DbName   string
	Timeout  string
}

var DB = db{
	Username: "root",
	Password: "root",
	Host:     "1.15.222.156",
	Port:     3307,
	DbName:   "douyin-lc",
	Timeout:  "10s",
}
