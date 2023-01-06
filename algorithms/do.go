package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

type QuestionModel struct {
	Question string `json:"question"`
	A        int    `json:"a"`
	B        int    `json:"b"`
	C        int    `json:"c"`
}

var models = make(map[string]QuestionModel)
var (
	a      int
	b      int
	answer int
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var i = 0

	for i < 50 {
		a = randomString(0, 100)
		b = randomString(0, 50)

		question := fmt.Sprintf("%v/%v", a, b)

		if b != 0 && a%b == 0 {
			answer = a / b
			models[strconv.Itoa(i+1)] = QuestionModel{
				Question: question,
				A:        answer,
				B:        answer + randomString(0, 50),
				C:        answer - randomString(0, 5),
			}
			i++
		}

	}
	file, _ := json.MarshalIndent(models, "", " ")
	_ = ioutil.WriteFile("easy/subtraction-easy.json", file, 0644)

}

func randomString(min int, max int) int {
	return min + rand.Intn(max-min)
}
