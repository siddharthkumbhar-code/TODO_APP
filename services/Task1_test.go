package services

import (
	"testing"

	"go-sqlite/repository"
	"go-sqlite/testutils"
)

func Getservice() *TaskServices {
	db := testutils.SetupTestDb()
	repo := repository.NewTaskRepository(db)
	service := NewTaskServices(repo)
	return service

}

func TestGetTaskByUserId_Success(t *testing.T) {

	service := Getservice()
	tasks, err := service.GetTaskByUserId(
		"1", // userid
		"",  // status
		"",  // sortby
		"",  // order
		"",  // cursor
		"",  // limit
		"",  // page
	)
	// 5. Assertions
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(tasks) == 0 {
		t.Fatalf("expected tasks, got empty")
	}
}

func TestGetTaskByUserId_Withstatus(t *testing.T) {
	service := Getservice()
	tasks, err := service.GetTaskByUserId(
		"1",
		"pending", "", "", "", "", "",
	)
	if err != nil {
		t.Fatalf("unexpectd error:%v", err)
	}
	if len(tasks) == 0 {
		t.Fatalf("expected task, got empty:")
	}
	for _, task := range tasks {
		if task.Status != "pending" {
			t.Errorf("status is not pending: %v", task.Status)
		}
	}
}

func TestGetTaskByUserId_Pagination(t *testing.T) {
	tasks, err := Getservice().GetTaskByUserId(
		"1",
		"", "", "", "",
		"1", "1",
	)
	if err != nil {
		t.Fatalf("unexpected error:%v", err)
	}
	if len(tasks) == 0 {
		t.Fatalf("expected tasks,got empty")
	}
}

func TestGetTaskByUserId_invaliduser(t *testing.T) {
	_, err := Getservice().GetTaskByUserId(
		"abc",
		"", "", "", "",
		"", "",
	)
	if err == nil {
		t.Fatalf("expected error for invalid user:%v", err)
	}
}

func TestGetTaskByUserId_emptyuser(t *testing.T) {
	_, err := Getservice().GetTaskByUserId(
		"",
		"", "", "", "",
		"", "",
	)
	if err == nil {
		t.Fatalf("expected error for invalid user:%v", err)
	}

}

func TestGetTaskByUserId_Nodata(t *testing.T) {
	tasks, err := Getservice().GetTaskByUserId(
		"999",
		"", "", "", "", "", "",
	)
	if err != nil {
		t.Fatalf("unexpected error :%v", err)
	}

	if len(tasks) != 0 {
		t.Errorf("expecting 0 tasks but got :%d", len(tasks))
	}
}
