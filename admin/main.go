package main

import (
	"golang.org/x/crypto/bcrypt"
	"my_blog/admin/internal/di"
)

func main() {
	//_ = ScryptPw("123456")

	_, err := di.Init()
	if err != nil {
		panic(err)
	}
}

func ScryptPw(password string) string {
	const cost = 10

	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		//log.Fatal(err)
	}
	s := string(HashPw)

	return s
}
