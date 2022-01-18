package connect

import (
	"fmt"
	_ "fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func init() {
	var err error
	username := "wineSong" //账号
	password := "123456"   //密码
	//host := "123.57.173.34" //数据库地址，可以是Ip或者域名（测试环境）
	host := "47.96.68.74" //数据库地址，可以是Ip或者域名（生产环境）
	port := 3306          //数据库端口
	Dbname := "winesong"  //数据库名

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。

	//gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	Db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println(err)
	Db.SingularTable(true)
}
