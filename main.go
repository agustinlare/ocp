package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// Context - stop asking me for the comment
type Context struct {
	Name     string `yaml:"name"`
	URL      string `yaml:"url"`
	Contexto string `yaml:"contexto"`
}

// Env - stop aspking me for the the comment
type Env struct {
	Username string    `yaml:"username"`
	Password string    `yaml:"password"`
	Context  []Context `yaml:"env"`
}

func main() {
	var argsLen int = len(os.Args)
	var config Env

	config = getConfig()

	var resp []byte
	n := "default"

	switch argsLen {
	case 1:
		resp = loginCluster(selectCluster(config), n)
	case 2:
		resp = loginCluster(getURL(config), os.Args[1])
	case 3:
		if os.Args[1] == "-n" {
			resp = ocCaller([]string{"project", os.Args[2]})
		} else {
			resp = loginCluster(getURL(config), os.Args[2])
		}
		// default:
		// 	var args []string

		// 	for _, v := range os.Args[1:] {
		// 		args = append(args, v)
		// 	}

		// 	resp = ocCaller(args)
	}

	fmt.Println(string(resp))
}

func loginCluster(c string, n string) []byte {
	var namespace string = "--namespace=" + n
	var server = "--server=" + c

	username := "--username=" + getUsername()
	password := "--password=" + decrypt()

	resp := ocCaller([]string{"login", username, password, server, "--insecure-skip-tls-verify=true", namespace})

	return resp
}

func selectCluster(config Env) string {
	var o int
	var separator string = "-----------------------------------------------------------"
	var header string = "N | NAME\t| URL"

	fmt.Println("Ingrese opcion:\n" + header + "\n" + separator)

	for i, v := range config.Context {
		fmt.Println(i, "|", v.Name, "\t|", v.URL)
	}

	fmt.Println(separator)
	fmt.Scan(&o)

	return config.Context[o].URL
}

func ocCaller(args []string) []byte {
	var stdByte []byte

	cmd := exec.Command("oc", args...)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	errByte, _ := ioutil.ReadAll(stderr)

	if len(errByte) > 0 {
		stdByte = tryThis(string(errByte))
	}

	stdByte, _ = ioutil.ReadAll(stdout)

	return stdByte
}
