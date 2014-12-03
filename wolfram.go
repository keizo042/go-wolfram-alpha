package wolfram


type Context struct {
	xml.Header
	appid string
	Query QueryResult
}

type QueryResult struct {
	//tag
	XMLName		xml.Name `xml:"queryresult"`
	Success       bool   `xml:"success,attr"`
	Error         bool   `xml:"error,attr"`
	Numpods       int32  `xml:"numpods,attr"`
	Datatypes     string `xml:"datatypes,attr"`
	Timedout      string `xml:"timedout,attr"`
	Timedoutpods  string `xml:"timedoutpods,attr"`
	Timing        int32  `xml:"timing,attr"`
	Parsetiming   int32  `xml:"parsetiming,attr"`
	Parsetimedout bool   `xml:"parsetimedout"`
	Recalculate   string `xml:"recalculate,attr"`
	Id            string `xml:"id,attr"`
	Host          string `xml:"host,attr"`
	Server        int32  `xml:"server,attr"`
	Related       string `xml:"related,attr"`
	Version       string `xml:"version,attr"`

	Pods    []Pod
	Assumpt []Assumptions
}

type Pod struct {
	XMLName xml.Name `xml:pod`
	//tag
	Title      string `xml:title,attr`
	Scanner    string `xml:scanner,attr`
	Id         int32  `xml:id,attr`
	Position   string `xml:position,attr`
	Err        string `xml:error,attr`
	Numsubpods int32  `xml:numsubpods,attr`

	Subpods []SubPod
}

type States struct {
    Stat []State `xml:"states"`
}

type State struct {
    Name string `xml:"name"`
    Input string `xml:"input"`

}

type SubPod struct {
	XMLName   xml.Name `xml:"subpod,attr"`
	Title     string   `xml:"title,attr"`
	PlainText string   `xml:"plaintext"`
	Image       Img	 `xml:"img"`
}

type Img struct {
	Src    string `xml:"src"`
	Alt    string `xml:"alt"`
	Title  string `xml:"title"`
	Width  string `xml:"width"`
	height string `xml:"height"`
}

type Assumptions struct {
	count int32 `xml:"count,attr"`
	Assum []Assumption
}

type Assumption struct {
	Type   string `xml:"type,attr"`
	Word   string `xml:"word,attr"`
	Template string `xml:"template,attr"`
	Count  int32  `xml:count,attr"`
	Values []Value
}

type Value struct {
	XMLName xml.Name `xml:"name"`
	Name  string `xml:"name,attr"`
	desc  string `xml:"desc,attr"`
	input string `xml:"input,attr"`
}

type Sources struct {
    XMLsource []Source
}

type Source struct {
	XMLName xml.Name `xml:"source"`
	Url string `xml:"url,attr"`
	Text string `xml:"text,attr"`
}

