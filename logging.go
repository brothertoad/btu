package btu

import (
  "log"
  "os"
)

const FATAL = 1
const ERROR = 2
const WARN  = 3
const INFO  = 4
const DEBUG = 5
const minLevel = FATAL
const maxLevel = DEBUG

var logLevel = ERROR

func SetLogLevel(level int) {
  if level < minLevel || level > maxLevel {
    log.Fatalf("log level %s is out of range (valid range is %d to %d)\n", level, minLevel, maxLevel)
  }
  logLevel = level
}

func logMsg(level int, msg string, a ...any) {
  if level <= logLevel {
    log.Printf(msg, a...)
  }
}

func Fatal(msg string, a ...any) {
  logMsg(FATAL, msg, a...)
  os.Exit(1)
}

func Error(msg string, a ...any) {
  logMsg(ERROR, msg, a...)
}

func Warn(msg string, a ...any) {
  logMsg(WARN, msg, a...)
}

func Info(msg string, a ...any) {
  logMsg(INFO, msg, a...)
}

func Debug(msg string, a ...any) {
  logMsg(DEBUG, msg, a...)
}
