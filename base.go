package wolfram

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func New(id string) *Client {
    // Make New Client Structure
    //
	var ctx Client
	ctx.appid = id

	return &ctx

}

func (c *Client) Get(data string)  (*QueryResult,error) {
    //export function
    // atodeyaru
    err := c.request(data)
    return &c.Query,err
}

func (c *Client) request(input string) error{
    // internal process to request wolfram-alpha.com engine.
    // 
    // ***future implementation***
    // you can choice format : image, plaintext, mathematica input, and other
    // we can choice some way
    // args -> const variable like os.O_XXX
    //      -> make strucure filed {image bool,plaintext bool,mathematica_input bool}

	var url string = "http://api.wolframalpha.com/v2/query?appid=" + c.appid + "&input=" + input + "&format=" + "image,plaintext"
	var query QueryResult

    //
	res, err := http.Get(url)
	if err != nil {
		fmt.Errorf("http Get failed %v", err)
		return err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("can not read resoponce body %v", err)
		return err
	}

	err = xml.Unmarshal(b, &query)
	if err != nil {
		fmt.Errorf("unmarshal error:%v", err)
		return err
	}
	c.Query = query

	return nil
}

func (c *Client) IsSuccessed() bool {
    //check to query-result server-side states
	return  c.Query.success
}

func (c *Client) ShowClient() error {
    //maybe delete

    return nil
}

func flagCheck(f flag) string {

    return ""

}
