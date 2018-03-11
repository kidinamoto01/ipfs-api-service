package main

import (
	c "github.com/kidinamoto01/ipfs-client/cient"
	"fmt"
)


func main(){


	//example := "Hello IPFS Shell tests"

	//c.AddString(example)
	//c.GetVersion()
	p,err:= c.ShowPeers()
	if err != nil{
       panic(err)
	} else{
		fmt.Println(p)
	}

}