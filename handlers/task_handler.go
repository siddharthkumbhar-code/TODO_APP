package handlers 

import(
	"log"
	"net/http"
	"go-sqlite/services"
	"encoding/json"
)

type TaskHandler struct{
	service *services.TaskServices
}

func NewTaskHandler(service *services.TaskServices) *TaskHandler{
	return &TaskHandler{service:service}
}

func(h *TaskHandler)GetTaskByUserId(writer http.ResponseWriter,request *http.Request){

	    useridstr := request.PathValue("userid")
		status := request.URL.Query().Get("status")
		sortby := request.URL.Query().Get("sortby")
		order := request.URL.Query().Get("order")
		cursor:= request.URL.Query().Get("cursor")
		limitstr := request.URL.Query().Get("limit")
		pagenostr := request.URL.Query().Get("pageno")

		tasks,err := h.service.GetTaskByUserId(useridstr,status,sortby,order,cursor,limitstr,pagenostr)
		if err!=nil{
			log.Println("error in service function call",err)
		}
		json.NewEncoder(writer).Encode(map[string]interface{}{
			"message":"the tasks of the users are as follows",
			"tasks":tasks,
		})
}