package helper

var Connection string

func init() {
	Connection = "MySQL"
}

func GetDatabase() string {
	return Connection
}