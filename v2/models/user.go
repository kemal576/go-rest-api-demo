package models

type User struct {
	ID        int    `json:"ID"`
	FirstName string `json:"Firstname"`
	LastName  string `json:"Lastname"`
	UserName  string `json:"Username"`
	Status    bool   `json:"Status"`
	Age       int    `json:"Age"`
}

func NewUser(Id, age int, firstName, lastName, userName string, status bool) *User {
	x := new(User)
	x.ID = Id
	x.FirstName = firstName
	x.LastName = lastName
	x.UserName = userName
	x.Age = age
	x.Status = status
	return x
}
