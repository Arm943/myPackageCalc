package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// v2.1.1
var name string
var newName string

// ФУНКЦИИ ДЛЯ РАБОТЫ С ФАЙЛАМИ
// создание файла
func createFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
	}
	defer file.Close()
	fmt.Println("✅ Файл создан:", name)
}

// переименовать файл
func renameFile(name, newName string) {
	err := os.Rename(name, newName)
	if err != nil {
		fmt.Println("Ошибка при переименовывании файла:", err)
	}
	fmt.Println("✅ Новое имя файла:", newName)
}

// проверка существует ли файл
func existFile(name string) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		fmt.Println("Файл не существует")
	} else {
		fmt.Println("✅ Файл существует")
	}
}

// удаление файла
func deleteFile(name string) {
	err := os.Remove(name)
	if err != nil {
		fmt.Println("Ошибка при удалении файла:", err)
	}
	fmt.Printf("✅ Файл %s удален \n", name)
}

// ФУНКЦИИ ДЛЯ РАБОТЫ С СОДЕРЖИМЫМ ФАЙЛОВ

// запись текста в файл
func writeNewText(fileName string, userText string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Ошибка отрытия файла:", err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte(userText))
	if err != nil {
		fmt.Println("Ошибка записи:", err)
	}
	fmt.Println("✅ Ваш текст успешно сохранен")

}

// чтение текста в файле
func readFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println("Ошибка чтения:", err)
		return
	}
	fmt.Println("✅ Содержимое файла:")
	fmt.Println(string(buf[:n]))
}

// копирование текста из одно файла в другой
func copyText(fileOne, fileTwo string) {
	fOne, err := os.Open(fileOne)
	if err != nil {
		fmt.Println("ошибка при отрытии первого файла: ", err)
		return
	}
	defer fOne.Close()

	fTwo, err := os.OpenFile(fileTwo, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("ошибка при отрытии второго файла: ", err)
		return
	}
	defer fTwo.Close()

	_, err = io.Copy(fTwo, fOne)
	if err != nil {
		fmt.Println("Ошибка при копировании:", err)
		return
	}

	fmt.Printf("✅ Все данные из файла %v успешно скопированы в файл %v ", fileOne, fileTwo)
}

// построчный вывод текста из файла
func bufScan(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ошибка при отрытии первого файла: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// поиск строки по слову
func finder(fileName, text string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ошибка при отрытии первого файла: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), text) {
			fmt.Println(scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла: ", err)
	}
}

func menu() {
	for {
		var userInput string
		fmt.Println()
		fmt.Println(`--------------------------------
👉 Выберите действие:
   1️⃣ - создать файл
   2️⃣ - переименовать файл
   3️⃣ - проверить файл
   4️⃣ - удалить файл
	5️⃣ - запись текста в файл
	6️⃣ - показать текст файле
	7️⃣ - скопировать текст в другой файл
	8️⃣ - построчный вывод текста из файла
	9️⃣ - найти строку по ключевому слову
--------------------------------
🅾️ - выход
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
		case "5":
			var fileName string
			fmt.Print("Введите название файла: ")
			fmt.Scan(&fileName)
			fmt.Scanln()

			fmt.Print("Введите свой текст: ")
			reader := bufio.NewReader(os.Stdin)
			userText, _ := reader.ReadString('\n')
			userText = strings.TrimSpace(userText)
			writeNewText(fileName, userText)
		case "6":
			var fileName string
			fmt.Print("Введите название файла: ")
			fmt.Scan(&fileName)
			fmt.Println()
			readFile(fileName)
		case "7":
			var fileOne string
			var fileTwo string

			fmt.Println("Введите название файла ОТКУДА хотите скопировать текст: ")
			fmt.Scan(&fileOne)
			fmt.Println("Введите название файла КУДА хотите скопировать текст: ")
			fmt.Scan(&fileTwo)
			fmt.Println()
			copyText(fileOne, fileTwo)
		case "8":
			var fileName string
			fmt.Print("Введите название файла: ")
			fmt.Scan(&fileName)
			fmt.Println()
			bufScan(fileName)
		case "9":
			var fileName string
			var text string

			fmt.Print("Введите название файла: ")
			fmt.Scan(&fileName)
			fmt.Println()
			fmt.Print("Введите текст для поиска строки: ")
			fmt.Scan(&text)
			fmt.Println()

			finder(fileName, text)
		case "0":
			os.Exit(0)
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
