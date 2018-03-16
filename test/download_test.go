package storage

import (
"fmt"
"testing"


"bytes"

"github.com/cheekybits/is"
ipfs "github.com/kidinamoto01/go-ipfs-api"
)


// 测试用例上传文件路径=nil
func TestDownloadNil(t *testing.T) {
	fmt.Println("测试用例上传文件路径=nil")
	is := is.New(t)
	s := ipfs.NewShell(shellUrl)

	err := s.Get("","")

	is.Equal(err, fmt.Errorf("Empty Path"))
}


func TestDownloadTxt(t *testing.T) {
	fmt.Println("测试下载TXT")
	is := is.New(t)
	s := ipfs.NewShell(shellUrl)

	mhash, err := s.Add(bytes.NewBufferString("Hello IPFS Shell tests"))
	is.Nil(err)
	is.Equal(mhash, "QmUfZ9rAdhV5ioBzXKdUTh2ZNsz9bzbkaLVyQ8uc8pj21F")
	err = s.Get(mhash,"/")
	is.Nil(err)

}


func TestDownloadJPG(t *testing.T) {
	fmt.Println("测试上传JPG,需要绝对路径")
	is := is.New(t)
	s := ipfs.NewShell(shellUrl)

	mhash, err := s.AddFile("/Users/b/Documents/go/src/github.com/irisnet/iris-hub/testdata/test.png")
	is.Nil(err)
	is.Equal(mhash, "QmaWsjUEsUV8bLCfkzyecyC3hnubeyZ9PVhNLiZ5JAvAfr")
}