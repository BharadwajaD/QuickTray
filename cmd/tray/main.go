package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func main() {
	systray.Run(onReady, nil)
}

func onReady() {
	fmt.Println("Systray started...")
	systray.SetIcon(icon.Data)

	k8Menu := systray.AddMenuItem("K8 Selector", "Select Kubernetes Context")
	json_string := systray.AddMenuItem("Json Encoder", "Convert json to string")
	string_json := systray.AddMenuItem("Json Decoder", "Convert string to json")

	go func() {
		for {
			select {
			case <-k8Menu.ClickedCh: launchTool("k8er")
			case <-json_string.ClickedCh: launchTool("json",  "-marshal")
			case <-string_json.ClickedCh: launchTool("json")
			}
		}
	}()
}

func launchTool(name string, opts ...string) {
	cwd, _ := os.Executable()
	cwd, _ = filepath.EvalSymlinks(cwd)
	binPath := filepath.Join(cwd, "../", name)

	cmd := exec.Command(binPath, opts...)
	out, err := cmd.CombinedOutput()
	log.Println("Executing:", binPath, opts)

	if err != nil {
		log.Printf("Failed to start %s: %s : %s\n", name, err, string(out))
	} else {
		log.Printf("%s Output:\n%s", name, string(out))
	}
}
