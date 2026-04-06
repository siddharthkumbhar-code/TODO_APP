package services

import "go-sqlite/models"

type FakeUserRepo struct {
	users []models.Users
	err error
}

func (f *FakeUserRepo)InsertUser(user models.Users)error{
	if f.err!=nil{
		return f.err
	}
	f.users=append(f.users, user)
	return nil
}

func (f *FakeUserRepo)GetAllUsers()([]models.Users, error){
	if f.err!=nil{
		return nil,f.err
	}
	return f.users,nil
}

