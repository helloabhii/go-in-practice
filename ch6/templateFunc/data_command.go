package main

import (
	"html/template"
)

func main() {
}

var tpl = `<!DOCTYPE HTML>
  <html>
  <head>
  <meta charset="utf-8">
  <title>DATE Example</title>
  </head>
  <body>
  <p>{{.Date | dateFormat "Feb 1, 2024"}}</p>
  </body>
  </html>`

var funcMap = template.FuncMap{
	"dateFormat": dateFormat,
}
