package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	var sqlWithParams string

	fmt.Println("Enter your sql with params:")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input := scanner.Text()
		sqlWithParams = input
	}

	fmt.Println("\n Got your params, thanks", sqlWithParams)
	fmt.Print("\n\n\n\n")

	paramsMarker := "-- PARAMETERS: "
	paramsStart := strings.Index(sqlWithParams, paramsMarker)
	paramsString := sqlWithParams[paramsStart+len(paramsMarker):]

	var params []interface{}

	err := json.Unmarshal([]byte(paramsString), &params)
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	fmt.Println("Params found", params)

	sqlWithParams = sqlWithParams[:paramsStart]

	for i, param := range params {
		if parsed, ok := param.(string); ok {
			params[i] = `'` + parsed + `'`
		}

		sqlWithParams = strings.Replace(sqlWithParams, fmt.Sprintf("$%d", i+1), fmt.Sprint(params[i]), -1)
	}

	fmt.Print("\n\n", sqlWithParams, "\n\n")
}
