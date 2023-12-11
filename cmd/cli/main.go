package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

type Param struct {
	File  string
	Value string
}

func main() {
	flag.Parse()
	params := flag.Args()
	resultMap := map[string][]Param{}
	for _, envFile := range params {
		envMap, err := godotenv.Read(envFile)
		if err != nil {
			log.Fatal(err)
		}

		for i, val := range envMap {
			if _, ok := resultMap[val]; ok {
				// add variable if it exists
				resultMap[val] = append(resultMap[val], Param{envFile, i})
				continue
			}
			// if variable does not exist than create
			resultMap[val] = []Param{
				{envFile, i},
			}
		}
	}

	printToSTDOUT(resultMap)
}

func printToSTDOUT(mapOfVariables map[string][]Param) {
	for val, listOfKeys := range mapOfVariables {
		if len(listOfKeys) >= 2 {

			fmt.Printf("Duplicate value: %s\n", val)
			for _, el := range listOfKeys {
				fmt.Printf("File: %s Variable: %s \n", el.File, el.Value)
			}
			fmt.Println()
		}
	}
}
