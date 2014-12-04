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

	return ctx

}

func (c Context) Reader() error {

	var data []byte
	xml.Unmarshal(data, &c)

	return nil
}

func (c Context) connect() {
}

func (c Context) ShowContext() error {

}

func (c Context) Get(data string) (string, error) {

	return "", nil
}

func (c *Context) request() {
	var url string = "http://api.wolframalpha.com/v2/query?appid=" + c.appid + "&input=" + "hello" + "&format=" + "image,plaintext"
	res, err := http.Get(url)
	if err != nil {
		fmt.Errorf("http Get failed %v", err)
		return
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("can not read resoponce body %v", err)
		return
	}
	var query QueryResult
	err = xml.Unmarshal(b, &query)
	if err != nil {
		fmt.Errorf("unmarshal error:%v", err)
		return
	}
	c.Query = query
}

func (q QueryResult) ConnectionSucess() bool {
	return q.Success
}

func (p Pod) PrintlnSubPod() error {
	for _, a := range p.Subpods {
		fmt.Println(a.Title)
	}

	return nil
}
