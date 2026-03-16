package main

import (
	
	"net/http"
	"database/sql"
	"encoding/json"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

type Task struct{
	ID int  `json:"id"`
	NAME string `json:"name"`
	STATUS  string `json:"status"`
}

func main(){
	db, err := sql.Open("sqlite3", "./test.db")
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close();

	query := `CREATE TABLE IF NOT EXISTS tasks(
	     id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		 name TEXT ,
		 status TEXT
	
	);`

	_,err = db.Exec(query)

	if err!=nil{
		log.Fatal(err)
	}
	log.Println("table creates succesfully")

	// to insert a task into database
	http.HandleFunc("/insert",InsertTask(db))

	// to get the all tasks from database
	http.HandleFunc("/getall",GetAll(db))

	//to get one task by id 
	http.HandleFunc("/get",gettask(db))

	// to rename the task by id 
	http.HandleFunc("/rename",RenameTask(db))

	//to change the status of the task
	http.HandleFunc("/ChangeStatus",ChangeStatus(db))

	//to delete the task from the table 
	http.HandleFunc("/delete",DeleteTask(db))

	//delete the task which are done 
	http.HandleFunc("/DeleteCompletedTask",DeleteCompletedTask(db))

	http.ListenAndServe(":8080", nil)
}

func DeleteCompletedTask(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter , r *http.Request){

		query :=`DELETE FROM tasks
		        WHERE status="DONE"    
		`

		_,err := db.Exec(query)
		
		 if err!=nil{
			log.Println("somthing went wrong  while deleting the task ")
			return 
		}
		json.NewEncoder(w).Encode(map[string]string{
			"message":"completed task deleted succesfully ",
		})
	}
}


func DeleteTask(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		id := r.URL.Query().Get("id")

		query :=`DELETE FROM tasks
		        WHERE id=?`

		_,err := db.Exec(query,id)
		
		if err!=nil{
			log.Println("error while deleting the task")
			return 
		}

		json.NewEncoder(w).Encode(map[string]string{
			"message":"delete the task succesfully",
		})
	}
}


func ChangeStatus(db *sql.DB) http.HandlerFunc{
	return func( w http.ResponseWriter, r *http.Request){

		id := r.URL.Query().Get("id")

		query := `UPDATE tasks
		          SET status=CASE
				  when status="pending" THEN "DONE"
				  ELSE "pending"
				  END
				  WHERE id=?
				  `

	   _,err := db.Exec(query,id)
	   
	   if err!=nil{
		 log.Println("error in update  the status")
		 return 
	   }

	   json.NewEncoder(w).Encode(map[string]string{
		"message":"update the status succesfully ",
	   })
	}
}


func RenameTask(db *sql.DB) http.HandlerFunc{
	return func (w http.ResponseWriter , r *http.Request){
	

		id:= r.URL.Query().Get("id")
		name := r.URL.Query().Get("name")

		query:=`UPDATE tasks
		        SET name=?
				WHERE  id=? `

		_,err:= db.Exec(query,name,id)

		if err!=nil{
			log.Println("error in updatting the data ")
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"message":"rename the task succesfully",
		})
	}
}

func gettask(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter , r *http.Request){
		var task Task;

	 id := r.URL.Query().Get("id")

	 query := `SELECT id,name,status FROM tasks where id=?`

	 err:= db.QueryRow(query,id).Scan(&task.ID,&task.NAME,&task.STATUS)

	 if err!=nil{
		log.Println("error in fetching the data");
		return 
	 }
	 json.NewEncoder(w).Encode(task)
	}
}

func GetAll(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var list []Task;

		query := `SELECT * FROM tasks`

		res,err := db.Query(query);
		 if err!=nil{
			log.Println("wrong in fetching the data")
		 }
		 defer res.Close();

        for res.Next(){
			var task Task;

			err := res.Scan(&task.ID,&task.NAME,&task.STATUS)
			if err!=nil{
				log.Println("wrong in the scanning the data")
				return 
			}

			list = append(list,task)
		}  

		json.NewEncoder(w).Encode(list)
	}
}


func InsertTask(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
        var newtask Task;

		 err := json.NewDecoder(r.Body).Decode(&newtask)

		 if err!=nil{
			log.Println("error in fetching the data")
		 }

		 query := `INSERT INTO tasks (name , status) VALUES(?,?)`

		 _, err = db.Exec(query,newtask.NAME,newtask.STATUS)

		 if err != nil{
			log.Println("somthing went wrong to inserting the data ")
			return 
		 }

	     json.NewEncoder(w).Encode(map[string]string{
			"message":"the task inserted succesfully into database ",
		 })

	}
}