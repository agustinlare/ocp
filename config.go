package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"

	"gopkg.in/yaml.v2"
)

// Config oc config view --minify
type Config struct {
	Clusters []Clusters `yaml:"clusters"`
}

// Clusters {range .clusters[*]}
type Clusters struct {
	Name    string            `yaml:"name"`
	Cluster map[string]string `yaml:"cluster"`
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func getConfig() Env {
	var c Env
	yamlFile, err := ioutil.ReadFile(getConfigfile())

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &c)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func getContext() string {
	// var jsonParse string = "-o=jsonpath='{range .clusters[*]}{.cluster}'"
	var jsonParse string = "-o=json"

	cmd := exec.Command("oc", "config", "view", "--minify", jsonParse)

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	errByte, _ := ioutil.ReadAll(stderr)

	if len(errByte) > 0 {
		panic(string(errByte))
	}

	stdByte, _ := ioutil.ReadAll(stdout)
	var m Config
	err := yaml.Unmarshal(stdByte, &m)

	if err != nil {
		panic(err)
	}

	return m.Clusters[0].Cluster["server"]
}

func getConfigfile() string {
	var configFile string

	if isLinux() {
		configFile = "/etc/occonfig"
	} else {
		configFile = os.Getenv("USERPROFILE") + "\\.kube\\occonfig"
	}

	if !fileExists(configFile) {
		fmt.Println(configFile)
		panic("No existe archivo de configuracion.")
	}

	return configFile
}

func getURL(config Env) string {
	var url string

	for _, v := range config.Context {
		if v.Name == os.Args[1] {
			url = v.URL
			break
		}
	}

	if url == "" {
		panic("Contexto no encontrado")
	}

	return url
}

func isLinux() bool {
	var b bool = false

	if runtime.GOOS != "windows" {
		b = true
	}

	return b
}
