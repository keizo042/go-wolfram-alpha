package wolfram

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func New(id string) Client {
    // Make New Client Structure
    //
	var ctx Client
	ctx.appid = id

	return ctx

}

func (c *Client) Reader() error {
    // it is fulled to io.Reader & io.Writer
    // i don have some idea

	var data []byte
	xml.Unmarshal(data, &c)

	return nil
}

func (c *Client) connect() {
    //some function to request?
}

func (c *Client) ShowClient() error {
    //maybe delete

    return nil
}

func (c *Client) Get(data string) (string, error) {
    //export function
    // atodeyaru

	return "", nil
}

func (c *Client) request(input string) {
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
		return
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("can not read resoponce body %v", err)
		return
	}

	err = xml.Unmarshal(b, &query)
	if err != nil {
		fmt.Errorf("unmarshal error:%v", err)
		return
	}
	c.Query = query
}

func (c *Client) ConnectionSucess() bool {
    //check to query-result server-side states
	return  c.Query.success
}

func (p Pod) PrintlnSubPod() error {
	for _, a := range p.Subpods {
		fmt.Println(a.Title)
	}

	return nil
}
