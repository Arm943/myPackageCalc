package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// v2.1.1 версия программы

var name string
var newName string

// ФУНКЦИИ ДЛЯ РАБОТЫ С ФАЙЛАМИ
// создание файла
func createFile(name string) {
	file, err := os.Create(name)
	if err != nil {
		logErr("Ошибка при создании файла", err)
		return
	}
	defer file.Close()
	logInfoMsg("✅ Файл создан:", name)
}

// переименовать файл
func renameFile(name, newName string) {
	err := os.Rename(name, newName)
	if err != nil {
		logErr("Ошибка при переименовывании файла:", err)
		return
	}
	logInfoMsg("✅ Новое имя файла:", newName)
}

// проверка существует ли файл
func existFile(name string) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		fmt.Println("Файл не существует", name)
		logInfoMsg("Файл не существует", name)
	} else {
		logInfoMsg("✅ Файл существует", name)
	}
}

// удаление файла
func deleteFile(name string) {
	err := os.Remove(name)
	if err != nil {
		logErr("Ошибка при удалении файла:", err)
		return
	}
	logInfoMsg("✅ Файл %s удален \n", name)
}

// ФУНКЦИИ ДЛЯ РАБОТЫ С СОДЕРЖИМЫМ ФАЙЛОВ

// запись текста в файл
func writeNewText(fileName string, userText string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		logErr("Ошибка отрытия файла:", err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte(userText))
	if err != nil {
		logErr("Ошибка записи:", err)
	}
	logInfoMsg("✅ Ваш текст успешно сохранен в ", fileName)

}

// чтение текста в файле
func readFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		logErr("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil && err != io.EOF {
		logErr("Ошибка чтения:", err)
		return
	}
	fmt.Println("✅ Содержимое файла:")
	fmt.Println(string(buf[:n]))
}

// копирование текста из одно файла в другой
func copyText(fileOne, fileTwo string) {
	fOne, err := os.Open(fileOne)
	if err != nil {
		logErr("ошибка при отрытии первого файла: ", err)
		return
	}
	defer fOne.Close()

	fTwo, err := os.OpenFile(fileTwo, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		logErr("ошибка при отрытии второго файла: ", err)
		return
	}
	defer fTwo.Close()

	_, err = io.Copy(fTwo, fOne)
	if err != nil {
		logErr("Ошибка при копировании:", err)
		return
	}

	logTreeInfoMsg("✅ Все данные из файла %v успешно скопированы в файл %v ", fileOne, fileTwo)
}

// построчный вывод текста из файла
func bufScan(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		logErr("ошибка при отрытии первого файла: ", err)
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
		logErr("ошибка при отрытии первого файла: ", err)
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
		logErr("Ошибка при чтении файла: ", err)
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
		logBlackBox(string(userInput))
		switch userInput {
		case "1":
			var name string
			fmt.Print("Введите название файла: ")
			fmt.Scan(&name)
			createFile(name)
			logBlackBox(name)
		case "2":
			var name string
			var newName string
			fmt.Print("Введите старое название: ")
			fmt.Scan(&name)
			fmt.Print("Введите новое название: ")
			fmt.Scan(&newName)
			renameFile(name, newName)
			logBlackBox(newName)
		case "3":
			var name string
			fmt.Print("Введите название файла: ")
			fmt.Scan(&name)
			existFile(name)
			logBlackBox(name)
		case "4":
			var name string
			fmt.Print("Введите название файла: ")
			fmt.Scan(&name)
			deleteFile(name)
			logBlackBox(name)
		case "5":
			var fileName string
			fmt.Print("Введите название файла: ")
			fmt.Scan(&fileName)
			fmt.Scanln()
			logBlackBox(name)

			fmt.Print("Введите свой текст: ")
			reader := bufio.NewReader(os.Stdin)
			userText, _ := reader.ReadString('\n')
			userText = strings.TrimSpace(userText)
			writeNewText(fileName, userText)
			logBlackBox(userText)
		case "6":
			var fileName string
			fmt.Print("Введите название файла: ")
			fmt.Scan(&fileName)
			fmt.Println()
			readFile(fileName)
			logBlackBox(fileName)
		case "7":
			var fileOne string
			var fileTwo string

			fmt.Println("Введите название файла ОТКУДА хотите скопировать текст: ")
			fmt.Scan(&fileOne)
			logBlackBox(fileOne)
			fmt.Println("Введите название файла КУДА хотите скопировать текст: ")
			fmt.Scan(&fileTwo)
			logBlackBox(fileTwo)
			fmt.Println()
			copyText(fileOne, fileTwo)
		case "8":
			var fileName string
			fmt.Print("Введите название файла: ")
			fmt.Scan(&fileName)
			fmt.Println()
			bufScan(fileName)
			logBlackBox(fileName)
		case "9":
			var fileName string
			var text string

			fmt.Print("Введите название файла: ")
			fmt.Scan(&fileName)
			fmt.Println()
			logBlackBox(fileName)
			fmt.Print("Введите текст для поиска строки: ")
			fmt.Scan(&text)
			fmt.Println()
			logBlackBox(text)

			finder(fileName, text)
		case "0":
			os.Exit(0)
		default:
			log.Println("Выберите пункт из меню")
		}

	}
}

// отправляем логи в файл
func logSetOutput() {
	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

//варианты логов

func logErr(msg string, err error) {
	log.Println("⛔[ERROR] :", msg, err)
}
func logInfoMsg(msg, info string) {
	log.Println("[INFO] :", msg, info)
}
func logTreeInfoMsg(msg, info1, info2 string) {
	log.Println("[INFO] :", msg, info1, info2)
}
func logBlackBox(info1 string) { //users experience for analytics
	log.Println("[UX] :", info1)
}

func main() {
	logSetOutput()
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
