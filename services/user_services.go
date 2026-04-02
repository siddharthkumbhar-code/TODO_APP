package services

import(
  "go-sqlite/models"
  "log"
  "go-sqlite/repository"
  "strings"
  "net/mail"
)

type UserServices struct{
	repo *repository.UserRepository
}

func NewUserServices(repo *repository.UserRepository) *UserServices{
	return &UserServices{repo:repo}
}
func (userserv *UserServices) InsertUser(newuser models.Users)error{
    	if newuser.Username == "" && newuser.Email == "" {
			
			log.Println("username and email required ")
			return  nil
		}
		if newuser.Username == "" {
			
			log.Println("username required ")
			return nil
		}
		if newuser.Email == "" {
			//http.Error(writer, "Email  Required", 400)
			log.Println("Email required ")
			return nil
		}

		if strings.TrimSpace(newuser.Username) == "" {
			//http.Error(writer, "Username Required", 400)
			log.Println("Username Required")
			return nil
		}
		if strings.TrimSpace(newuser.Email) == "" {
			
			log.Println("Email is required")
			return nil
		}
		_, err := mail.ParseAddress(newuser.Email)
		if err != nil {
			//http.Error(writer, "Invalid Email", http.StatusBadRequest)
			log.Println("Enter a valid Email")
			return err
		}
		if len(newuser.Username) < 2 {
			//http.Error(writer, "Name should greater than 2 characters", 400)
			return nil
		}
		return userserv.repo.InsertUser(newuser)
}
