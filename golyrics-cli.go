package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/mamal72/golyrics"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "golyrics-cli"
	app.Usage = "A command line client for golyrics"
	app.UsageText = "golyrics-cli \"Artist:Song\" | Example: golyrics-cli \"Blackfield:Some Day\""
	app.Version = "1.0.0"

	app.Action = func(c *cli.Context) error {
		if len(c.Args()) != 0 {
			suggestions, err := golyrics.SearchTrack(c.Args().First())
			if err != nil || len(suggestions) == 0 {
				return cli.NewExitError("Couldn't find your requested song", 10)
			}

			track := suggestions[0]
			err = track.FetchLyrics()

			if err != nil {
				panic(err)
			}
			color.Set(color.FgCyan)
			fmt.Printf("\"%s\" by \"%s\":\n\n", track.Name, track.Artist)
			color.Unset()
			fmt.Print(track.Lyrics)
		} else {
			fmt.Println("Please give a track - For more information check --help")
		}
		return nil

	}
	app.Run(os.Args)
}
