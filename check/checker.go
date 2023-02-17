package check

import (
	"fmt"
	"log"
	"time"

	"github.com/xlab/treeprint"
)

type Checker struct {
	RootJs  string
	TagDate time.Time
}

func (ck *Checker) Process() error {
	log.Printf("Processing root: %s\ntag date: %s\n", ck.RootJs, ck.TagDate.String())
	// to add a custom root name use `treeprint.NewWithRoot()` instead
	tree := treeprint.New()

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
