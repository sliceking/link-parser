package main

import (
	"fmt"
	"strings"

	"github.com/svwielga4/link-parser"
)

var htmlString = `
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Document</title>
</head>
<body>
  <a href="/page-one">i am the link to page one
	<span>CURVEBALL</span>
  </a>
  <a href="/page-two">I am page 2</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(htmlString)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", links)
}
