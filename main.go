package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	pwdFile = "data/passwords.dat"
)

var mainSets = []string{"0123456789", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "abcdefghijklmnopqrstuvwxyz"}

/*
Создание  файла passwords.dat .
Там хранятся все предыдущие сгенерированные пароли
*/
func createFile() (*os.File, error) {
	dir := "data"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, err
	}
	
	file, err := os.Create(pwdFile)
	if err != nil {
		return nil, err
	}

	return file, nil
}

/*
	Проверка на уникальность.

по строчкам (отделенным пробелом) из passwords.dat сравниваем с сгенерированным паролем
возвращаем false , если совпадение есть
*/
func isUnique(password string) bool {
	if _, err := os.Stat(pwdFile); os.IsNotExist(err) {
		return true
	}

	file, err := os.Open(pwdFile)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == password {
			return false
		}
	}

	return true
}

/* сохранение созданного пароля в passwords.dat + переход на новую строчку */
func savePassword(password string) error {
	file, err := os.OpenFile(pwdFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(password + "\n")
	return err
}

/*создание пароля:
система такая (возможно, не самая эффективная, но простая)
создается временный слайс tSet, который изначально копия Set - слайса рун возможных символов
генерируется рандомный индекс k, берется tSet[k] , и удаляется k из tSet
далее, пока пароль не наберет l символов*/

func generatePassword(l int, set []rune) string {
	rand.Seed(time.Now().UnixNano())
	tSet := make([]rune, len(set))
	copy(tSet, set)

	var pwd string
	for i := 0; i < l; i++ {
		k := rand.Intn(len(tSet))
		pwd += string(tSet[k])
		tSet = append(tSet[:k], tSet[k+1:]...)
	}
	return pwd
}

/*
	для оптимизации и избавления от необходимости 2 раза проверять userInput

создаем слайс с индексами выбранных пользователем наборов символов
*/
func checkInput(userInput string) []int {
	var userSetsIndex []int
	if strings.Contains(userInput, "a") {
		userSetsIndex = append(userSetsIndex, 0)
	}
	if strings.Contains(userInput, "b") {
		userSetsIndex = append(userSetsIndex, 1)
	}
	if strings.Contains(userInput, "c") {
		userSetsIndex = append(userSetsIndex, 2)
	}
	if len(userSetsIndex) == 0 {
		panic("не выбран ни один набор символов")
	}

	return userSetsIndex
}

// создаем слайс рун всех возможных символов
func makeSet(setIDs []int, l int) []rune {
	var set []rune
	for _, e := range setIDs {
		set = append(set, []rune(mainSets[e])...)
	}
	if l < len(setIDs) || l > len(set) {
		fmt.Printf("длина должна быть от %d до %d\n", len(setIDs), len(set))
		panic("неверная длина")
	}

	return set
}

/*
	проверяет это правило:

Если пользователь выбрал несколько наборов символов, в пароле должен быть представлен хотя бы
один символ из каждого выбранного набора.
*/
func isCorrect(pwd string, setIDs []int) bool {
	for _, userSetIndex := range setIDs {
		if !(strings.ContainsAny(pwd, mainSets[userSetIndex])) {
			return false
		}
	}
	return true
}

func main() {
	for {

		if _, err := os.Stat(pwdFile); os.IsNotExist(err) {
			file, err := createFile()
			if err != nil {
				fmt.Println("ошибка создания файла паролей:", err)
				return
			}
			file.Close()
		}

		fmt.Println("Выбор набора символов (a, b, c)")
		fmt.Println("a - Цифры (0-9).")
		fmt.Println("b - Маленькие латинские буквы (a-z).")
		fmt.Println("c -  Большие латинские буквы (A-Z).")
		fmt.Println("Можно выбрать несколько, например: abc или bc.")
		var userInput string
		fmt.Scan(&userInput)

		setIDs := checkInput(userInput)

		fmt.Print("Введите длину пароля: ")
		var l int
		fmt.Scan(&l)

		set := makeSet(setIDs, l)

		var pwd string
		for {
			pwd = generatePassword(l, set)
			if isCorrect(pwd, setIDs) {
				if isUnique(pwd) {
					break
				}
			}
		}
		fmt.Println(pwd)

		if err := savePassword(pwd); err != nil {
			fmt.Println("Ошибка сохранения пароля:", err)
			return
		}
		fmt.Println("введите q, если хотите выйти. любой другой символ , если нет")
		var rep string
		fmt.Scan(&rep)
		if rep == "q" {
			break
		}
	}

}
