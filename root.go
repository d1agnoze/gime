package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	i "gime/internal"
)

const VERSION = "v0.0.1"

var filePath string
var media i.Media
var isImageShortCut bool
var isVideoShortCut bool
var position string = ""

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

	PreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().Lookup("mime").Changed &&
			(cmd.Flags().Lookup("img").Changed || cmd.Flags().Lookup("vid").Changed) {
			return fmt.Errorf("Cannot use --mime and --img or --vid together")
		}
		if i.IsInputFromPipe() && cmd.Flags().Lookup("file").Changed {
			return fmt.Errorf("Cannot use --file and STDIN together")
		}
		if cmd.Flags().Lookup("img").Changed && cmd.Flags().Lookup("vid").Changed {
			return fmt.Errorf("Cannot use --img and --vid together")
		}

		if isImageShortCut {
			media.Mime = string(i.I_wildcard)
		}
		if isVideoShortCut {
			media.Mime = string(i.V_youtube)
		}
		position = strings.TrimSpace(position)
		position = strings.ToUpper(position)

		return nil
	},
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
		fmt.Printf("media link: %s\nmime type: %s\n", media.Link, media.Mime)
		app := NewApp(&media)
		app.Run()
	},
}

func init() {
	root.Flags().StringVarP(&media.Mime, "mime", "m", "", "input mime types")
	root.Flags().StringVarP(&filePath, "file", "f", "", "path to the file")

	root.Flags().BoolVarP(&isImageShortCut, "img", "i", false, "shorthand for image/*")
	root.Flags().BoolVarP(&isVideoShortCut, "vid", "v", false, "shorthand for video/x-youtube")

	root.Flags().StringVarP(&position, "position", "p", "", "window position (TL, TR, BL, BR)")
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
