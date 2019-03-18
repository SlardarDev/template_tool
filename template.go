package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"strconv"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"Title": func(args ...interface{}) string {
		if len(args) == 0 {
			return ""
		}
		arg1 := args[0]
		s := arg1.(string)
		return strings.Title(s)
	},
	"Concat": func(args ...interface{}) string {

		if len(args) == 0 {
			return ""
		}
		result := ""
		for _, arg := range args {
			result += arg.(string)
		}
		return result
	},

	"Replace": func(args ...interface{}) string {
		if len(args) != 4 {
			return ""
		}

		arg1 := args[0].(string)
		arg2 := args[1].(string)
		arg3 := args[2].(string)
		arg4 := args[3].(int)

		return strings.Replace(arg1, arg2, arg3, arg4)
	},
}

func GenTemplate(templateFile string, templateMap string) {

	arguments := map[string]interface{}{}
	err := json.Unmarshal([]byte(templateMap), &arguments)
	if err != nil {
		panic(err)
	}
	fileContent, err := ioutil.ReadFile(templateFile)
	if err != nil {
		panic(err)
	}
	tpl := template.New(templateFile)
	tpl.Funcs(funcMap)
	tpl, err = tpl.Parse(string(fileContent))
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBufferString("")

	tpl.Execute(buf, arguments)

	result := buf.String()

	if strings.Contains(result, "__LINE__") {
		result = processLineMacro(result)
	}
	fmt.Print(result)
}

func processLineMacro(s string) string {

	rs, err := format.Source([]byte(s))
	var arr []string
	if err != nil {
		arr = strings.Split(s, "\n")
	} else {
		arr = strings.Split(string(rs), "\n")
	}

	newFile := []string{}

	for idx, line := range arr {
		newFile = append(newFile, strings.Replace(line, "__LINE__", strconv.FormatInt(int64(idx)+1, 10), -1))
	}

	return strings.Join(newFile, "\n")

}
