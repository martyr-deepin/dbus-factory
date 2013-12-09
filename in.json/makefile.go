package main

import "fmt"
import "os/exec"
import "os"
import "path/filepath"
import "strings"
import "encoding/json"
import "flag"

func run_golang(name string) {
	defer func() {
		chs_go <- true
	}()
	f, err := os.Open(name)
	config := struct {
		Config struct {
			DestName string
		}
	}{}
	json.NewDecoder(f).Decode(&config)

	tmp := strings.Split(config.Config.DestName, ".")
	out := tmp[len(tmp)-1] + "_golang"

	stdout, err := exec.Command("proxyer", "-in", name, "-out", out, "-target", "golang").Output()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Generate " + name + " to golang failed!")
		return
	}
	fmt.Println("Generate " + name + " to golang OK!")
	fmt.Print(string(stdout))
	stdout, err = exec.Command("bash", "-c", "cd "+out+" && go build").Output()
	if err != nil {
		fmt.Println(err)
		fmt.Println("\tTest " + name + "golang code Failed!")
		return
	}
	fmt.Println("\tTest " + name + "golang code Ok!")
}

func run_qml(name string) {
	defer func() {
		chs_qml <- true
	}()
	f, err := os.Open(name)
	config := struct {
		Config struct {
			DestName string
		}
	}{}
	json.NewDecoder(f).Decode(&config)

	tmp := strings.Split(config.Config.DestName, ".")
	out := tmp[len(tmp)-1] + "_qml"

	stdout, err := exec.Command("proxyer", "-in", name, "-out", out, "-target", "qml").Output()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Generate " + name + " to QML failed!")
		return
	}
	fmt.Println("Generate " + name + " to QML OK !")
	fmt.Print(string(stdout))
	stdout, err = exec.Command("bash", "-c", "cd "+out+" && make ").Output()
	if err != nil {
		fmt.Println(err)
		fmt.Println("\tTest " + name + "QML code Failed!")

		return
	}
	fmt.Println("\tTest " + name + "QML code Ok!")
}

var chs_go chan bool
var chs_qml chan bool

func main() {
	var testQML bool
	var testGoLang bool
	flag.BoolVar(&testQML, "qml", false, "Test QML Target")
	flag.BoolVar(&testGoLang, "golang", true, "Test Golang Target (default enabled)")
	flag.Parse()
	modules := flag.Args()

	if len(modules) == 0 {
		var err error
		modules, err = filepath.Glob("*.in.json")
		if err != nil {
			panic(err)
		}
	}

	chs_go = make(chan bool, len(modules))
	chs_qml = make(chan bool, len(modules))
	for _, name := range modules {
		if testQML {
			go run_qml(name)
		}
		if testGoLang {
			go run_golang(name)
		}
	}
	if testQML {
		for i := 0; i < len(modules); i++ {
			<-chs_qml
		}
	}
	if testGoLang {
		for i := 0; i < len(modules); i++ {
			<-chs_go
		}
	}
}
