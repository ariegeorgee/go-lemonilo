package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	// register model

	// set default database
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		"root",
		"Kairos30121993",
		"localhost",
		"3306",
		"golemonilo",
		"utf8")

	orm.RegisterDataBase("default", "mysql", conn)
	orm.RegisterModel(new(User))
}

// TableName of User
func (u *User) TableName() string {
	return "users"
}

// UserTable Table
func UserTable() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(User))
}

// User Models
type User struct {
	ID       int    `json:"id" orm:"column(id);pk;auto;unique"`
	Name     string `json:"name" orm:"column(name)"`
	Email    string `json:"email" orm:"column(email)"`
	Address  string `json:"address" orm:"column(address)"`
	Password string `json:"password,omitempty" orm:"column(password)"`
	Token    string `json:"token" orm:"-"`
}

func AddUser(u User) string {
	o := orm.NewOrm()
	o.Begin()

	secretOwn := "golemonilo"
	u.Password, _ = HashPassword(u.Password + secretOwn)

	id, err := o.Insert(&u)

	if err == nil {
		o.Commit()
	} else {
		o.Rollback()
	}

	return string(id)
}

func GetUser(uid int) (u User, err error) {
	o := orm.NewOrm()
	u.ID = uid
	err = o.Read(&u, "ID")
	return u, err
}

func GetAllUsers() []User {
	var usr []User
	orm.NewOrm().QueryTable("users").All(&usr)
	return usr
}

func UpdateUser(usr User) (a User, err error) {
	var usrInput User
	o := orm.NewOrm()
	usrInput.ID = usr.ID
	err = o.Read(&usrInput, "ID")
	if err == nil && usrInput.ID != 0 {
		if usr.Password != "" {
			secretOwn := "golemonilo"
			usr.Password, _ = HashPassword(usr.Password + secretOwn)
		}
		if _, err := o.Update(&usrInput); err == nil {
			return usr, err
		}
	}
	return usr, err
}

func Login(usrs User) (isSuccess bool, usr User) {
	secretOwn := "golemonilo"
	// pwd, _ := HashPassword(usrs.Password + secretOwn)
	user := User{Email: usrs.Email}

	o := orm.NewOrm()
	err := o.Read(&user, "Email")

	if err == nil && user.ID != 0 {
		if CheckPasswordHash(user.Password, usrs.Password+secretOwn) {
			user.Token = user.Password
			user.Password = ""
			return true, user
		}
	}
	return false, user
}

func DeleteUser(id int) int64 {
	o := orm.NewOrm()
	num, _ := o.Delete(&User{ID: id})
	return num
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(existing, incoming string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(existing), []byte(incoming))
	return err == nil
}
