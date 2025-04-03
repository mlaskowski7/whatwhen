package bll

import (
	"time"
	"whatwhen/dal"
	"whatwhen/models"
)

func GetAllTasks() ([]models.Task, error) {
	return dal.FindAllTasks()
}

func FindTaskById(id uint) (*models.Task, error) {
	return dal.FindTaskById(id)
}

func DeleteTaskById(id uint) error {
	return dal.DeleteTaskById(id)
}

func CreateTask(title string, desc string, deadline time.Time, finished bool) (*models.Task, error) {
	task := models.NewTask(title, desc, deadline, finished)
	err := dal.CreateTask(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}
