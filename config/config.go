package config

import (
	"eletronicMall/dao"
	"fmt"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	HttpPort string
	AppModel string
	DB       string
	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
	DBName   string

	CacheDB     string
	RedisAddr   string
	RedisPass   string
	RedisDBName string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string

	Host        string
	ProductPath string
	AvatarPath  string
)

func Init() {
	// 本地读取环境变量
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		panic(err)
	}

	LoadServer(file)
	LoadDB(file)
	LoadRedis(file)
	LoadEmail(file)
	LoadPath(file)

	// mysql read, master db
	pathRead := strings.Join([]string{DBUser, ":", DBPass, "@tcp(", DBHost, ":", DBPort, ")/", DBName, "?charset=utf8mb4&parseTime=true"}, "")
	// mysql write, slave db
	pathWrite := strings.Join([]string{DBUser, ":", DBPass, "@tcp(", DBHost, ":", DBPort, ")/", DBName, "?charset=utf8mb4&parseTime=true"}, "")

	fmt.Println(pathRead)
	fmt.Println(pathWrite)
	dao.DataBase(pathRead, pathWrite)

}

// server
func LoadServer(file *ini.File) {
	AppModel = file.Section("service").Key("AppModel").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

// mysql 读写分离
func LoadDB(file *ini.File) {
	DB = file.Section("mysql").Key("DB").String()
	DBHost = file.Section("mysql").Key("DBHost").String()
	DBPort = file.Section("mysql").Key("DBPort").String()
	DBUser = file.Section("mysql").Key("DBUser").String()
	DBPass = file.Section("mysql").Key("DBPass").String()
	DBName = file.Section("mysql").Key("DBName").String()
}

func LoadRedis(file *ini.File) {
	CacheDB = file.Section("redis").Key("CacheDB").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPass = file.Section("redis").Key("RedisPass").String()
	RedisDBName = file.Section("redis").Key("RedisDBName").String()
}

func LoadEmail(file *ini.File) {
	ValidEmail = file.Section("email").Key("ValidEmail").String()
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String()
	SmtpPass = file.Section("email").Key("SmtpPass").String()
}

func LoadPath(file *ini.File) {
	Host = file.Section("path").Key("Host").String()
	ProductPath = file.Section("path").Key("ProductPath").String()
	AvatarPath = file.Section("path").Key("AvatarPath").String()
}
