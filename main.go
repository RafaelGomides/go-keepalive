package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

var keepCMD, logPATH string

func term() string {
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("pidof %v", keepCMD))
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("[ LOG ] EXECUCAO PARADA EM: %v - [ ERRO ] %v \n", time.Now().Format("Mon Jan _2 15:04:05 2006"), err)
	}
	return out.String()
}

func keepalive() {
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("nohup %v >  %v 2>&1 &", keepCMD, logPATH))
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("[ LOG ] PROBLEMAS AO REINICAR EM: %v - [ ERRO ] %v \n", time.Now().Format("Mon Jan _2 15:04:05 2006"), err)
	}
}

func main() {
	arg := os.Args[1:]
	if len(arg) < 2 {
		log.Fatal(`Utilização: ./go-keepalive <COMMAND_NAME> <PATH_TO_LOG_FILE>`)
	}
	keepCMD = arg[0]
	logPATH = arg[1]
	for {
		time.Sleep(time.Millisecond * 500)
		pid := term()
		if pid == "" {
			keepalive()
		}
	}
}
