package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	// --- Налаштування ---
	fileName := "dummy_file.txt" // Назва файлу, який буде змінюватись
	numCommits := 100            // Кількість комітів, яку потрібно створити
	// ---------------------

	// Перевіряємо, чи ініціалізовано репозиторій Git
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		fmt.Println("Помилка: Поточна папка не є Git-репозиторієм.")
		fmt.Println("Будь ласка, виконайте 'git init' перед запуском скрипта.")
		return
	}

	// Створюємо файл, якщо він не існує
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		err := ioutil.WriteFile(fileName, []byte("a"), 0644)
		if err != nil {
			fmt.Printf("Не вдалося створити файл: %v\n", err)
			return
		}
	}

	fmt.Printf("Починаю створювати %d комітів...\n", numCommits)

	for i := 1; i <= numCommits; i++ {
		// Читаємо файл
		content, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Printf("Помилка читання файлу: %v\n", err)
			return
		}

		// Змінюємо останній символ
		if len(content) > 0 {
			lastChar := content[len(content)-1]
			content[len(content)-1] = lastChar + 1 // Просто збільшуємо ASCII-код
		} else {
			content = []byte("a") // Якщо файл порожній
		}

		// Записуємо зміни назад у файл
		err = ioutil.WriteFile(fileName, content, 0644)
		if err != nil {
			fmt.Printf("Помилка запису у файл: %v\n", err)
			return
		}

		// Виконуємо git add
		cmdAdd := exec.Command("git", "add", ".")
		if err := cmdAdd.Run(); err != nil {
			fmt.Printf("Помилка виконання 'git add': %v\n", err)
			return
		}

		// Створюємо коміт
		commitMessage := fmt.Sprintf("Автоматичний коміт #%d", i)
		cmdCommit := exec.Command("git", "commit", "-m", commitMessage)
		if err := cmdCommit.Run(); err != nil {
			fmt.Printf("Помилка виконання 'git commit': %v\n", err)
			return
		}

		fmt.Printf("Створено коміт %d/%d: %s\n", i, numCommits, commitMessage)
	}

	fmt.Println("\nЗавершено! Успішно створено", numCommits, "комітів.")
}
