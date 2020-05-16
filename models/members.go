package models

import (
	"DellDaven_API/repository"
	"errors"
	"github.com/astaxie/beego/orm"
)

var (
	UserList map[string]*Members
)

//Members
type Members struct {
	Id       int    `json:"id"`
	Username string `orm:"column(username)" json:"username"`
	Pass     string `orm:"column(pass)" json:"pass"`
	Admin    bool   `orm:"column(admin)" json:"admin"`
}

func init() {
	orm.RegisterModel(new(Members))
}

func AddUser(u *Members) (id int64, err error) {
	db := repository.GetSession()
	count, err := db.QueryTable(u).Filter("username", u.Username).Count()
	if count == 0 {
		id, err = db.InsertOrUpdate(u)
	} else {
		err = errors.New("-1")
	}

	return id, err
}

func GetAllUsers() map[string]*Members {
	return UserList
}

func Login(username, password string) (userLog Members, err error) {
	db := repository.GetSession()
	err = db.QueryTable(userLog).Filter("username", username).Filter("pass", password).One(&userLog)
	return userLog, err
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
