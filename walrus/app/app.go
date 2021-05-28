package app

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
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

var log = logrus.New()

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
		panic(err)
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

func fileWriter(f string) io.Writer {
	file, err := os.Open(f)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileWriter := bufio.NewWriter(file)

	return fileWriter
}

// set do path de arquivo dos logs
// escrever no arquivo
// set do level
// set true ou false para o tipo de logs, se terá json além dos logs
func initLogger() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic Recuperado  na função", r)
		}
	}()

	var logPath string = "./logger_config.json"
	logC, err := loadConfigLogger(logPath)

	if err != nil {
		panic(err)
	}

	if logC.JSON {
		// faz o set do formato para json
		log.SetFormatter(&logrus.JSONFormatter{})

		// Faz o output para o arquivo json
		log.SetOutput(fileWriter(logC.FileJSONPath))
	}

	// Log as TEXT instead of the default ASCII formatter.
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Faz o output para o arquivo de logs
	log.SetOutput(fileWriter(logC.FileTextPath))

	// Only logger the warning severity or above.
	switch logC.Level {
	case "E_PRODUCTION":
		log.SetLevel(logrus.InfoLevel)
	case "E_DEVEL":
		log.SetLevel(logrus.TraceLevel)
	default:
		log.SetLevel(logrus.DebugLevel)
	}
}

func Trace(msg string) {
	log.WithFields(logrus.Fields{
		"field":  "test",
		"field2": 10,
	}).Trace(msg)
}

func Debug(msg string) {
	log.WithFields(logrus.Fields{
		"field":  "test",
		"field2": 10,
	}).Debug(msg)
}

func Info(msg string) {
	log.WithFields(logrus.Fields{
		"field":  "test",
		"field2": 10,
	}).Info(msg)
}
