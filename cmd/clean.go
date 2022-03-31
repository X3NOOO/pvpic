/*
Copyright © 2022 X3NO <X3NO@disroot.org> [https://github.com/X3NOOO]

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"io/ioutil"
	"os"

	"github.com/X3NOOO/logger"
	"github.com/X3NOOO/pvpic/pvpic"
	"github.com/X3NOOO/pvpic/utils"
	"github.com/spf13/cobra"
)

var (
	files []string
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean metadata",
	Long: `remove metadata from your images`,
	Run: func(cmd *cobra.Command, args []string) {
		clean(args)
	},
}

/*
 * 1. Read image from path in argument
 * 2. Remove metadata from image
 * 3. Write image to {files}.clean.{files[extension]}
 */
func clean(args []string){
	// configurate logger
	l := logger.NewLogger("clean.go")
	l.SetVerbosity(Verbose)
	l.Debugln("Verbosity:", Verbose)

	// get existing files from args
	files, err := utils.CheckFiles(args)
	if(err != nil){
		l.Fatalln(1, "while checking files:", err)
	}

	// read all files and pass it to cleaner, then write cleaner output to {files}.clean.{files[extension]}
	for _, file := range files {
		l.Debugln("file:", file)
	
		// read file content
		fileContent, err := os.ReadFile(file)
		if err != nil {
			l.Fatalln(1, "Could not read file content", file)
		}

		// remove metadata
		cleaned, err := pvpic.Clean(fileContent)
		if err != nil {
			l.Fatalln(1, "Could not clean file", file)
		}

		// get extension of file
		new_name := utils.AddBeforeDot(file, "_clean")
		l.Debugln("new name:", new_name)

		// write cleaned file
		if !Testing {
			err = ioutil.WriteFile(new_name, cleaned, 0644)
			if err != nil {
				l.Fatalln(1, "Could not write file")
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
