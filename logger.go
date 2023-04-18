package debugger

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"time"
)

func CreateFol() string {
	homeDir, err := os.UserHomeDir()
	e(err)
	err = os.Mkdir(homeDir+"/.Tsunagu", 0755)
	if err != nil {
		fmt.Print()
	}
	return homeDir
}
func ReadFile() {
	homeDir, err := os.UserHomeDir()
	e(err)
	data, err := ioutil.ReadFile(homeDir + "/.Tsunagu/Debugger.log")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
func Messages(s string) {
	homeDir, _ := os.UserHomeDir()
	time := time.Now()
	var file, err = os.OpenFile(homeDir+"/.Tsunagu/Debugger.log", os.O_APPEND|os.O_WRONLY, fs.ModeAppend)
	e(err)
	defer file.Close()
	date := fmt.Sprintf("[%s]", time.Format("2006-01-02 15:04:05"))
	data := date + " MESSAGE: " + s
	_, err = file.WriteString(data + "\n")
	e(err)
}
func Warning(s string) {
	homeDir, _ := os.UserHomeDir()
	time := time.Now()
	var file, err = os.OpenFile(homeDir+"/.Tsunagu/Debugger.log", os.O_APPEND|os.O_WRONLY, fs.ModeAppend)
	e(err)
	defer file.Close()
	date := fmt.Sprintf("[%s]", time.Format("2006-01-02 15:04:05"))
	data := date + " WARNING: " + s
	_, err = file.WriteString(data + "\n")
	e(err)
}
func Error(s string) {
	homeDir, _ := os.UserHomeDir()
	time := time.Now()
	var file, err = os.OpenFile(homeDir+"/.Tsunagu/Debugger.log", os.O_APPEND|os.O_WRONLY, fs.ModeAppend)
	e(err)
	defer file.Close()
	date := fmt.Sprintf("[%s]", time.Format("2006-01-02 15:04:05"))
	data := date + " ERROR: " + s
	_, err = file.WriteString(data + "\n")
	e(err)
}

func e(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
