package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/evaluate", PostHandler).Methods("POST")
	http.ListenAndServe(":5000", r)
}
func PostHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Expression string
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	log.Println(data)
	answer := Evaluate(data.Expression)

	w.Header().Set("Content-Type", "application/text")
	// w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500/index.html")
	// w.Header().Set("Access-Control-Request-Method", "POST")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(answer))
}

func Evaluate(input string) string {
	var answer string
	var validExp = regexp.MustCompile(`([0-9]+\.?[0-9]*)(\+|-|\*|\/)([0-9]+\.?[0-9]*)`)
	m := validExp.FindStringSubmatch(input)
	log.Println(m)
	if m == nil {
		answer = "Incorrect expression"
	} else {
		operator := m[2]
		log.Println("operator", operator)
		operand1, _ := strconv.Atoi(m[1])
		log.Println("operand1", operand1)
		operand2, _ := strconv.Atoi(m[3])
		log.Println("operand2", operand2)
		answerInt, err := MathOperator(operator, operand1, operand2)
		if err == nil {
			answer = strconv.Itoa(answerInt)
		} else {
			answer = err.Error()
		}
	}
	log.Println(answer)
	return answer
}

func MathOperator(operator string, operand1, operand2 int) (int, error) {
	switch operator {
	case "+":
		log.Println("+")
		return operand1 + operand2, nil
	case "-":
		log.Println("-")
		return operand1 - operand2, nil
	case "*":
		log.Println("*")
		return operand1 * operand2, nil
	case "/":
		if operand2 != 0 {
			log.Println("/")
			return operand1 / operand2, nil
		} else {
			err := fmt.Errorf("Can divide to 0")
			return 0, err
		}
	}
	return 0, nil
}
