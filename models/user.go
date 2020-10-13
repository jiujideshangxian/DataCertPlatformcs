package models

import "crypto/md5"

type User struct {
       Id  int `form:id`
       Phone  int `form:phone`
       Password  int `form:password`
}


/**
 * 将用户的信息保存到数据库中
 */
func (u User)AddUser()(int64,error)  {
       hashMd5:=md5.New()

}
