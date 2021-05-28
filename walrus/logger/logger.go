package logger

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type (
	LogT struct {
		FileTextPath string `json:"fileTextPath"`
		FileJSONPath string `json:"fileJsonPath"`
		JSON         bool   `json:"json"`
		Level        string `json:"level"`
		Environment  string `json:"environment"`
	}
)

var (
	log = logrus.New()
)

func init() {
	initLogger()
}

func loadConfigLogger(path string) (*LogT, error) {

	l := new(LogT)
	path = strings.TrimSpace(filepath.Clean(path))

	if len(path) == 0 {
		fmt.Println("O caminho do arquivo passado está vazio.")
		return l, nil
	}

	f, err := os.Open(path)

	if err != nil {
		panic(err.Error())
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	bFile, err := ioutil.ReadAll(reader)

	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(bFile, &l); err != nil {
		panic(err)
	}

	return l, nil
}

// set do path de arquivo dos logs
// escrever no arquivo
// set do level
// set true ou false para o tipo de logs, se terá arquivo em formato json além dos logs
func initLogger() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic Recuperado na função", r)
		}
	}()

	var logPath string = "walrus/logger/config.json"
	logC, err := loadConfigLogger(logPath)

	// fileLog, _ := os.OpenFile(logC.FileTextPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	// fileJson, _:= os.OpenFile(logC.FileJSONPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)

	if err != nil {
		panic(err)
	}

	if logC.JSON {
		log.Formatter = new(logrus.JSONFormatter)
		// faz o set do formato para json
		//log.SetFormatter(&logrus.JSONFormatter{})
		// Faz o output para o arquivo json
		//log.Out = fileJson // TODO ver como gravar em dois arquivos diferentes utilizando uma única instância
		// logrus.SetOutput(fileJson)
	}

	// Log as TEXT instead of the default ASCII formatter.
	log.Formatter = new(logrus.TextFormatter)
	// TODO parte de configuração
	// log.Formatter.(*logrus.TextFormatter).DisableColors = true    // remove colors
	// log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	// log.SetFormatter(&logrus.TextFormatter{
	//	FullTimestamp: true,
	//})
	log.Out = os.Stdout
	// Faz o output para o arquivo de logs
	logrus.SetOutput(log.Out)
	// logrus.SetOutput(os.Stdout) // com Stdout

	// Only logger the warning severity or above.
	switch logC.Level {
	case "E_PRODUCTION":
		log.Level = logrus.InfoLevel
		// log.SetLevel(logrus.InfoLevel)
	case "E_DEVEL":
		log.Level = logrus.DebugLevel
		// log.SetLevel(logrus.DebugLevel)
	default:
		log.Level = logrus.TraceLevel
		// log.SetLevel(logrus.TraceLevel)
	}
}

func Trace(msg string) {
	log.WithFields(logrus.Fields{
		"field":  "Trace",
		"field2": 10,
	}).Trace(msg)
}

func Debug(msg string) {
	log.WithFields(logrus.Fields{
		"field":  "Debug",
		"field2": 10,
	}).Debug(msg)
}

func Info(msg string) {
	log.WithFields(logrus.Fields{
		"field":  "Info",
		"field2": 10,
	}).Info(msg)
}

func Warn(msg string) {
	log.WithFields(logrus.Fields{
		"field":  "Warn",
		"field2": 10,
	}).Warn(msg)
}

func Error(msg string) {
	log.WithFields(logrus.Fields{
		"field":  "Error",
		"field2": 10,
	}).Error(msg)
}
