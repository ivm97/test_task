package worker

import (
	"fmt"
	"log"
	"os"
	"strings"
	"test_task/models"
)

type WorkQueue struct {
	Task *models.Action
}

// Функция реализующая действия над файлами
func (t *WorkQueue) Working() *models.Action {
	checkValue := strings.ToLower(t.Task.Title)
	switch {
	case strings.Contains(checkValue, "создать"):
		if _, ok := t.Task.Params["fileName"]; ok {
			err := create(t.Task.Params["fileName"])
			if err != nil {
				t.Task.Result = fmt.Sprintf("Ошибка при создании файла: %v", err)
				return t.Task
			}
			t.Task.Result = "Создан файл с именем: " + t.Task.Params["fileName"]
			return t.Task
		} else {
			t.Task.Result = "файла с таким именем не существует!"
			return t.Task
		}

	case strings.Contains(checkValue, "удалить"):
		if _, ok := t.Task.Params["fileName"]; ok {
			err := delete(t.Task.Params["fileName"])
			if err != nil {
				t.Task.Result = fmt.Sprintf("Ошибка во время удаления файла: %v", err)
				return t.Task
			}
			t.Task.Result = "Удален файл с именем: " + t.Task.Params["fileName"]
			return t.Task
		} else {
			t.Task.Result = "файла с именем: " + t.Task.Params["fileName"] + "не существует!"
			return t.Task
		}
	case strings.Contains(checkValue, "переименовать"):
		if _, ok := t.Task.Params["currentName"]; ok {
			if _, ok := t.Task.Params["newName"]; ok {
				err := rename(t.Task.Params["currentName"], t.Task.Params["newName"])
				if err != nil {
					t.Task.Result = fmt.Sprintf("Ошибка при во время смены имени файла: %v", err)
					log.Println(t.Task.Result)
					return t.Task
				}
				t.Task.Result = "Файл переименован!"
				return t.Task
			}
			t.Task.Result = "Не задан файл с новым именем для файла"
			return t.Task
		} else {
			t.Task.Result = "файла с таким именем не существует!"
			return t.Task
		}

	case strings.Contains(checkValue, "записать"):
		if _, ok := t.Task.Params["fileName"]; ok {
			inpText := []byte(t.Task.Params["content"])
			err := writeToFile(t.Task.Params["fileName"], &inpText)
			if err != nil {
				t.Task.Result = fmt.Sprintf("Ошибка при попытке записи в файл: %v", err)
				return t.Task
			} else {
				t.Task.Result = "Данные записаны!"
				return t.Task
			}
		} else {
			t.Task.Result = "файла с именем: " + t.Task.Params["fileName"] + "не существует!"
		}
	case strings.Contains(checkValue, "время"):
		if _, ok := t.Task.Params["fileName"]; ok {
			resp, err := timeInfo(t.Task.Params["fileName"])
			if err != nil {
				t.Task.Result = fmt.Sprintf("Ошибка! Не удалось получить время создания файла: %v", err)
				return t.Task
			}
			t.Task.Result = "Файл создан: " + resp
			return t.Task
		} else {
			t.Task.Result = "файла с именем: " + t.Task.Params["fileName"] + "не существует!"
			return t.Task
		}

	}
	t.Task.Result = "Такого делать не умею..."
	return t.Task

}

func create(name string) error {
	file, err := os.Create(name)
	file.Close()
	return err
}

func delete(name string) error {
	err := os.Remove(name)
	return err
}

func rename(old, new string) error {
	err := os.Rename(old, new)
	return err
}

func timeInfo(name string) (string, error) {
	info, err := os.Stat(name)
	if err != nil {
		return "", err
	}

	return info.ModTime().String(), nil

}

func writeToFile(name string, data *[]byte) error {
	err := os.WriteFile(name, *data, 0666)
	return err
}
