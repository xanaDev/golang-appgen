package types

import (
	"sync"
)
type Messages struct {
	Error string `json:"ERROR"`
	Info  string `json:"INFO"`
	Warn  string `json:"WARN"`
	Fatal string `json:"FATAL"`
	Debug string `json:"DEBUG"`
}

type LoggingFramework struct {
	LibraryName string   `json:"LIBRARY_NAME"`
	ImportPath  string   `json:"IMPORT_PATH"`
	Version     string   `json:"VERSION"`
	Messages    Messages `json:"MESSAGES"`
}

type Configuration struct {
	AppName string
	Logging *LoggingFramework
}

var Mutex = &sync.Mutex{}