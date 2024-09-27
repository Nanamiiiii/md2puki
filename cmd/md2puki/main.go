package main

import (
    "flag"
	"fmt"
	"os"
    
    "github.com/yuin/goldmark"
    "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/text"
	"github.com/Nanamiiiii/md2puki/pkg/renderer"
)

type Options struct {
    outfile string
    inputfile string
}

func parse(b []byte) ast.Node {
	md := goldmark.New(
		goldmark.WithExtensions(extension.NewTable()),
	)

	return md.Parser().Parse(text.NewReader(b))
}

func processOptions() *Options {
    var opts Options

    flag.StringVar(&opts.outfile, "out", "", "Output filename.")
    flag.Parse()

    args := flag.Args()
    if len(args) != 1 {
        fmt.Fprintln(os.Stderr, "Invalid argument. Specify only one input filename.")
        os.Exit(1)
    }

    opts.inputfile = args[0]

    return &opts
}

func main() {
    opts := processOptions()

    bytes, err := os.ReadFile(opts.inputfile)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Cannot read file: ", opts.inputfile);
    }

    if opts.outfile != "" { 
        fout, err := os.Create(opts.outfile)
        if err != nil {
            fmt.Fprintln(os.Stderr, "Cannot create output file: ", opts.outfile);
        }

        defer fout.Close()

        r := renderer.NewRenderer()
        r.Render(fout, bytes, parse(bytes))
    } else {
        r := renderer.NewRenderer()
        r.Render(os.Stdout, bytes, parse(bytes))
    }
}
