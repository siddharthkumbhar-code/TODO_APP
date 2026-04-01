package repository

import(
	"database/sql"
	"go-sqlite/models"
	"log"
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