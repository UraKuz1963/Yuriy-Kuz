package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func Contains(a string, x string) bool {
	for isimv := 0; isimv < len(a); isimv++ {
		for i := 0; i < len(x); i++ {
			if x[i] != a[isimv+i] {
				break
			}
			if i == len(x)-1 {
				return true
			}
		}
	}
	return false
}

func seekinBody(host string, strFind string, id *int) bool {
	resp, err := http.Get(host)
	if err != nil {
		fmt.Println(err)
		*id--
		return false
	}
	defer resp.Body.Close()
	strBody := ""
	for true {
		bs := make([]byte, 32768)
		n, err := resp.Body.Read(bs)
		strBody += string(bs[:n])
		if n == 0 || err != nil {
			if Contains(strBody, strFind) == true {
				fmt.Println("В хосте ", host, " найдена строка ", strFind, "\n")
				*id--
				return true
			}
			break
		}
	}
	*id--
	return false
}

func main() {
	iRut := 0
	f, err := os.Open("hello.txt")
	check(err)
	fmt.Println("Start:", "\n")
	if err != nil { // если возникла ошибка
		fmt.Println("Unable to open file:", err)
		os.Exit(1) // выходим из программы
	}
	b1 := make([]byte, 1024) //Предполагаем пока макс длинну файла 1024 байт
	n1, err := f.Read(b1)
	check(err)
	//fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	words := strings.Fields(string(b1[:n1]))
	var input string
	fmt.Println("\n", "Введите строку для поиска:")
	fmt.Scanln(&input)
	fmt.Println("\n", "Идет поиск:", "\n")
	for _, word := range words {
		iRut++
		go seekinBody(word, input, &iRut)
		for iRut > 20 { // Пауза если процессов более 4
		}
	}
	for iRut > 0 { // Ожидание окончания всех процессов
	}
	fmt.Println("End:", "\n")
	defer f.Close()                       // закрываем файл
	fmt.Println(" Имя файла: ", f.Name()) // hello.txt
	return
}
