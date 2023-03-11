package main

//Penamaan struct harus diawali dengan huruf kapital
type User struct {
	id       int
	username int
	password int
}

//Penamaan struct harus diawali dengan huruf kapital dan harus menggunakan PascalCase dan penamaan variable jangan terlalu singkat
type UserService struct {
	tipe []User
}

//penamaan function harus menggunakan camelCase
func (user UserService) getAllUsers() []User {
	return user.tipe
}

//penamaan function harus menggunakan camelCase
func (user UserService) getUserById(id int) User {
	for _, value := range user.tipe {
		if id == value.id {
			return value
		}
	}

	return User{}
}
