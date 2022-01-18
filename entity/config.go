package entity

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Host string `json:"host"`
}

func GetConfig() Config {
	var config Config //定义一个结构体变量

	//读取yaml文件到缓存中
	r, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		fmt.Print(err)
	}

	//yaml文件内容影射到结构体中
	err1 := yaml.Unmarshal(r, &config)
	if err1 != nil {
		fmt.Println("error")
	}

	//通过访问结构体成员获取yaml文件中对应的key-value
	return config
}
