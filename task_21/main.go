package main

import "fmt"

// --------------------------
type OldLogger struct{}

func (l *OldLogger) Log(msg string) {
	fmt.Println("Log: ", msg)
}

//--------------------------

type Logger interface {
	LogInfo(msg string)
	LogError(msg string)
}

type LoggerAdapter struct {
	oldLogger *OldLogger
}

func (l *LoggerAdapter) LogInfo(msg string) {
	l.oldLogger.Log(msg + "Level: INFO")
}

func (l *LoggerAdapter) LogError(msg string) {
	l.oldLogger.Log(msg + "Level: ERROR")
}

func main() {
	OldLogger := &OldLogger{}
	logger := &LoggerAdapter{OldLogger}
	logger.LogInfo("Hello, World")
	logger.LogError("Smth error")

}
