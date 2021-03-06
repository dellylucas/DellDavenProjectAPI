package models

import (
	"DellDaven_API/repository"
	"crypto/rand"
	"fmt"
	"github.com/astaxie/beego/orm"
)

var (
	UserList map[string]*Members
)

//Members
type Members struct {
	Id       int    `json:"id"`
	FullName string `orm:"column(fullname)" json:"fullname"`
	Email    string `orm:"column(email);unique" json:"email"`
	Pass     string `orm:"column(pass)" json:"pass"`
	Admin    bool   `orm:"column(admin)" json:"admin"`
	Token    string `orm:"column(token)" json:"token"`
}

func init() {
	orm.RegisterModel(new(Members))
}

func AddUser(u *Members) (value int32) {
	db := repository.GetSession()
	count, err := db.QueryTable(u).Filter("email", u.Email).Count()
	if err == nil && count == 0 {
		_, err = db.InsertOrUpdate(u)
		if err == nil {
			value = 100 //OK
		} else {
			value = 300 //error BD
		}
	} else if err == nil && count > 0 {
		value = 200 //ya existe
	} else {
		value = 300 //error BD
	}
	return value
}

func GetAllUsers() map[string]*Members {
	return UserList
}

func Login(user *Members) error {
	db := repository.GetSession()
	err := db.QueryTable(user).Filter("email", user.Email).Filter("pass", user.Pass).One(user)
	if err == nil && user.Id != 0 {
		user.Token = tokenGenerator()
		db.Update(user, "token")
	}
	return err
}

func tokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

/*
func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Usernamew != "" {
			u.Usernamew = uu.Usernamew
		}
		if uu.Passwordw != "" {
			u.Passwordw = uu.Passwordw
		}

		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Usernamew == username && u.Passwordw == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}*/
