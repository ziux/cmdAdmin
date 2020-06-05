package vars

import "fmt"

type DBVar struct {
	Addr string
	User string
	Password string
	DataBase string
}

type RedisVar struct {
	Used bool
	Addr string
	Password string
	DB int

}
func (db *DBVar) MakeStr() string{
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",db.User,db.Password,db.Addr,db.DataBase)

}


var (
	DB = &DBVar{"127.0.0.1:3306","root","0000","test"}
	Redis = &RedisVar{false,"127.0.0.1:3546","",1}
)
