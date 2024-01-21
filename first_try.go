package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var rim_table = [][]string{
	{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
	{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
	{"", "C"},
}

func arab_to_rim(arab int) string {
	rim := ""
	digit := 100
	if arab <= 0 || arab > 100 {
		panic(1)
	}

	for i := 2; i >= 0; i-- {
		j := arab / digit
		rim += rim_table[i][j]
		arab %= digit
		digit /= 10
	}

	return rim
}

func rim_to_arab(rim string) int {
	arab := 0
	rex := regexp.MustCompile(`^(C{0,1})(L?X{0,3}|X[LC])(V?I{0,3}|I[VX])$`)

	if !rex.MatchString(rim) {
		panic(1)
	}

	digits := rex.FindAllStringSubmatch(rim, 1)
	digit := 100

	for i := 1; i < len(digits[0]); i++ {
		for k, v := range rim_table[3-i] {
			if digits[0][i] == v {
				arab += k * digit
				break
			}
		}

		digit /= 10
	}

	return arab
}

func add(a int, b int) int {
	return a + b
}

func diff(a int, b int) int {
	return a - b
}

func mult(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}

func main_action_arabic(a int, b int, action string) int {
	switch action {
	case "+":
		return (add(a, b))
	case "-":
		return (diff(a, b))
	case "*":
		return (mult(a, b))
	case "/":
		return (div(a, b))
	default:
		panic(1)
	}
}

func get_number(number string) (int, bool) {
	is_rim := false
	a, err := strconv.Atoi(number)
	if err != nil {
		a = rim_to_arab(number)
		is_rim = true
	}
	if a > 10 || a < 0 {
		panic(1)
	}
	return a, is_rim
}

func main() {
	scanner := bufio.NewScanner((os.Stdin))
	scanner.Scan()
	example_splited := strings.Split(scanner.Text(), " ")

	var pred_ []string
	for i := 0; i < len(example_splited); i++ {
		if strings.Compare(example_splited[i], "") != 0 {
			pred_ = append(pred_, example_splited[i])
		}
	}

	if len(pred_) != 3 {
		panic(1)
	}

	a, is_rim_a := get_number(pred_[0])
	b, is_rim_b := get_number(pred_[2])

	res := main_action_arabic(a, b, pred_[1])

	if is_rim_a == is_rim_b {

		if is_rim_a == true && is_rim_b == true {
			fmt.Println(arab_to_rim(res))
		}
		if is_rim_a == false && is_rim_b == false {
			fmt.Println(res)
		}
	} else {
		panic(1)
	}
}
