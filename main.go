package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func JsonCheck() []Task {
	var tasks []Task
	_, err := os.Stat("tasks.json")
	if os.IsNotExist(err) {
		fmt.Println("tasks.json не существует,создаем...")
		file, err := os.Create("tasks.json")
		if err != nil {
			fmt.Println("Ошибка при создании файла:", err)
			return tasks
		}
		defer file.Close()
		_, err = file.WriteString("[]")
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
		}
		fmt.Println("Файл успешно создан")
	} else if err != nil {
		fmt.Println("Ошибка при проверке файла:", err)
		return tasks
	} else {
		fmt.Println("Файл существует,читаем...")
		data, err := os.ReadFile("tasks.json")
		if err != nil {
			fmt.Println("Ошибка при чтении файла", err)
			return tasks
		}
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			fmt.Println("Ошибка при десериализации JSON:", err)
			return tasks
		}
		fmt.Println("Загружено задач:", len(tasks))
	}
	return tasks
}

func Create(description string, tasks []Task) ([]Task, error) {
	tasks = append(tasks, Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	jsonData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Ошибка при сериализации JSON:", err)
		return tasks, err
	}
	err = os.WriteFile("tasks.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return tasks, err
	}

	return tasks, nil
}
