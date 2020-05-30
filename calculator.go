package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Calc struct {
	operatorOne int
	operatorTwo int
}

func (Calc) parseString(operator string) (int, error) {
	result, err := strconv.Atoi(operator)
	return result, err
}

func (c Calc) getOperators() {
	log.Println(c.operatorOne, c.operatorTwo, "Op1 = %s \n Op2 = %s")
}

func (c Calc) Operate(operation string) int {
	switch operation {
	case "+":
		return c.operatorOne + c.operatorTwo
	case "-":
		return c.operatorOne - c.operatorTwo
	case "*":
		return c.operatorOne * c.operatorTwo
	case "/":
		return c.operatorOne / c.operatorTwo
	default:
		log.Println(operation, " Operation is not supported!")
		return 0
	}
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func processResult(input string, operator string) {
	c := Calc{}

	cleanInput := strings.Split(input, operator)

	first, err := c.parseString(cleanInput[0])
	second, err := c.parseString(cleanInput[1])

	if err != nil {
		fmt.Println(err)
	} else {
		c.operatorOne = first
		c.operatorTwo = second
		value := c.Operate(operator)
		fmt.Println("Result of", input, " equals to : ", value)
	}
}

func validateInput(input string) (bool, string, string) {
	match, _ := regexp.MatchString("[0-9]([/]|[*]|[-]|[+])[0-9]", input)
	r, _ := regexp.Compile("[0-9]([/]|[*]|[-]|[+]+)[0-9]")

	if match && len(r.FindStringSubmatch(input)) > 1 {
		operatorVal := r.FindStringSubmatch(input)[1]
		return match, "Input valid!", operatorVal
	} else {
		return match, "Input is not valid!", ""
	}
}
