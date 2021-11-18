package main

import (
	"fmt"
	"net/http"

	"github.com/akamensky/argparse"
	"github.com/spf13/cobra"
)

var p = argparse.NewParser("fs", "")

var RootCmd = &cobra.Command{
	Use:   "fs",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path := p.String("d", "dir", &argparse.Options{Default: "."})
		port := p.Int("p", "port", &argparse.Options{Default: 8888})
		if err := p.Parse(append(append([]string{"fs"}, args...))); err != nil {
			cmd.PrintErrln(err)
			_ = cmd.Help()
		}

		fs := http.FileServer(http.Dir(*path))
		if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), fs); err != nil && err != http.ErrServerClosed {
			cmd.PrintErrln(err)
			_ = cmd.Help()
		}
	},
}

func main() {
	cobra.CheckErr(RootCmd.Execute())
}
