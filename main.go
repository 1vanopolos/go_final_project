package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

const (
	defaultPort = "7540"
	webDir      = "./web"
)

func main() {
	// Определяем порт, который будет слушать сервер
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = defaultPort
	}

	// Настраиваем файловый сервер для обслуживания статических файлов из директории web
	fs := http.FileServer(http.Dir(webDir))
	http.Handle("/", fs)

	//Проверяю есть ли файл базы

	appPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dbFile := filepath.Join(filepath.Dir(appPath), "scheduler.db")
	_, err = os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}

	// Открываем базу данных
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if install {
		// Создаем таблицу и индексы
		database.CreateDB(db)
	} else {
		fmt.Println("База данных уже существует")
	}
	// если install равен true, после открытия БД требуется выполнить
	// sql-запрос с CREATE TABLE и CREATE INDEX

	// Запускаем сервер
	log.Printf("Сервер запущен на порту %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}

}
