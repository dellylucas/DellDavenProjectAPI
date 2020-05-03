package models

import (
	"DellDaven_API/repository"
	"github.com/astaxie/beego/orm"
	"log"
)

//Members
type Members struct {
	ID        int    `orm:"column(ID)"`
	Usernamew string `orm:"column(Usernamew)" json:"usernamew"`
	Passwordw string `orm:"column(Passwordw)" json:"passwordw"`
}

func init() {
	orm.RegisterModel(new(Members))
}

func AddUser(u Members) int {
	u.ID = 0
	return u.ID
}

/*
func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}*/

func GetAllUsers(user *Members) {
	db := repository.GetSession()

	if _, err := db.InsertOrUpdate(user); err != nil {
		log.Println(err)
	}
	//return user
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
