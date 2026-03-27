package handlers

import (
	"database/sql"
	"encoding/json"
	"go-sqlite/models"
	"log"
	"net/http"
)

func GetAllUsers(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

		query:=`SELECT * FROM Users`
		rows,err:=db.Query(query)
		if err!=nil{
			log.Fatal("Internal server Error")
		}
		Users:=[]models.User{}
		for rows.Next(){
			var user models.User
			rows.Scan(&user.Userid,&user.UserName,&user.Email)
			Users = append(Users, user)
		}
		rows.Close()
		w.Header().Set("Content-type","application/json")
		json.NewEncoder(w).Encode(Users)
	}
}