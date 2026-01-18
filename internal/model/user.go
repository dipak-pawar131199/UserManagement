package model

type User struct {
	UserId    int    `db:"userId"`
	UserName  string `db:"userName"`
	UserEmail string `db:"userEmail"`
	Password  string `db:"password"`
}
