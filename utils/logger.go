package utils

import (
	"encoding/json"
	"log"
	"os"
)

type Logger struct {
	Method     string      `json:"method"`
	Path       string      `json:"path"`
	Response   interface{} `json:"response"`
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
}

func Log(l *Logger) {
	file, err := os.OpenFile("./info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println("Error with creating / opening log file: ", err)
	}
	defer file.Close()

	logHead := "[info] "
	if l.Status != "success" {
		logHead = "[error] "
	}

	log_ := log.New(file, logHead, log.LstdFlags|log.Lmicroseconds)
	logData, _ := json.Marshal(l)

	// writing logs to file
	log_.Println(string(logData))
	// logging in the terminal
	log.Println(string(logData))
}
