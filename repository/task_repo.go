package repository

import(
	"database/sql"
	"go-sqlite/models"
	"log"
	"time"
	
)


type TaskRepository struct{
	db *sql.DB
}

func NewTaskRepository (db *sql.DB) *TaskRepository{
	return &TaskRepository{db:db}
}

func (r *TaskRepository) GetTaskByUserId(query string,params []interface{})([]models.Task,error){


	var tasklist []models.Task
	rows,err := r.db.Query(query,params...)
	if err!=nil{
		log.Println("error in execution the query",err)
		return nil,err
	}
	for rows.Next(){
		var task models.Task

		err = rows.Scan(
			&task.Id,
			&task.Name,
			&task.Status,
			&task.UserId,
			&task.CreatedAt,
			&task.UpdatedAt,

		)
		if err!=nil{
			log.Println("error in scanning the  data",err)
			return nil,err
		}
		tasklist = append(tasklist,task)
	}
	return tasklist,nil
}


func(r *TaskRepository) InsertTask(newtask models.Task) error{
	
	    query := `INSERT INTO tasks1 (name ,status,userid,createdAt,updatedAt) VALUES(?,?,?,?,?)`
	
		now := time.Now().UTC().Format(time.RFC3339)
		_, err := r.db.Exec(query, newtask.Name, newtask.Status, newtask.UserId, now, now)

		if err != nil {
			log.Println("somthing went wrong to inserting the data ", err)
			//http.Error(writer,"Error while creating the task",500)
			return err
		}
   return nil
}