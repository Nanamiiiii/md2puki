package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

    mathjax "github.com/litao91/goldmark-mathjax"
	"github.com/Nanamiiiii/md2puki/pkg/renderer"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/text"
)

type Options struct {
    outfile string
    inputfile string
}

func parse(b []byte) ast.Node {
	md := goldmark.New(
		goldmark.WithExtensions(extension.NewTable(), mathjax.NewMathJax()),
	)

	return md.Parser().Parse(text.NewReader(b))
}

func processOptions() *Options {
    var opts Options

    flag.StringVar(&opts.outfile, "out", "", "Output filename.")
    flag.StringVar(&opts.inputfile, "in", "", "Input filename.")
    flag.Parse()

    return &opts
}

func main() {
    opts := processOptions()
    var bytes []byte
    var err error

    if opts.inputfile != "" {
        bytes, err = os.ReadFile(opts.inputfile)
        if err != nil {
            fmt.Fprintln(os.Stderr, "Cannot read file: ", opts.inputfile);
        }
    } else {
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Split(bufio.ScanBytes)
        for scanner.Scan() {
            bytes = append(bytes, scanner.Bytes()...)
        }
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
