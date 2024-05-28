package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	i "gime/internal"
)

const VERSION = "v0.0.1"

var filePath string
var media i.Media

var root = &cobra.Command{
	Use:   "gime",
	Short: "Gime is a simple program to display web image and youtube link",
	Example: `use file flag: gime --mime=image/png --file=path/to/file.png
  or pipe from stdin: echo image/link.png | gime --mime=image/png `,
	Version: VERSION,
	Long: `
  GIME
  ==========================================================================
  A simple program to display images and youtube links from the command line`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := inputHandler(); err != nil {
			fmt.Println(err.Error())
			cmd.Help()
			os.Exit(0)
		}
		if err := media.MediaProcess(); err != nil {
			fmt.Printf("%s", err.Error())
			os.Exit(1)
		}
		fmt.Printf("media link: %s\nmime type: %s", media.Link, media.Mime)
		app := NewApp(&media)
		app.Run()
	},
}

func init() {
	root.Flags().StringVarP(&media.Mime, "mime", "m", "", "input mime types")
	root.Flags().StringVarP(&filePath, "file", "f", "", "path to the file")

	root.Flags().StringVarP(&media.Mime, "img", "i", string(i.I_wildcard), "shorthand for image/*")
	root.Flags().StringVarP(&media.Mime, "vid", "v", string(i.V_youtube), "shorthand for video/x-youtube")
}

func Excute() {
	root.ParseFlags([]string{"mime", "file"})

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func inputHandler() error {
	var err error
	if i.IsInputFromPipe() {
		//NOTE: PIPE
		if err = i.ReadFromPipe(os.Stdin, &media.Link); err != nil {
			return fmt.Errorf("%s", err.Error())
		}
	} else {
		//NOTE: file
		var file *os.File
		if file, err = i.GetFile(filePath); err != nil {
			return fmt.Errorf("%s", err.Error())
		}

		defer file.Close()

		if err = i.ReadFromFile(file, &media.Link); err != nil {
			return fmt.Errorf("%s", err.Error())
		}
	}
	return nil
}
