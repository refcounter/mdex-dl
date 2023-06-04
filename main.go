package main

import (
	"errors"
	"fmt"

	"github.com/leaanthony/clir"
)

type Args struct {
  Url     string  `name:"url"  description:"Manga's url from https://mangadex.org"`
  Lang    string  `name:"lang" description:"Translator language" default:"en"`
  Dir     string  `name:"dir"  description:"Custom Download Directory" default:"."`
  Start   int     `name:"sc"   description:"Start Chapter" default:1`
  End     int     `name:"ec"   description:"End Chapter" default:0`
  Saver   bool    `name:"ds"   description:"Data Saver mode" default:"true"`
  Single  bool    `name:"s"    description:"Link is Single Chapter?" default:"false"`
}


func runCli() {
  // Create new cli
	cli := clir.NewCli("mdex-dl", "A simple mangadex downloader", "v0.1")

  args := &Args{}
  cli.AddFlags(args)

	// Define action for the command
	cli.Action(func() error {
    if args.Start < 0 || args.End < 0 || args.Url == "" {
      return errors.New("Invalid Arguments")
    }

    if args.Single {
      SingleDownload(args.Url, args.Saver)
    } else {
      DownloadManga(args.Url, args.Lang ,args.Start, args.End, args.Saver)
    }

		return nil
	})

	if err := cli.Run(); err != nil {
		fmt.Printf("Error encountered: %v\n", err)
	}
}
func main() {
  runCli()
}
