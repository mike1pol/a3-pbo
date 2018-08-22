package main

import (
	"flag"
	pbo "github.com/TheMysteriousVincent/go-pbo"
)

var (
	from   = flag.String("from", "", "directory from")
	to     = flag.String("to", "file.pbo", "to pbo file")
	prefix = flag.String("prefix", "", "Password")
)

func main() {
	flag.Parse()
	if *from == "" {
		flag.Usage()
		return
	}

	nw := pbo.New()
	nw.From = *from
	nw.To = *to
	nw.Prefix = *prefix
	nw.Generate()
	nw.Save()

}
