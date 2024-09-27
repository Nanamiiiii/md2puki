package main

import (
	"fmt"
	"os"
    
    "github.com/yuin/goldmark"
    "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/text"
	"github.com/Nanamiiiii/md2puki/pkg/renderer"
)

func parse(b []byte) ast.Node {
	md := goldmark.New(
		goldmark.WithExtensions(extension.NewTable()),
	)

	return md.Parser().Parse(text.NewReader(b))
}

func main() {
    args := os.Args
    if len(args) != 2 {
        fmt.Fprintln(os.Stderr, "Invalid argument. Specify only single filename.")
        os.Exit(1)
    }

    bytes, err := os.ReadFile(args[1])
    if err != nil {
        fmt.Fprintln(os.Stderr, "Cannot read file: ", args[1]);
    }

    outname := args[1] + ".puki"
    fout, err := os.Create(outname)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Cannot create output file: ", outname);
    }

    defer fout.Close()

    r := renderer.NewRenderer()
    r.Render(fout, bytes, parse(bytes))
}
