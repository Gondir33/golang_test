package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputRoma = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

var outputRoma = map[int]string{
	1:   "I",
	2:   "II",
	3:   "III",
	4:   "IV",
	5:   "V",
	6:   "VI",
	7:   "VII",
	8:   "VIII",
	9:   "IX",
	10:  "X",
	11:  "XI",
	12:  "XII",
	13:  "XIII",
	14:  "XIV",
	15:  "XV",
	16:  "XVI",
	17:  "XVII",
	18:  "XVIII",
	19:  "XIX",
	20:  "XX",
	21:  "XXI",
	24:  "XXIV",
	25:  "XXV",
	27:  "XXVII",
	28:  "XXVIII",
	30:  "XXX",
	32:  "XXXII",
	35:  "XXXV",
	36:  "XXXVI",
	40:  "XL",
	42:  "XLII",
	45:  "XLV",
	48:  "XLVIII",
	49:  "XLIX",
	50:  "L",
	54:  "LIV",
	56:  "LVI",
	60:  "LX",
	63:  "LXIII",
	64:  "LXIV",
	70:  "LXX",
	72:  "LXXII",
	80:  "LXXX",
	81:  "LXXXI",
	90:  "XC",
	100: "C",
}

func check_numbers(text string, sep string) (bool, error) {
	var args []string
	var roma bool
	roma = false
	args = strings.Split(text, sep)
	_, fir := inputRoma[args[0]]
	_, sec := inputRoma[args[1]]
	if fir == true && sec == true {
		roma = true
	} else if (fir == true && sec == false) || (fir == false && sec == true) {
		return roma, errors.New("Вывод ошибки, так как используются одновременно разные системы счисления")
	} else {
		arg1, fir := strconv.Atoi(args[0])
		arg2, sec := strconv.Atoi(args[1])
		if fir == nil && sec == nil {
			if (arg1 < 1 || arg1 > 10) || (arg2 < 1 || arg2 > 10) {
				return roma, errors.New("Неправильный формат введения данных должны быть числа от 1 до 10")
			}
		} else {
			return roma, errors.New("Неправильный формат введения данных должны быть числа от 1 до 10 или римские цифры от I до X")
		}
	}
	return roma, nil
}

func plus(args []string, roma bool) (int, error) {
	var args1, args2 int

	if roma == true {
		args1, args2 = inputRoma[args[0]], inputRoma[args[1]]
	} else {
		args1, _ = strconv.Atoi(args[0])
		args2, _ = strconv.Atoi(args[1])
	}
	return args1 + args2, nil
}

func minus(args []string, roma bool) (int, error) {
	var args1, args2 int

	if roma == true {
		args1, args2 = inputRoma[args[0]], inputRoma[args[1]]
	} else {
		args1, _ = strconv.Atoi(args[0])
		args2, _ = strconv.Atoi(args[1])
	}
	if roma == true && (args1-args2) <= 0 {
		return 0, errors.New("Вывод ошибки, так как в римской системе нет отрицательных чисел и нуля.")
	}
	return args1 - args2, nil
}

func divide(args []string, roma bool) (int, error) {
	var args1, args2 int

	if roma == true {
		args1, args2 = inputRoma[args[0]], inputRoma[args[1]]
	} else {
		args1, _ = strconv.Atoi(args[0])
		args2, _ = strconv.Atoi(args[1])
	}
	if roma == true && (args1/args2) <= 0 {
		return 0, errors.New("Вывод ошибки, так как в римской системе нет отрицательных чисел и нуля.")
	}
	return args1 / args2, nil
}

func multiply(args []string, roma bool) (int, error) {
	var args1, args2 int

	if roma == true {
		args1, args2 = inputRoma[args[0]], inputRoma[args[1]]
	} else {
		args1, _ = strconv.Atoi(args[0])
		args2, _ = strconv.Atoi(args[1])
	}
	return args1 * args2, nil
}
func main() {

	var res int

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, " ", "")
	i := 0
	for i < len(text) && text[i] != '+' && text[i] != '-' && text[i] != '*' && text[i] != '/' {
		i++
	}
	if i == len(text) {
		fmt.Println(errors.New("Вывод ошибки, так как строка не является математической операцией"))
		return
	}
	j := i + 1
	for j < len(text) && text[j] != '+' && text[j] != '-' && text[j] != '*' && text[j] != '/' {
		j++
	}
	if j < len(text) {
		fmt.Println(errors.New("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."))
		return
	}
	if i < len(text) {
		roma, err := check_numbers(text, string(text[i]))
		if err != nil {
			fmt.Println(err)
			return
		}
		if text[i] == '+' {
			res, err = plus(strings.Split(text, "+"), roma)
		} else if text[i] == '-' {
			res, err = minus(strings.Split(text, "-"), roma)
		} else if text[i] == '*' {
			res, err = multiply(strings.Split(text, "*"), roma)
		} else {
			res, err = divide(strings.Split(text, "/"), roma)
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		if roma == true {
			fmt.Println(outputRoma[res])
		} else {
			fmt.Println(res)
		}

	}
}
