package cient

import (

	ipfs "github.com/kidinamoto01/go-ipfs-api"

	"fmt"
)



func ShowPeers()(string, error){
	s := ipfs.NewShell(shellUrl)

	peers, err := s.ShowPeers()
	if err!= nil{
		panic(err)
	}else{
	// "QmUfZ9rAdhV5ioBzXKdUTh2ZNsz9bzbkaLVyQ8uc8pj21F")
		fmt.Println("get output",peers)
		return peers,nil
	}


}
