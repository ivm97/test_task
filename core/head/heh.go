package head

import (
	"test_task/core/queue"
	"test_task/core/worker"
	"test_task/models"
)

// Функция запускающая главный процесс
func (app *application) start() {
	var result []models.Action
	for _, e := range *app.task {
		app.iLog.Printf("Работаю с заданием: %s\n", e.Title)
		var work worker.WorkQueue
		work.Task = &e
		field := work.Working()
		app.iLog.Printf("Результат выполнения: %s\n", field.Result)
		result = append(result, *field)
	}

	queue.ToFile("task/task.json", &result)

}
