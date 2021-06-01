// Go program to illustrate how to
// access the bytes of the string
package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func hostDial(host string, id int) bool {
	port := "80"
	timeout := time.Duration(4 * time.Second)
	//_, err := net.Dial("tcp", host)
	//_, err := net.DialTimeout("tcp", host+":"+port, timeout)
	_, err := net.DialTimeout("tcp", host, timeout)
	if err != nil {
		fmt.Printf("%s %s %s\n", host, "not responding", err.Error())
		return false
	} else {
		fmt.Printf("%s %s %s\n", host, "responding ....", port)
		return true
	}
}

func main() {

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
	for idx, word := range words {
		//timeout := time.Duration(4 * time.Second)
		//_, err := net.Dial("tcp", host)
		//_, err := net.DialTimeout("tcp", host+":"+port, timeout)
		//_, err := net.DialTimeout("tcp", host, timeout)
		go hostDial(word, idx)
		//if hostDial(word, 4) == false {
		//fmt.Printf("%s %s %s\n", host, "not responding", err.Error())
		//} else {
		//fmt.Printf("%s %s %s\n", host, "responding ....", port)
		//}
	}
	var input string
	fmt.Scanln(&input)                    // Для паузы перед выходом из программы
	defer f.Close()                       // закрываем файл
	fmt.Println(" Имя файла: ", f.Name()) // hello.txt
}
