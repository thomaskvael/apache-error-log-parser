package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/CrowdSurge/banner"
	"github.com/spf13/viper"
)

// LogFormat : Format specification for error log entries
type LogFormat struct {
	Time     time.Time
	Loglevel string
	Pid      int
	Tid      int
	Apr      string
	Client   string
	Message  string
}

// RegexCollection : Regex used to catch Apache error log entries
var RegexCollection = []string{
	`\[(.+?)\]`,
	`\s\[(.*?)\]`,
	`\s\[pid\s(.*?):`,
	`tid\s(\d+?)\]`,
	`\]\s(\(.*?\).*?\:)\s\[`,
	`\s\[client\s(.+?)]`,
	`\d\]\s(\w[^,\n]*)`,
}

// ConfigFile : Define path and filename (without extension) of config file
type ConfigFile struct {
	Path     string
	Filename string
}

// ConfigContent : Entries in config file
type ConfigContent struct {
	Log        string
	TimeFormat string
}

// C : Stores ConfigContent
var C ConfigContent

// Welcome : Print startup
func Welcome() {
	banner.Print("apache error log parser")
	banner.Print("-----------------------")
}

// Config : Handle configuration file
func Config(path, filename string) {
	config := ConfigFile{path, filename}
	viper.AddConfigPath(config.Path)
	viper.SetConfigName(config.Filename)

	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Get config entry from Configuration file
	json := viper.Sub("config")

	err := json.Unmarshal(&C)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	Welcome()
	parseLog(C.Log)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func matchRegex(line string) {
	logformat := LogFormat{}
	logMap := map[string]string{}
	e := reflect.ValueOf(&logformat).Elem()
	logCounter := -1
	for _, errorRegex := range RegexCollection {
		logCounter++
		re1, err := regexp.Compile(errorRegex)
		if err != nil {
			log.Fatalf("regexp: %s", err)
		}
		result := re1.FindStringSubmatch(line)

		LogItem := e.Type().Field(logCounter).Name

		if len(result) > 0 {
			fmt.Println(LogItem, result[1])
			logMap[strings.ToLower(LogItem)] = result[1]
		} else {
			fmt.Println(LogItem, "")
		}
	}
	insertDatabase(logMap)
	/*
		fmt.Println(logMap)
		for k, v := range logMap {
			fmt.Println(k, v)

		}
	*/
	return
}

func parseLog(file string) ([]LogFormat, error) {
	var items []LogFormat

	lines, err := readLines(file)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for _, line := range lines {
		fmt.Println("##############################################")
		fmt.Println("")
		matchRegex(line)
		fmt.Println("")
	}
	return items, nil
}
