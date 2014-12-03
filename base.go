package wolfram

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func NewContext(id string) Context {
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
    var url string ="http://api.wolframalpha.com/v2/query?appid=" + c.appid + "&input=" +"hello" +"&format=" +"image,plaintext"
    res, err := http.Get(url)
    if err != nil {
        fmt.Errorf(err)
        return
    }
   b, err  := ioutil.ReadAll(res.Body)
   if err != nil {
       fmt.Errorf(err)
       return
   }
   var query QueryResult
   err = xml.Unmarshal(b,&query)
   if err != nil {
       fmt.Errorf(err)
       return
   }
   c.Query = query
}

func (q QueryResult) IsSccuess() bool {
    return q.Success

}

func (p Pod) PrintlnSubPod() (string, error) {
    
    for a,_ := range p.Subpods{
        fmt.Println(a.Title)
    }

}

