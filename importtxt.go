package main

import (
	"bufio"
	"os"
)

func ImportTxt() []string { //import le fichier contenant les mots
	var word []string

	readFile, _ := os.Open("hangman.txt")

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		word = append(word, fileScanner.Text())
	}

	readFile.Close()
	return word
}
