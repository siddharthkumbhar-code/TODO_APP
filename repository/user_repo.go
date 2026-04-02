package repository

import(
	"database/sql"
	"go-sqlite/models"
	"log"
	
	
)

type UserRepository struct{
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository{
	return &UserRepository{db:db}
}

func (repo *UserRepository) InsertUser(newuser models.Users) error{
	query :=`INSERT INTO users(username,email) VALUES(?,?) `

	_, err := repo.db.Exec(query, newuser.Username, newuser.Email)

		if err != nil {
			log.Println("error while inserting the user ", err)
			return err
		}
     
      return nil
}