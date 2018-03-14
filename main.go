package main

import (
	c "github.com/kidinamoto01/ipfs-client/cient"
)


func main(){


	example := "/Users/b/Documents/go/src/github.com/irisnet/iris-hub/README.md"

	c.AddFile(example)
	//c.GetVersion()
	//p,err:= c.ShowPeers()
	//if err != nil{
     //  panic(err)
	//} else{
	//	fmt.Println(p)
	//}

}