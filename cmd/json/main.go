package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path"

	webview "github.com/webview/webview_go"
)

func Marshal(obj interface{}) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		log.Fatalf("Error encoding %+v\n", obj)
	}

	return string(bytes)
}

func Unmarshal(str string) interface{}{
	var v interface{}
	err := json.Unmarshal([]byte(str), &v)

	if err != nil {
		log.Fatalf("Error decoding %+s\n", str)
	}

	return v
}

type JsonString struct{
	Output interface{} `json:"output"`
}

func main(){

	marshal := flag.Bool("marshal" , false, "default is unmarshaling usecase")
	flag.Parse()

	cwd, _ := os.Getwd()
	htmlFile := path.Join(cwd, "../", "../", "html", "json_string.html")
	htmlByte, err := os.ReadFile(htmlFile)
	if err != nil {
		log.Fatalf("Error reading file %s: %s\n", htmlFile, err.Error())
	}

	html := string(htmlByte)

	ui := webview.New(false)
	defer ui.Destroy()
	ui.SetTitle("Json Un/Marshaler")
	ui.SetSize(480, 320, webview.HintNone)


	ui.Bind("json_string", func(input any) JsonString{
		var output any
		if *marshal {
			output = Marshal(input)
		}else {
			str := input.(string)
			output = Unmarshal(str)
		}

		return JsonString{ Output: output }
	})

	ui.Bind("close_window", func(){
		ui.Terminate()
	})

	ui.SetHtml(html)
	ui.Run()
}
