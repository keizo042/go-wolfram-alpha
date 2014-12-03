package wolfram

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

func NewContext(id string) wolfram.Context {
    var ctx Context
    ctx.appid = id

    if false {

    }
    return ctx

}

func (c Context) Reader() error {

    var data string
    xml.Unmarshal(data,&c)

}

func (c Context) connect () {

    if err != nil {
    }
}

func (c Context) ShowContext () error {

}

func (c Context) Get(data string) (string, error) {

}

func (c *Context) request() {
    var url string ="http://api.wolframalpha.com/v2/query?appid=" + c.appid + "&input=hello&format=image,plaintext"
}

func (q QueryResult) IsSccuess() (bool) {
    return q.Success

}

func (p Pod) PrintlnSubPod() (string, error) {
    
    for a,_ := range p.Subpods{
        fmt.Println(a.Title)
    }

}

