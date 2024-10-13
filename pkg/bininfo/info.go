package bininfo

import "fmt"

var (
	SystemName string
	Version    string
	CommitID   string
	BuildTime  string
	GoVersion  string
)

var (
	GinLogMode string
)

var (
	StartTime int64
)

func String() string {
	s := fmt.Sprintln("\n******************************")
	s += fmt.Sprintf("*\t SystemName : %s\n", SystemName)
	s += fmt.Sprintf("*\t    Version : %s\n", Version)
	s += fmt.Sprintf("*\t   CommitID : %s\n", CommitID)
	s += fmt.Sprintf("*\t  BuildTime : %s\n", BuildTime)
	s += fmt.Sprintf("*\t  GoVersion : %s\n", GoVersion)
	s += fmt.Sprintln("******************************")
	return s
}
