package fancy_logger

import (
	"fmt"
	"os"
	"regexp"
	"sync"
	"time"
)

type FileLogFilter int

// use OR to include multiple filters
const (
	Filter_Log     FileLogFilter = 1 << iota // 1 << 0 = 1
	Filter_Warning                           // 1 << 1 = 2
	Filter_Error                             // 1 << 2 = 4
	Filter_Info                              // 1 << 3 = 8
)

type LoggerConfig struct {
	//custom tag for declaring a struct log or a custom package name
	Tag          string
	UseTimeStamp bool
	LogToFile    bool
	TagColor     Colors
	ErrorColor   Colors
	InfoColor    Colors
	LogColor     Colors
	WarningColor Colors
	FileFilters  FileLogFilter
}

var file *os.File
var currentDay int
var fileMutex = sync.Mutex{}

type logger struct {
	conf LoggerConfig
}

func NewLogger(conf *LoggerConfig) *logger {
	if conf == nil {
		return &logger{
			conf: LoggerConfig{
				TagColor:     Colors_Magenta,
				ErrorColor:   Colors_Red,
				InfoColor:    Colors_Cyan,
				LogColor:     Colors_White,
				WarningColor: Colors_Yellow,
				Tag:          "APP",
				UseTimeStamp: true,
				LogToFile:    true,
				FileFilters:  Filter_Log | Filter_Error | Filter_Info | Filter_Warning,
			},
		}
	}
	return &logger{
		conf: *conf,
	}
}

func (l *logger) log(tag string, filter FileLogFilter, msgc Colors, message string, ts bool) {
	var timeStamp string
	if ts {
		now := time.Now()
		timeStamp = fmt.Sprintf("|%d-%02d-%02d %02d:%02d:%02d", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	}
	msg := fmt.Sprintf("%s[%s]%s: %s%s%s\n", l.conf.TagColor, tag+timeStamp, Colors_Reset, msgc, message, Colors_Reset)

	fmt.Print(msg)

	if l.conf.LogToFile {
		go logToFile(l, msg, l.conf.FileFilters, filter)
	}
}

func (l *logger) Log(message string) {
	l.log(l.conf.Tag+"|LOG", Filter_Log, l.conf.LogColor, message, l.conf.UseTimeStamp)
}

func (l *logger) LogWarning(message string) {
	l.log(l.conf.Tag+"|WAR", Filter_Warning, l.conf.WarningColor, message, l.conf.UseTimeStamp)
}

func (l *logger) LogError(message string) {
	l.log(l.conf.Tag+"|ERR", Filter_Error, l.conf.ErrorColor, message, l.conf.UseTimeStamp)
}

func (l *logger) LogInfo(message string) {
	l.log(l.conf.Tag+"|INF", Filter_Info, l.conf.InfoColor, message, l.conf.UseTimeStamp)
}

func logToFile(l *logger, message string, currentFilters, filterToCheck FileLogFilter) {
	fileMutex.Lock()
	defer fileMutex.Unlock()
	now := time.Now()

	if currentDay != now.Day() {
		if file != nil {
			file.Close()
			file = nil
		}
		fileName := fmt.Sprintf("LOGS-%d-%02d-%02d.txt", now.Year(), int(now.Month()), now.Day())
		var err error
		file, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			l.LogError(err.Error())
			return
		}
		currentDay = now.Day()
	}

	if currentFilters&filterToCheck != 0 {

		if file != nil {
			_, err := file.WriteString(removeANSIColors(message))
			if err != nil {
				l.LogError(err.Error())
			}
		}
	}
}

func removeANSIColors(input string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(input, "")
}
