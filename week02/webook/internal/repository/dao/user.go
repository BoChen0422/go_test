package dao

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"time"
)
import "gorm.io/gorm"

var (
	ErrDuplicateEmail = errors.New("邮箱冲突")
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

type UserDAO struct {
	db *gorm.DB
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Ctime = now
	u.Utime = now
	err := dao.db.WithContext(ctx).Create(&u).Error
	if me, ok := err.(*mysql.MySQLError); ok {
		const duplicateErr uint16 = 1062
		if me.Number == duplicateErr {
			//用户冲突，邮箱冲突
			return ErrDuplicateEmail
		}
	}
	return err
}

func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {

	var u User
	err := dao.db.WithContext(ctx).Where("email=?", email).First(&u).Error
	return u, err
}

func (dao *UserDAO) Updates(ctx context.Context, u User) error {

	return dao.db.Debug().WithContext(ctx).Model(&u).Where("id=?", u.Id).
		Updates(map[string]any{
			"utime":    time.Now().UnixMilli(),
			"nickname": u.Nickname,
			"birthday": u.Birthday,
			"about_Me": u.AboutMe,
		}).Error
}

func (dao *UserDAO) FindById(ctx *gin.Context, id int64) (User, error) {
	var u User
	err := dao.db.Debug().WithContext(ctx).Where("id=?", id).First(&u).Error
	return u, err
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

type User struct {
	Id       int64  `gorm:"primaryKey,autoIncrement"` //主键，自增
	Email    string `gorm:"unique"`                   //唯一
	Password string

	Nickname string `gorm:"type=varchar(128)"`
	Birthday int64
	AboutMe  string `gorm:"type=varchar(4096)"`

	Ctime int64
	Utime int64

	//时区,UTC 0的毫秒数

	//json 存储
	//Addr string
}
