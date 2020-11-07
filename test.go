package main

import (
	"encoding/xml"
	"fmt"
)

const data = `<?xml version="1.0" encoding="UTF-8"?>
<root recordcount="1">
	<area id="1" priority="0" physicalname="VACC13" volume_sum="1">
		<volume volumename="CSB05" lower="7201" upper="9550" shape="P" point_sum="16">
			<points longitude="1121000" latitude="292500" longi_xpos="-110932.210519" lati_ypos="668112.608812" pointname="CX128" radius="">
			</points>
			<points longitude="1130712" latitude="292300" longi_xpos="-18417.707381" lati_ypos="663889.290109" pointname="CX15" radius="">
			</points>
			<points longitude="1143400" latitude="290200" longi_xpos="122441.098586" lati_ypos="625732.512220" pointname="CX16" radius="">
			</points>
			<points longitude="1141326" latitude="274501" longi_xpos="90127.677621" lati_ypos="483219.885923" pointname="CX171" radius="">
			</points>
			<points longitude="1133142" latitude="273742" longi_xpos="21576.332868" lati_ypos="469390.979555" pointname="CX170" radius="">
			</points>
			<points longitude="1123800" latitude="265430" longi_xpos="-67190.120148" lati_ypos="389769.866708" pointname="CX130" radius="">
			</points>
			<points longitude="1121729" latitude="273718" longi_xpos="-100515.767562" lati_ypos="469047.300899" pointname="CX169" radius="">
			</points>
			<points longitude="1115624" latitude="274111" longi_xpos="-135123.713768" lati_ypos="476555.790553" pointname="CX168" radius="">
			</points>
			<points longitude="1114758" latitude="274810" longi_xpos="-148832.998420" lati_ypos="489617.983838" pointname="CX167" radius="">
			</points>
			<points longitude="1114253" latitude="275805" longi_xpos="-156944.710274" lati_ypos="508043.246248" pointname="CX166" radius="">
			</points>
			<points longitude="1114249" latitude="280947" longi_xpos="-156770.692733" lati_ypos="529659.360798" pointname="CX165" radius="">
			</points>
			<points longitude="1114720" latitude="281906" longi_xpos="-149159.657164" lati_ypos="546776.159831" pointname="CX161" radius="">
			</points>
			<points longitude="1115540" latitude="282629" longi_xpos="-135379.830793" lati_ypos="560252.302039" pointname="CX162" radius="">
			</points>
			<points longitude="1120434" latitude="283007" longi_xpos="-120778.394548" lati_ypos="566806.341196" pointname="CX163" radius="">
			</points>
			<points longitude="1121144" latitude="283100" longi_xpos="-109068.022236" lati_ypos="568323.781534" pointname="CX164" radius="">
			</points>
			<points longitude="1121000" latitude="292500" longi_xpos="-110932.210519" lati_ypos="668112.608812" pointname="CX128" radius="">
			</points>
		</volume>
	</area>
</root>
`

// root
type Root struct {
	XMLName xml.Name `xml:"root"`
	Recordcount string `xml:"recordcount,attr"`
	Areas Area `xml:"area"`
}

// area
type Area struct {
	Id int `xml:"id,attr"`
	Priority int `xml:"priority,attr"`
	Physicalname string `xml:"physicalname,attr"`
	Volume_sum int `xml:"volume_sum,attr"`
	Volumes Volume `xml:"volume"`	
}

// volume
type Volume struct {
	Volumename string `xml:"volumename,attr"`
	Lower int `xml:"lower,attr"`
	Upper int `xml:"upper,attr"`
	Shape string `xml:"shape,attr"`
	Point_sum int `xml:"point_sum,attr"`
	Points []*Point `xml:"points"`
}

// points
type Point struct {
	Longitude int `xml:"longitude,attr"`
	Latitude int `xml:"latitude,attr"`
	Longi_xpos float64 `xml:"longi_xpos,attr"`
	Lati_ypos float64 `xml:"lati_ypos,attr"`
	Pointname string `xml:"pointname,attr"`
	Radius int `xml:"radius,attr"`
}


func main() {	  

	v := Root{}
	err1 := xml.Unmarshal([]byte(data), &v)
	if  err1 != nil {
		panic(err1)
	}

	fmt.Println(v)
	fmt.Println("root.Recordcount: ",v.Recordcount)
	fmt.Println("Area.id: ",v.Areas.Id)
	fmt.Println("Area.volume_sum: ",v.Areas.Volume_sum)
	fmt.Println("Volume.lower: ",v.Areas.Volumes.Lower)
	fmt.Println("Volume.upper: ",v.Areas.Volumes.Upper)
	fmt.Println("Volume.point_sum: ",v.Areas.Volumes.Point_sum)
	for i,element:= range v.Areas.Volumes.Points {
		fmt.Println(i,element)
	}

}