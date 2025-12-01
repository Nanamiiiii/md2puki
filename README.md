# md2puki

Markdown to pukiwiki notation converter

## Installation
### Pre-built binary
You can download from [Releases](https://github.com/Nanamiiiii/md2puki/releases).
- Supported OS
    - Windows
    - macOS
    - Linux
- Supported Architecture
    - amd64
    - arm64

### Nix
```
# Use flakes
nix run github:Nanamiiiii/md2puki#md2puki

# Use NUR
nix-shell -p nur.repos.nanamiiiii.md2puki
```

## Usage
### Options
- `-in`: input Markdown file
- `-out`: output filename of pukiwiki notation

### from file to stdout
```
$ md2puki -in sample.md
* hoge
** fuge
- foo
-- bar
```

### from file to file
```
$ md2puki -in sample.md -out result.puki
```

### from stdin to file
```
$ md2puki -out result.puki
# hoge
## fuga
- foo
  - bar
(EOF)
```

### from stdin to stdout
```
$ md2puki
# hoge
## fuga
- foo
  - bar
(EOF)

* hoge
** fuge
- foo
-- bar
```

### heredoc
```
$ cat << EOF | md2puki
# hoge
## fuga
- foo
  - bar
EOF

* hoge
** fuga
- foo
-- bar
```

## Markdown Compatibility
`md2puki` is now supporting following markdown elements.
- Text
- Headings
    - pukiwiki natively supports H1~H3. After the H4, `md2puki` uses emphasized text.
- List / Unordered List
    - Support up to 3-levels of nesting.
- Code Block
- Inline Code
    - Covert to emphasized text.
- Link
- Image
- Table
    - Now Support Partially
- Block Quote
- Bold
- Italic
- Strikethrough
- Math Block / Inline Math
    - Need to install mathjax plugin to pukiwiki
- Frontmatter
    - Ignored

Some of elements of Obsidian flavor are supported.
- WikiLink


Original README below
---

# md2puki

Markdown to Pukiwiki converter

# License
- Under the MIT License
- Copyright (c) 2021 Tsuzu
