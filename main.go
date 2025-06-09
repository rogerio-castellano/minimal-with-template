package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var addr = flag.String("addr", "127.0.0.1:1718", "http service address") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	http.Handle("/original", http.HandlerFunc(QROriginal))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

type Data struct {
	Values map[string]string
}

// It ran three times before adding the /favicon.ico route rule
func QR(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	values := map[string]string{"text": req.FormValue("text"), "caption": req.FormValue("caption")}
	data := &Data{Values: values}
	fmt.Print(req.URL.Path)
	templ.Execute(w, data)
}

const templateStr = `
					<html class="text-gray-900 antialiased leading-tight">
					<head>
						<title>QR Link Generator</title>
						<meta charset="UTF-8" />
						<meta name="viewport" content="width=device-width, initial-scale=1.0" />
						<script src="https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"></script>

					</head>
					<body class="min-h-screen bg-gray-100 ml-5">
						<div class="mb-5">
						{{$text := index .Values "text"}}
						</div>
						<div class="mb-5">
						{{$caption := index .Values "caption"}}
						</div>
						{{if $text }}
							<div class="w-full max-w-lg rounded bg-gray-300 inline-block text-center px-2 py-2 mb-5">
								<img class="mx-auto" src="https://quickchart.io/qr?text={{ $text }}&caption={{ $caption }}" />
								<br>
								<p class="text-indigo-800 font-medium text-center">{{$text}}</p>
							</div>
						{{end}}
						<form action="/" name=f method="GET" class="w-full max-w-lg">
							<div class="mb-5">
							<label for="text" class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2">Text to QR Encode</label>
							<input maxLength=1024 size=70 name=text value="" title="Text to QR Encode" class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500">
							</div>
							<div class="mb-5">
							<label for="caption" class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2">Caption</label>
							<input maxLength=1024 size=70 name=caption value="" title="Caption" class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500">
							</div>

							<input type=submit value="Show QR" name=qr class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">
						</form>
					</body>
					</html>
`

// It works, running once per submission
var templOriginal = template.Must(template.New("qr_original").Parse(templateStrOriginal))

func QROriginal(w http.ResponseWriter, req *http.Request) {
	fmt.Println("QROriginal running with text", req.FormValue("s"))
	templOriginal.Execute(w, req.FormValue("s"))
}

const templateStrOriginal = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="https://quickchart.io/qr?text={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/original" name=f method="GET">
    <input maxLength=1024 size=70 name=s value="" title="Text to QR Encode">
    <input type=submit value="Show QR" name=qr>
</form>
</body>
</html>
`
