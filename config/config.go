package config

import (
	"encoding/json"
	"os"
)

//服务器配置
type AppConfig struct{
	AppName string `json:"app_name"`
	Port    string `json:"port"`
	StaticPath string `json:"static_path"`
	Mode string `json:"mode"`
}

var ServConfig AppConfig

//初始化服务器配置
func InitConfig() *AppConfig{
	file,err:=os.Open("E://project/iris/ShiZhanCatalogue/config.json")  //有争议
	if err!=nil{
		panic(err.Error())
	}
	decoder:=json.NewDecoder(file)
	conf:=AppConfig()
	err=decoder.Decode(&conf)
	if err!=nil{
		panic(err.Error())
	}
	return &conf
}
