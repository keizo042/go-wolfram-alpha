package mathematica

import (
    wolfram "github.com/KeizoBookman/go-wolfram-alpha"
)

//wolfram mathematica api binding
//wolfram mathematica 

func (c *wolfram.Client )Fetch(t string) ( wolfram.QueryResult, err){
     q,err := c.Get(t)

     if err != nil {
	 return nil, err
     }
     return q, nil
}

func New(id string) wolfram.Client{
    return wolfram.New(id)
}

