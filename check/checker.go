package check

import (
	"fmt"
	"log"
	"path"
	"time"

	"github.com/xlab/treeprint"
)

type Checker struct {
	RootJs  string
	TagDate time.Time
}

func (ck *Checker) Process() error {
	rootName := path.Base(ck.RootJs)
	log.Printf("Processing root: %s\ntag date: %s\n", ck.RootJs, ck.TagDate.String())
	// to add a custom root name use `treeprint.NewWithRoot()` instead
	//treeRoot := treeprint.New()
	tree := treeprint.New()
	tree.AddNode(rootName)
	// tree.AddNode("Dockerfile")
	// tree.AddNode("Makefile")
	// tree.AddNode("aws.sh")
	// tree.AddMetaBranch(" 204", "bin").
	// 	AddNode("dbmaker").AddNode("someserver").AddNode("testtool")
	// tree.AddMetaBranch(" 374", "deploy").
	// 	AddNode("Makefile").AddNode("bootstrap.sh")
	// tree.AddMetaNode("122K", "testtool.a")

	fmt.Println(tree.String())

	return nil
}
