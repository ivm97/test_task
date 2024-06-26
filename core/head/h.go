package head

import (
	"log"
	"os"
	"test_task/core/queue"
	"test_task/models"
)

type application struct {
	iLog *log.Logger
	eLog *log.Logger
	task *[]models.Action
}

func Entry() {
	//Простой настроенный логгер, можно было бы сделать вывод в файл
	//или же в брокер сообщений, так же можно просто писать в базу данных выделив под это
	//сервис сбора ошибок

	iLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	eLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//Структура приложения, если бы была работа с базой данных
	//то можно было бы указать пул подключения
	//а так же кеш приложения
	app := application{
		iLog: iLog,
		eLog: eLog,
		task: queue.Get("task/task.json"),
	}

	app.start()

}
