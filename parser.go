package parser

import "time"

// LogFormat : Format specification for error log entries
type LogFormat struct {
	Time     time.Time
	Loglevel string
	Pid      int
	Tid      int
	Source   string
	Apr      string
	Client   string
	Message  string
	Request  string
}
