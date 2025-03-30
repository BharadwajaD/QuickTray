package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func main() {
	systray.Run(onReady, nil)
}

func onReady() {
	fmt.Println("Systray started...")
	systray.SetIcon(icon.Data)
	systray.SetTitle("QuickTray")

	k8Menu := systray.AddMenuItem("K8 Selector", "Select Kubernetes Context")

	go func() {
		for {
			<-k8Menu.ClickedCh
			launchK8Selector()
		}
	}()
}

func launchK8Selector() {
	cwd, _ := os.Getwd()
	cmdPath := path.Join(cwd, "cmd", "k8er", "k8er")
	cmd := exec.Command(cmdPath) 

	out, err := cmd.CombinedOutput()
	log.Println(exec.LookPath(cmdPath))

	if err != nil {
		msg := fmt.Sprintf("Failed to start K8 selector: %s", err.Error())
		log.Println(msg)
		beeep.Notify("K8 Selector Error", msg, "")
	} else {
		msg := fmt.Sprintf("K8 Selector Output:\n%s", string(out))
		log.Println(msg)
		beeep.Notify("K8 Selector", msg, "")
	}
}
