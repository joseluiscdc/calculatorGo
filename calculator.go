package calculator

import (
	"bufio"
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

func (c Calc) operate(operation string) int {
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

func (c Calc) processResult(input string, operator string) (string, int) {
	cleanInput := strings.Split(input, operator)

	first, err := c.parseString(cleanInput[0])
	second, err := c.parseString(cleanInput[1])

	if err != nil {
		return err.Error(), 0
	} else {
		c.operatorOne = first
		c.operatorTwo = second
		value := c.operate(operator)
		return "Result of " + input + " equals to : ", value
	}
}

func (c Calc) CalculateInput(input string) (string, int) {
	match, _ := regexp.MatchString("[0-9]([/]|[*]|[-]|[+])[0-9]", input)
	r, _ := regexp.Compile("[0-9]([/]|[*]|[-]|[+]+)[0-9]")

	if match && len(r.FindStringSubmatch(input)) > 1 {
		operator := r.FindStringSubmatch(input)[1]
		return c.processResult(input, operator)
	} else {
		return "Input is not valid!", 0
	}
}

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
