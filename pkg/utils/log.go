package utils

import (
	"log"
	"os"
	"time"
)

var (
	LoggerInfo  *log.Logger
	LoggerError *log.Logger
	LoggerWarn  *log.Logger
	logFile     *os.File
	Info        = "\033[34m" // Blue
	Error       = "\033[31m" // Red
	Warn        = "\033[33m" // Yellow
	Reset       = "\033[0m"
)

const LogMaxSize = 10 * 1024

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func Init() {
	var (
		err   error
		check bool
	)
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0755)
		if err != nil {
			log.Println("Error creating log directory: " + err.Error())
			return
		}
	}

	if fileExists("logs/app.log") {
		check = true
	}

	logFile, err = os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Error opening log file: " + err.Error())
		return
	}

	LoggerInfo = log.New(logFile, "INFO: ", log.LstdFlags|log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
	LoggerError = log.New(logFile, "ERROR: ", log.LstdFlags|log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
	LoggerWarn = log.New(logFile, "WARNING: ", log.LstdFlags|log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)

	if !check {
		Welcome()
	}
}

func Rotate() {
	if stat, err := logFile.Stat(); err == nil {
		if stat.Size() > LogMaxSize {
			err := logFile.Close()
			if err != nil {
				log.Println("Error closing log file: " + err.Error())
				return
			}

			newName := "logs/app-" + time.Now().Format("2006-01-02_15-04-05") + ".log"
			err = os.Rename("logs/app.log", newName)
			if err != nil {
				log.Println("Error renaming log file: " + err.Error())
				return
			}
			Init()
		}
	}
}

func CleanUp() {
	if logFile != nil {
		Close()
		err := logFile.Close()
		if err != nil {
			LoggerError.Println("Error closing file: " + err.Error() + Reset)
			return
		}
	}
}

func Welcome() {
	LoggerInfo.Println(Info + "" + Reset)
	LoggerInfo.Println(Info + " ▗▄▄▖ ▗▄▖      ▗▄▄▖▗▄▄▄▖▗▄▖ ▗▄▄▖▗▄▄▄▖▗▄▄▄▖▗▄▄▖ " + Reset)
	LoggerInfo.Println(Info + "▐▌   ▐▌ ▐▌    ▐▌     █ ▐▌ ▐▌▐▌ ▐▌ █  ▐▌   ▐▌ ▐▌" + Reset)
	LoggerInfo.Println(Info + "▐▌▝▜▌▐▌ ▐▌     ▝▀▚▖  █ ▐▛▀▜▌▐▛▀▚▖ █  ▐▛▀▀▘▐▛▀▚▖" + Reset)
	LoggerInfo.Println(Info + "▝▚▄▞▘▝▚▄▞▘    ▗▄▄▞▘  █ ▐▌ ▐▌▐▌ ▐▌ █  ▐▙▄▄▖▐▌ ▐▌" + Reset)
	LoggerInfo.Println(Info + "" + Reset)
}

func Close() {
	LoggerInfo.Println(Info + "Closing logger..." + Reset)
}
