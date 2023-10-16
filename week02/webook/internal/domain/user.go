package domain

import "time"

type User struct {
	Id       int64
	Email    string
	Password string

	Nickname string
	Birthday time.Time
	AboutMe  string

	//UTC 0 的时区
	Ctime time.Time

	//Addr Address
}

//type Address struct {
//	Province string
//	Region   string
//}
