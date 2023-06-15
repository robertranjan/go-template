# ${APPNAME}

template for go-cli projects

## Use snippet to generate MD document for all commands

    err := doc.GenMarkdownTree(kubectl, "./")
    if err != nil {
      log.Fatal(err)
    }

Markdown doc will be at /tmp/kubectl.md

[Read more here](https://github.com/spf13/cobra/blob/main/doc/md_docs.md)

## Use snippet to generate man page

    func main() {
      cmd := &cobra.Command{
        Use:   "test",
        Short: "my test program",
      }
      header := &doc.GenManHeader{
        Title: "MINE",
        Section: "3",
      }
      err := doc.GenManTree(cmd, header, "/tmp")
      if err != nil {
        log.Fatal(err)
      }
    }

Man page will be at /tmp/test.man

[Read more here](https://github.com/spf13/cobra/blob/main/doc/man_docs.md)
