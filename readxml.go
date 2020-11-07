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
		for _, attr := range area.Attr {
			fmt.Printf("%s=%s\t", attr.Key, attr.Value)
		}
		fmt.Printf("\n")

		//volume
		for _,volume := range area.SelectElements("volume"){
			// for _, attr := range volume.Attr {
			// 	fmt.Printf("  ATTR: %s=%s\t", attr.Key, attr.Value)
			// }
			volumename := volume.SelectAttrValue("volumename", "unknown")
			lower := volume.SelectAttrValue("lower", "unknown")
			upper := volume.SelectAttrValue("upper", "unknown")
			point_sum := volume.SelectAttrValue("point_sum", "unknown")
			fmt.Printf("volumename=%s\tlower=%s\tupper=%s\tpoint_sum=%s\n", volumename,lower,upper,point_sum)
			
			// points
			for _,points := range volume.SelectElements("points"){
				// fmt.Println("CHILD element:", points.Tag)
				// for _, attr := range points.Attr {
				// 	fmt.Printf("  ATTR: %s=%s\t", attr.Key, attr.Value)
				// }
				longitude := points.SelectAttrValue("longitude", "unknown")
				latitude := points.SelectAttrValue("latitude", "unknown")
				fmt.Printf("%s,%s\n", longitude,latitude)	
			}
		}
	}
}

func main() {

	readXml("physical_sector.xml")
}
