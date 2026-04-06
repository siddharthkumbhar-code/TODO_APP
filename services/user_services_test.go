package services

import "go-sqlite/models"

type FakeUserRepo struct {
	
}

func (f *FakeUserRepo) InsertUser(user models.Users) error{
	return f.InsertUser()
}