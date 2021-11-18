package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	Path string
	Port int
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&Path, "dir", "d", "./", "server dir")
	RootCmd.PersistentFlags().IntVarP(&Port, "port", "p", 8989, "server port")
}

var RootCmd = &cobra.Command{
	Use:   "fs",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var dir http.Dir
		if filepath.IsAbs(Path) {
			dir = http.Dir(Path)
		} else {
			pwd, err := os.Getwd()
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
			dir = http.Dir(filepath.Join(pwd, Path))
		}

		addr := fmt.Sprintf(":%d", Port)
		if err := http.ListenAndServe(addr, http.FileServer(dir)); err != nil && err != http.ErrServerClosed {
			cmd.PrintErrln(err)
			return
		}
	},
}

func main() {
	cobra.CheckErr(RootCmd.Execute())
}
