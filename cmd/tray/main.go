//TODO: Move this stuff to systray
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

func GetK8Contexts() []string{
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = path.Join(os.Getenv("HOME"), ".kube", "config")
	}

	k8cmd :=  exec.Command("kubectl", "config", "get-contexts")
	contexts_b, _ := k8cmd.Output()
	contexts := strings.Split(string(contexts_b), "\n")

	context_names := []string{}
	for _, ctx := range contexts {
		ctx_info := strings.Fields(ctx)
		if len(ctx_info) == 0 {
			continue
		}
		if ctx_info[0] == "*" {
			context_names = append(context_names, strings.TrimSpace(ctx_info[1]))
		} else{
			context_names = append(context_names, strings.TrimSpace(ctx_info[0]))
		}
	}

	return context_names
}

func K8Selector(ui webview.WebView) {

	ui.SetTitle("K8 Selector")
	ui.SetSize(480, 320, webview.HintNone)

	// A binding that increments a value and immediately returns the new value.
	context_names := GetK8Contexts()
	ui.Bind("filter", func(str string) Filter {
		return Filter{ FilteredStrings: StringMatch(context_names, str)}
	})

	ui.Bind("select_context", func(context_name string) error {
		k8cmd :=  exec.Command("kubectl", "config", "use-context", context_name)
		err := k8cmd.Run()
		if err != nil {
			log.Fatalf("Error switching context %s\n", err.Error())
		}
		return  err
	})

}

func main() {

	fname := "./ui.html"
	html_byte, err := os.ReadFile(fname); 
	if err != nil {
		log.Fatalf("Error reading file %s: %s\n", fname, err.Error())
	}
	html := string(html_byte)

	ui := webview.New(false)
	defer ui.Destroy()
	ui.Bind("close_window", func() {
		ui.Terminate()
	})

	K8Selector(ui)
	ui.SetHtml(html)
	ui.Run()
}
