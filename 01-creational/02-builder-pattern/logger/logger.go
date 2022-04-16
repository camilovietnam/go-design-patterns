package main

import "fmt"

/*
	This is the interface that the logger objects need to implement in order to
	be built by the Builder object.
*/
type Logger interface {
	ReadLine() string
}

/*
	LogAdministrator is our Builder, it will build the loggers depending
	on the type of logger that we pass to it.
*/

// these are the logger types, database and file.
type loggerType int

const (
	DATABASE_LOGGER loggerType = iota
	FILE_LOGGER
)

// LogAdministrator class will have an internal logger object from where it will
// read its lines.
type LogAdministrator struct {
	logger Logger
}

// BuildLogger function will receive the type of logger we want to build, and use
// the internal object to store the logger built.
func (a *LogAdministrator) BuildLogger(t loggerType) {
	switch t {
	case DATABASE_LOGGER:
		a.logger = databaseLogger{}
	case FILE_LOGGER:
		a.logger = fileLogger{}
	}
}

// ReadLine will use the internal logger object to read a line, regardless
// of the origin of it.
func (a *LogAdministrator) ReadLine() string {
	return a.logger.ReadLine()
}

// These are the specific implementations of the Logger interface, with their
// ReadLine function.
type databaseLogger struct{}

func (d databaseLogger) ReadLine() string {
	return "[>] This line was read from a database"
}

type fileLogger struct{}

func (f fileLogger) ReadLine() string {
	return "[>] This line was read from a file"
}

// Main function.
func main() {
	echo("[>] Builder test")

	admin := LogAdministrator{}

	admin.BuildLogger(DATABASE_LOGGER)
	echo(admin.ReadLine())

	admin.BuildLogger(FILE_LOGGER)
	echo(admin.ReadLine())
}

// echo method to simplify fmt.Println.
func echo(args ...interface{}) {
	fmt.Println(args...)
}
