package wolfram

// wolfram alpha api binding
//
//
// *** sample ***
// package main
// import (
//		"fmt"
//		"github/KeizoBookman/go-wolfram-alpha"
// )
// ctx := wolfram.New(<your-app-id>)
// ctx.Get()
// fmt.Println(ctx.Query)
//

import (
	"encoding/xml"
)

//
//  xml structs
//

type Client struct {
	appid string
	Query QueryResult
}

type QueryResult struct {

	//tag
	success bool `xml:"success,attr"`
	//
	//
	XMLName       xml.Name `xml:"queryresult"`
	Error         bool     `xml:"error,attr"`
	Numpods       int32    `xml:"numpods,attr"`
	Datatypes     string   `xml:"datatypes,attr"`
	Timedout      string   `xml:"timedout,attr"`
	Timedoutpods  string   `xml:"timedoutpods,attr"`
	Timing        int32    `xml:"timing,attr"`
	Parsetiming   int32    `xml:"parsetiming,attr"`
	Parsetimedout bool     `xml:"parsetimedout,attr"`
	Recalculate   string   `xml:"recalculate,attr"`
	Id            string   `xml:"id,attr"`
	Host          string   `xml:"host,attr"`
	Server        int32    `xml:"server,attr"`
	Related       string   `xml:"related,attr"`
	Version       string   `xml:"version,attr"`

	Pods        []Pod         `xml:"pod"`
	Assumpt     []Assumptions `xml:"assumptions"`
	Sources     Sources       `xml:"sources"`
	ExamplePage ExamplePage   `xml:"examplepage"`
	Scripts     string        `xml:"scripts"`
	Css         string        `xml;"css"`
}

type Pod struct {
	XMLName xml.Name `xml:"pod"`
	//tag
	Title      string `xml:"title,attr"`
	Scanner    string `xml:"scanner,attr"`
	Id         string `xml:"id,attr"`
	Position   string `xml:"position,attr"`
	Err        string `xml:"error,attr"`
	Numsubpods int32  `xml:"numsubpods,attr"`

	Subpods []SubPod `xml:"subod"`
}

type States struct {
	XMLName xml.Name `xml:"states"`
	Count   int32    `xml:"count"`
	State   []State  `xml:"state"`
}

type State struct {
	XMLName xml.Name `xml:"state"`
	Name    string   `xml:"name"`
	Input   string   `xml:"input"`
}

type SubPod struct {
	XMLName   xml.Name `xml:"subpod,attr"`
	Title     string   `xml:"title,attr"`
	PlainText string   `xml:"plaintext"`
	Image     Img      `xml:"img"`
}

type Img struct {
	XMLName xml.Name `xml:"img"`
	Src     string   `xml:"src"`
	Alt     string   `xml:"alt"`
	Title   string   `xml:"title"`
	Width   string   `xml:"width"`
	height  string   `xml:"height"`
}

type Assumptions struct {
	XMLName xml.Name     `xml:"assumptions"`
	count   int32        `xml:"count,attr"`
	Assum   []Assumption `xml:"assumption"`
}

type Assumption struct {
	Type     string  `xml:"type,attr"`
	Word     string  `xml:"word,attr"`
	Template string  `xml:"template,attr"`
	Count    int32   `xml:"count,attr"`
	Values   []Value `xml:"value"`
}

type Value struct {
	XMLName xml.Name `xml:"value"`
	Name    string   `xml:"name,attr"`
	desc    string   `xml:"desc,attr"`
	input   string   `xml:"input,attr"`
}

type Sources struct {
	XMLName xml.Name `xml:"sources"`
	Source  []Source `xml:"source"`
	Count   int32    `xml:"count,attr"`
}

type Source struct {
	XMLName xml.Name `xml:"source"`
	Url     string   `xml:"url,attr"`
	Text    string   `xml:"text,attr"`
}

type ExamplePage struct {
	Category string `xml:"category,attr"`
	Url      string `xml:"url,attr"`
}

//
// flag
//
const (
	//prime number
	A_PLAIN  int = 2
	A_IMAGE  int = 3
	A_SOUND  int = 5
	A_HTML   int = 7
	A_CELL   int = 11
	A_MINPUT int = 13

	A_PI int = A_PLAIN * A_IMAGE
	A_PS int = A_PLAIN * A_SOUND
	A_PH int = A_PLAIN * A_HTML
	A_PC int = A_PLAIN * A_CELL
	A_PM int = A_PLAIN * A_MINPUT

	A_IS int = A_IMAGE * A_SOUND
	A_IH int = A_IMAGE * A_SOUND
	A_IC int = A_IMAGE * A_SOUND
	A_IM int = A_IMAGE * A_SOUND

	A_SH int = A_SOUND * A_HTML
	A_SC int = A_SOUND * A_CELL
	A_SM int = A_SOUND * A_MINPUT

	A_HC int = A_HTML * A_CELL
	A_HM int = A_HTML * A_MINPUT

	A_PIS int = A_PLAIN * A_IMAGE * A_SOUND
	A_PIH int = A_PLAIN * A_IMAGE * A_HTML
	A_PIC int = A_PLAIN * A_IMAGE * A_CELL
	A_PIM int = A_PLAIN * A_IMAGE * A_MINPUT
	A_PSH int = A_PLAIN * A_SOUND * A_HTML
	A_PSC int = A_PLAIN * A_SOUND * A_CELL
	A_PSM int = A_PLAIN * A_SOUND * A_MINPUT
	A_PHC int = A_PLAIN * A_HTML * A_CELL
	A_PHM int = A_PLAIN * A_HTML * A_MINPUT
	A_ISH int = A_IMAGE * A_SOUND * A_HTML
	A_ISC int = A_IMAGE * A_SOUND * A_CELL
	A_ISM int = A_IMAGE * A_SOUND *  A_MINPUT
	A_IHC int = A_IMAGE * A_HTML * A_CELL
	A_IHM int = A_IMAGE * A_HTML * A_MINPUT
	A_ICM int = A_IMAGE * A_CELL * A_MINPUT
	A_SHC int = A_SOUND * A_HTML * A_CELL
	A_SHM int = A_SOUND * A_HTML * A_MINPUT
	A_SCM int = A_SOUND * A_CELL * A_MINPUT
	A_ALL int = A_PLAIN * A_IMAGE * A_SOUND * A_HTML * A_CELL * A_MINPUT
)
