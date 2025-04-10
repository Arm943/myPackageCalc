package main

import (
	"fmt"
	"os"
)

var name string
var newName string

// создание файла
func createFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
	}
	defer file.Close()
	fmt.Println("Файл создан:", name)
}

// переименовать файл
func renameFile(name, newName string) {
	err := os.Rename(name, newName)
	if err != nil {
		fmt.Println("Ошибка при переименовывании файла:", err)
	}
	fmt.Println("Новое имя файла:", newName)
}

// проверка существует ли файл
func existFile(name string) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		fmt.Println("Файл не существует")
	} else {
		fmt.Println("Файл существует")
	}
}

// удаление файла
func deleteFile(name string) {
	err := os.Remove(name)
	if err != nil {
		fmt.Println("Ошибка при удалении файла:", err)
	}
	fmt.Printf("Файл %s удален \n", name)
}

func menu() {
	for {
		var userInput string
		fmt.Println(`Выберите действие:
	1 - создать файл
	2 - переименовать файл
	3 - проверить файл
	4 - удалить файл
	`)
		fmt.Scan(&userInput)
		switch userInput {
		case "1":
			var name string
			fmt.Print("Введите название файла: ")
			fmt.Scan(&name)
			createFile(name)
		case "2":
			var name string
			var newName string
			fmt.Print("Введите старое название: ")
			fmt.Scan(&name)
			fmt.Print("Введите новое название: ")
			fmt.Scan(&newName)
			renameFile(name, newName)
		case "3":
			var name string
			fmt.Print("Введите название файла: ")
			fmt.Scan(&name)
			existFile(name)
		case "4":
			var name string
			fmt.Print("Введите название файла: ")
			fmt.Scan(&name)
			deleteFile(name)
		}
	}
}

func main() {
	menu()
}

/*
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Использование: create|rename|check|delete и нужные аргументы")
		return
	}

	command := args[1]

	switch command {
	case "create":
		if len(args) < 3 {
			fmt.Println("Укажи имя файла для создания")
			return
		}
		createFile(args[2])

	case "rename":
		if len(args) < 4 {
			fmt.Println("Укажи старое и новое имя файла")
			return
		}
		renameFile(args[2], args[3])

	case "check":
		if len(args) < 3 {
			fmt.Println("Укажи имя файла для проверки")
			return
		}
		existFile(args[2])

	case "delete":
		if len(args) < 3 {
			fmt.Println("Укажи имя файла для удаления")
			return
		}
		deleteFile(args[2])

	default:
		fmt.Println("Неизвестная команда. Используй: create, rename, check, delete")
	}

}
*/
