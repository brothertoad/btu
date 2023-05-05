package btu

import (
  "log"
  "os"
  "strconv"
  "strings"
)

const FATAL = 1
const ERROR = 2
const WARN  = 3
const INFO  = 4
const DEBUG = 5
const TRACE = 6
const minLevel = FATAL

var logLevel = ERROR

func SetLogLevel(level int) {
  if level < minLevel {
    log.Fatalf("log level %s is out of range, must be at least %d)\n", level, minLevel)
  }
  logLevel = level
}

func SetLogLevelByName(name string) {
  lcName := strings.ToLower(name)
  switch lcName {
  case "fatal":
    SetLogLevel(FATAL)
  case "error":
    SetLogLevel(ERROR)
  case "warn":
    SetLogLevel(WARN)
  case "info":
    SetLogLevel(INFO)
  case "debug":
    SetLogLevel(DEBUG)
  case "trace":
    SetLogLevel(TRACE)
  default:
    // try to convert to a number
    level, err := strconv.Atoi(name)
    CheckError(err)
    SetLogLevel(level)
  }
}

func logMsg(level int, msg string, a ...any) {
  if level <= logLevel {
    log.Printf(msg, a...)
  }
}

// An exported version of logMsg.
func Log(level int, msg string, a ...any) {
  logMsg(level, msg, a...)
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

func Trace(msg string, a ...any) {
  logMsg(TRACE, msg, a...)
}
