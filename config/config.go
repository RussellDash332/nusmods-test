package config

import (
	"fmt"
	"strings"
	"io/ioutil"
)

var ShowDiff = true

var modulesConfig, _ = ioutil.ReadFile("../../config/modules.txt")
var Modules []string = strings.Split(string(modulesConfig), "\n")

var semestersConfig, _ = ioutil.ReadFile("../../config/semesters.txt")
var Semesters []string = strings.Split(string(semestersConfig), "\n")

// Change package config to package main for debugging purposes
func main() {
	// Rewrite to resolve wd issue
	var modulesConfig, _ = ioutil.ReadFile("modules.txt")
	var Modules []string = strings.Split(string(modulesConfig), "\n")
	var semestersConfig, _ = ioutil.ReadFile("semesters.txt")
	var Semesters []string = strings.Split(string(semestersConfig), "\n")

	for i, mod := range Modules {
		fmt.Printf("%d => %s\n", i, mod)
	}
	fmt.Println()
	for i, sem := range Semesters {
		fmt.Printf("%d => %s\n", i, sem)
	}
}

type Scenario struct {
	Name             string
	ExpectedJsonPath string
}

var Scenarios = []Scenario{
	{
		Name:             "2018-2019",
		ExpectedJsonPath: "testdata/2018-2019.json",
	},
	{
		Name:             "2019-2020",
		ExpectedJsonPath: "testdata/2019-2020.json",
	},
	{
		Name:             "2020-2021",
		ExpectedJsonPath: "testdata/2020-2021.json",
	},
	{
		Name:             "2021-2022",
		ExpectedJsonPath: "testdata/2021-2022.json",
	},
}