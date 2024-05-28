package cmd

import (
  "os"
  "fmt"
  "embed"

  "github.com/spf13/cobra"
	i "gime/internal"
)

const VERSION = "v0.0.1"

type Assets struct { Asset embed.FS }
var filePath string
var media i.Media
var AssetMap Assets

var root = &cobra.Command{
	Use:   "gime",
	Short: "Gime is a simple program to display web image and youtube link",
	Long:  `A simple program to display images and youtube links from the command line`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := InputHanlder(); err != nil {
			fmt.Printf("%s", err.Error())
			os.Exit(0)
		}
		if err := media.MediaProcess(); err != nil {
			fmt.Printf("%s", err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s - %s\n", media.Mime, media.Link)
		RunGUI()
	},
}

var version = &cobra.Command{
	Use:   "version",
	Short: "Display app version",
	Long:  `Display app version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gime version: %s", VERSION)
		os.Exit(0)
	},
}

func init() {
	root.AddCommand(version)
}

func Excute() {
	root.Flags().StringVarP(&media.Mime, "mime", "m", "", "input mime types")
	root.Flags().StringVarP(&filePath, "file", "f", "", "path to the file")
	// root.MarkFlagRequired("mime")
	root.ParseFlags([]string{"mime", "file"})

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
