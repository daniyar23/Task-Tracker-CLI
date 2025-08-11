package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

func ReadAllTasks(tasks []Task) {
	for _, task := range tasks {
		fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s\n", task.ID, task.Description, task.Status, task.CreatedAt.Format("2006-01-02"))
	}
}
func Update(task []Task, id int, status string) {
	flag := false
	for i := range task {

		if task[i].ID == id {
			task[i].Status = status
			task[i].UpdatedAt = time.Now()
			flag = true
			break
		}
	}
	if !flag {
		fmt.Println("Задача с ID", id, "не найдена")
		return
	}
	jsonData, err := json.MarshalIndent(task, "", " ")
	if err != nil {
		fmt.Println("Ошибка при сериализации JSON:", err)
		return
	}
	err = os.WriteFile("tasks.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}
}

func Delete(tasks []Task, id int) []Task {
	flag := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			flag = true
			break
		}
	}
	if !flag {
		fmt.Println("Задача с ID", id, "не найдена")
		return tasks
	}
	jsonData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Ошибка при сериализации JSON:", err)
		return tasks
	}
	err = os.WriteFile("tasks.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return tasks
	}
	return tasks
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Пожалуйста,укажите команду: add,list,update,delete")
		return
	}
	command := os.Args[1]
	tasks := JsonCheck()
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Пожалуйста,укажите описание задачи")
			return
		}
		description := os.Args[2]
		tasks, _ = Create(description, tasks)
		fmt.Println("Задача успешно добавлена")

	case "list":
		ReadAllTasks(tasks)
	case "updete":
		if len(os.Args) < 4 {
			fmt.Println("Пожалуйста,укажите ID и новое описание задачи")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("неверный id:", os.Args[2])

		}
		status := os.Args[3]
		Update(tasks, id, status)
		fmt.Println("Задача успешно обновлена")
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Пожалуйста, укажите ID задачи для удаления")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Неверный ID:", os.Args[2])
			return
		}
		tasks = Delete(tasks, id)
		fmt.Println("Задача удалена")

	default:
		fmt.Println("Неизвестная команда:", command)
	}
}
