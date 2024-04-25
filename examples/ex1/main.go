package main

import (
	"fmt"
	link "html_link_parser"
	"strings"
)

var exampleHtml = `
  <html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">
  A link to another page
  <span>sone span</span>
  </a>
  <a href="/first-page">First page link</a>
</body>
</html>
  `

func main() {
  r := strings.NewReader(exampleHtml)
  links, err := link.Parse(r)
  if err != nil {
    panic(err)
  }
  fmt.Printf("%+v\n", links)
}
