package handlers

import (
	"database/sql"
	"encoding/json"
	"go-sqlite/models"
	"log"
	"net/http"
	"strconv"
)

func PendingTasks(db *sql.DB) http.HandlerFunc{

	return  func(w http.ResponseWriter, r *http.Request) {

		userId:=r.URL.Query().Get("userId")
		if userId==""{
			log.Println("Enter a userId")
		}
		uid,err:=strconv.Atoi(userId)
		if err!=nil{
			log.Println("Enter a valid userId")
		}

		query:=`SELECT * FROM tasks
				WHERE userId=?
				AND status='pending'`
		rows,err:=db.Query(query,uid)
		tasks:=[]models.Task{}
		for rows.Next(){
			var task models.Task
			rows.Scan(&task.ID,&task.NAME,&task.STATUS,&task.CreatedAt,&task.UpdatedAt,&task.Userid)
			tasks = append(tasks, task)
		}
		w.Header().Set("Content-type","application/json")
		json.NewEncoder(w).Encode(tasks)
	}
}