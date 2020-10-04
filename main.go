package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const messages string = "messages.txt"

// Screen to print different screens
func Screen(n string) {
	switch n {
	case "1":
		fmt.Println("***************************")
		fmt.Println("**.:|              8:48am**")
		fmt.Println("**                       **")
		fmt.Println("**     Adrián Olmedo     **")
		fmt.Println("**                       **")
		fmt.Println("**        [Menu]         **")
		fmt.Println("***************************")
	case "2":
		fmt.Println("***************************")
		fmt.Println("**.:|              8:48am**")
		fmt.Println("**                       **")
		fmt.Println("**[1] Messages.          **")
		fmt.Println("**[2] Back.              **")
		fmt.Println("**[3] Turn off.          **")
		fmt.Println("***************************")
	case "3":
		fmt.Println("***************************")
		fmt.Println("**                       **")
		fmt.Println("**                       **")
		fmt.Println("**     Turning off...    **")
		fmt.Println("**                       **")
		fmt.Println("**                       **")
		fmt.Println("***************************")
	}
}

// Xor Encryption (Exclusive-OR encryption)
func Xor(line string) (output string) {
	key := "K"

	for i := 0; i < len(line); i++ {
		output += string(line[i] ^ key[i%len(key)])
	}

	return output
}

// ClearScreen clear screen in terminal
func ClearScreen() {
	fmt.Println("\033[H\033[2J")
}

// Menu is a main menu of options
func Menu() {
	var intro, opc string

	for {
		ClearScreen() // Clear screen
		Screen("1")
		_, err := fmt.Scan(&intro)

		for {
			ClearScreen()
			Screen("2")
			_, err := fmt.Scan(&opc)

			switch opc {
			case "1":
				Chat()
			case "2":
				Menu()
			case "3":
				ClearScreen()
				Screen("3")
				time.Sleep(2000 * time.Millisecond)
				os.Exit(3)
			}

			if err != nil && opc < "1" || opc > "3" {
				break
			}
		}

		if err != nil {
			break
		}
	}
}

// SendMessage write ciphered messages with Xor in messages.txt file, line by line
func SendMessage() {
	file, err := os.OpenFile(messages, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}

	// Se pone aqui un defer para asegurarnos
	// de una buena vez que se cierre el archivo al final
	defer file.Close()

	fmt.Println("Send message:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	msg := scanner.Text()
	file.WriteString(Xor(msg) + "\n")
}

// read line by line into memory
// all file contents is stores in lines[]
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Decipher open file for reading, read line by line and print file contents
func Decipher() {
	lines, err := readLines(messages)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for i, line := range lines {
		fmt.Println(i, Xor(line))
	}
}

// Chat is menu to manage SendMessage()
func Chat() {
	var resp string
	ClearScreen()
	Decipher()
	for {
		fmt.Println("Message? (y/n) r to update")
		_, err := fmt.Scan(&resp)

		if resp == "y" || resp == "Y" {
			ClearScreen()
			Decipher()
			SendMessage()
		}

		if resp == "r" || resp == "R" {
			Chat()
		}

		if err != nil && resp == "n" || resp == "N" {
			break
		}
	}
}

func main() {
	Menu()
}
