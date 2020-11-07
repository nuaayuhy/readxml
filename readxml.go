package main

import (
	"fmt"
	"github.com/beevik/etree"
)

func readXml(path string) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(path); err != nil {
		panic(err)
	}

	// root
	root := doc.SelectElement("root")
	fmt.Println("ROOT element:", root.Tag)

	// area
	for _,area  := range root.SelectElements("area") {
		fmt.Println("CHILD element:", area.Tag)
		if volume := area.SelectElement("volume"); volume != nil {
			lang := volume.SelectAttrValue("lang", "unknown")
			fmt.Printf("  TITLE: %s (%s)\n", title.Text(), lang)
		}
		for _, attr := range book.Attr {
			fmt.Printf("  ATTR: %s=%s\n", attr.Key, attr.Value)
		}
	}
}

func main()  {
	readXml("physical_sector.xml")
}
