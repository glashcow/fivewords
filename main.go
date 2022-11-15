package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	out             strings.Builder
	regex           string
	contains        string
	doesntContain   string
	containsBool    bool
	notContainsBool bool
)

const (
	//path = "C:\\Users\\ewana\\Documents\\aDevAux\\GO\\src\\words\\five.txt"
	path = "five.txt"
)

func main() {
	lines := getWordList()

	for {
		fmt.Println("enter your word you cheating bastard")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		split := strings.Split(text, " ")
		reg := split[0]
		reg = strings.TrimSpace(reg)
		containsBool = false
		notContainsBool = false
		if len(split) > 1 {
			if strings.TrimSpace(split[1]) != "|" {
				containsBool = true
				contains = split[1]
				contains = strings.TrimSpace(contains)
			}
		}
		if strings.Contains(text, "|") {
			notContainsBool = true
			doesntContain = split[len(split)-1]
			doesntContain = strings.TrimSpace(doesntContain)
		}
		fmt.Println()

		out.Reset()
		for _, s := range lines {
			matched, _ := regexp.MatchString(reg, s)
			if matched && containsBool {
				for _, ch := range contains {
					if !strings.Contains(s, string(ch)) {
						matched = false
					}
				}
			}
			if matched && notContainsBool {
				for _, ch := range doesntContain {
					if strings.Contains(s, string(ch)) {
						matched = false
					}
				}
			}
			if matched {
				out.WriteString(s + "\n")
			}
		}
		fmt.Println(out.String())
	}
}

func getWordList() []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := make([]string, 0)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}
