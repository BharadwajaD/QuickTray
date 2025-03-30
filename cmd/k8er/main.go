package main

import (
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	webview "github.com/webview/webview_go"
)

type Filter struct {
	FilteredStrings []string `json:"filtered_strings"`
}

func StringMatch(strs []string, prefix string) []string {
	matches := make([]string, 0)
	for _, str := range strs {
		if strings.Contains(str, prefix) {
			matches = append(matches, str)
		}
	}
	return matches
}

func GetK8Contexts() []string {
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = path.Join(os.Getenv("HOME"), ".kube", "config")
	}

	k8cmd := exec.Command("kubectl", "config", "get-contexts", "--output=name")
	contexts_b, _ := k8cmd.Output()
	contexts := strings.Split(string(contexts_b), "\n")

	var contextNames []string
	for _, ctx := range contexts {
		trimmed := strings.TrimSpace(ctx)
		if trimmed != "" {
			contextNames = append(contextNames, trimmed)
		}
	}

	return contextNames
}

func main() {
	cwd, _ := os.Getwd()
	htmlFile := path.Join(cwd, "html", "k8.html")
	htmlByte, err := os.ReadFile(htmlFile)
	if err != nil {
		log.Fatalf("Error reading file %s: %s\n", htmlFile, err.Error())
	}
	html := string(htmlByte)

	ui := webview.New(false)
	defer ui.Destroy()

	ui.SetTitle("K8 Selector")
	ui.SetSize(480, 320, webview.HintNone)

	contextNames := GetK8Contexts()

	ui.Bind("filter", func(str string) Filter {
		return Filter{FilteredStrings: StringMatch(contextNames, str)}
	})

	ui.Bind("select_context", func(contextName string) error {
		cmd := exec.Command("kubectl", "config", "use-context", contextName)
		err := cmd.Run()
		if err != nil {
			log.Printf("Error switching context: %s\n", err.Error())
		}

		log.Printf("Switched to %s\n", contextName)
		return err
	})

	ui.Bind("close_window", func(){
		ui.Terminate()
	})

	ui.SetHtml(html)
	ui.Run()
}

