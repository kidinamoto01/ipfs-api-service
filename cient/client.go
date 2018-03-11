package cient

import (
	"bytes"
	"fmt"
	ipfs "github.com/kidinamoto01/go-ipfs-api"
)
const (
	examplesHash = "QmS4ustL54uo8FzR9455qaxZwuMiUhyvMcX9Ba8nUH4uVv"
	shellUrl     = "localhost:5001"
)
func AddString(input string) {

	s := ipfs.NewShell(shellUrl)

	mhash, err := s.Add(bytes.NewBufferString(input))
	if err!= nil{
		panic(err)
	}else{
		// "QmUfZ9rAdhV5ioBzXKdUTh2ZNsz9bzbkaLVyQ8uc8pj21F")
		fmt.Println("get output",mhash)

	}

}

func GetVersion(){

	s := ipfs.NewShell(shellUrl)

	version, commit,err := s.Version()
	if err!= nil{
		panic(err)
	}else{
		// "QmUfZ9rAdhV5ioBzXKdUTh2ZNsz9bzbkaLVyQ8uc8pj21F")
		fmt.Println("get version",version,commit)

	}
}
