package main
/*

使用 etree 解析复杂结构的 xml 文件
https://godoc.org/github.com/beevik/etree
https://pkg.go.dev/github.com/beevik/etree?tab=doc
https://github.com/beevik/etree
*/

import (
	"fmt"
	"github.com/beevik/etree"// go get github.com/beevik/etree
)

func readXml(path string) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(path); err != nil {
		panic(err)
	}

	root := doc.SelectElement("bookstore")
	fmt.Println("ROOT element:", root.Tag)

	for _, book := range root.SelectElements("book") {
		fmt.Println("CHILD element:", book.Tag)
		if title := book.SelectElement("title"); title != nil {
			lang := title.SelectAttrValue("lang", "unknown")
			fmt.Printf("  TITLE: %s (%s)\n", title.Text(), lang)
		}
		for _, attr := range book.Attr {
			fmt.Printf("  ATTR: %s=%s\n", attr.Key, attr.Value)
		}
	}
}

func main()  {
	readXml("bookstores.xml")
}