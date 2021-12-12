package repositories

import (
	"errors"

	. "github.com/kemal576/go-rest-api-demo/models"
)

type UserRepository struct {
	users []User
}

func (u *UserRepository) AppendUsers() {
	u.users = append(u.users, *NewUser(1, 21, "Kemal", "Şahin", "kemal576", true),
		*NewUser(2, 30, "Cihan", "Özhan", "cihanozhan", true),
		*NewUser(3, 14, "Ahmet", "Denizci", "ahmet123", true),
		*NewUser(4, 45, "Fatma", "Arslan", "fatmarslan", true),
		*NewUser(5, 27, "Zeynep", "Erdoğdu", "zzeynep", true))
}

func (u *UserRepository) Add(user *User) {
	u.users = append(u.users, *user)
}

func (u *UserRepository) Update(_user *User) {
	for i, user := range u.users {
		if user.ID == _user.ID {
			u.users[i] = *_user
			println("update scope")
		}
	}
}

//Soft Delete
func (u *UserRepository) Delete(ID int) bool {
	for i, user := range u.users {
		if user.ID == ID {
			u.users[i].Status = false
			return true
		}
	}
	return false
}

func (u *UserRepository) GetActiveUsers() []User {
	var activeUsers []User
	for _, user := range u.users {
		if user.Status {
			activeUsers = append(activeUsers, user)
		}
	}
	return activeUsers
}

func (u *UserRepository) GetAll() []User {
	return u.users
}

func (u *UserRepository) GetById(ID int) (User, error) {
	for _, user := range u.users {
		if user.ID == ID && user.Status {
			return user, nil
		}
	}
	return *NewUser(0, 0, "", "", "", false), errors.New("kullanıcı bulunamadı")
}
