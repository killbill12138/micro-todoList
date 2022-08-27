package conf

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/ini.v1"
)

var Config *ConfigType
var MongoDBClient *mongo.Client

type ConfigType struct {
	Service `ini:"service"`
	Mysql   `ini:"mysql"`
	MongoDB `ini:"mongodb"`
	Redis   `ini:"redis"`
}

type Service struct {
	AppMode  string `ini:"AppMode"`
	HttpPort string `ini:"HttpPort"`
}

type Mysql struct {
	Db       string `ini:"Db"`
	Host     string `ini:"Host"`
	Port     string `ini:"Port"`
	User     string `ini:"User"`
	Password string `ini:"Password"`
	DbName   string `ini:"DbName"`
}

type MongoDB struct {
	UserName string `ini:"UserName"`
	Host     string `ini:"Host"`
	Port     string `ini:"Port"`
	Password string `ini:"Password"`
}

type Redis struct {
	Db          string `ini:"Db"`
	HostAndPort string `ini:"HostAndPort"`
	Password    string `ini:"Password"`
	Name        string `ini:"Name"`
}

func init() {
	Config = new(ConfigType)
	// 从本地读取环境(结构体反射法)
	fmt.Println("init")
	err := ini.MapTo(Config, "./conf/config.conf")
	if err != nil {
		log.Fatal("config field err:", err)
	} else {
		fmt.Println(Config)
	}
}
