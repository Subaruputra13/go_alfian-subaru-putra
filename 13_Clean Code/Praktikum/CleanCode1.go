package main

//Penamaan struct harus diawali dengan huruf kapital
type User struct {
	Id       int
	Username int
	Password int
}

//Penamaan struct harus diawali dengan huruf kapital dan harus menggunakan PascalCase dan penamaan variable jangan terlalu singkat
type UserService struct {
	Tipe []User
}

//penamaan function harus menggunakan camelCase
func (user UserService) getAllUsers() []User {
	return user.Tipe
}

//penamaan function harus menggunakan camelCase
func (user UserService) getUserById(id int) User {
	for _, value := range user.Tipe {
		if id == value.Id {
			return value
		}
	}

	return User{}
}
