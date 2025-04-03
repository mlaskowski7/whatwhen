package dal

import "whatwhen/models"

func CreateTask(task *models.Task) error {
	return Database.Create(task).Error
}

func FindAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := Database.Find(&tasks).Error
	return tasks, err
}

func FindTaskById(id uint) (*models.Task, error) {
	var task models.Task
	err := Database.First(&task, id).Error
	return &task, err
}

func DeleteTaskById(id uint) error {
	return Database.Delete(&models.Task{}, id).Error
}
