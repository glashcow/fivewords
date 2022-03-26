package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	out   strings.Builder
	regex string
)

const (
	path = "C:\\Users\\ewana\\Documents\\aDevAux\\GO\\src\\words\\five.txt"
)

func main() {
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

	for {
		fmt.Println("enter your word you prick")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		split := strings.Split(text, " ")
		reg := split[0]
		reg = strings.TrimSpace(reg)
		conBool := false
		var contains []string
		if len(split) > 1 {
			contains = split[1:]
			conBool = true
		}
		fmt.Println()

		out.Reset()
		for _, s := range lines {
			matched, _ := regexp.MatchString(reg, s)
			if matched && conBool {
				containsBool := true
				for _, c := range contains {
					c = strings.TrimSpace(c)
					if !strings.Contains(s, c) {
						containsBool = false
					}
				}
				if containsBool {
					out.WriteString(s + "\n")
				}

			} else if matched {
				out.WriteString(s + "\n")
			}
		}
		fmt.Println(out.String())
	}
}
