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
	"github.com/X3NOOO/pvpic/values"
	"github.com/spf13/cobra"
)

// fakeCmd represents the fake command
var fakeCmd = &cobra.Command{
	Use:   "fake",
	Short: "fake metadata",
	Long:  `add fake metadata to your images`,
	Run: func(cmd *cobra.Command, args []string) {
		fake(args)
	},
}

var (
	model 			string
	model_folder 	string
)

/*
 * 1. Read image from path in args
 * 2. Check if files in args exists
 * 3. Remove metadata from image
 * 4. Read read {string_after_-m_or_--model}.json from db path
 * 5. Add fake metadata from model to image
 * 6. Write image to {files}.fake.{files[extension]}
 */
func fake(args []string) {
	// configurate logger
	l := logger.NewLogger("fake.go")
	l.SetVerbosity(Verbose)
	l.Debugln("Verbosity:", Verbose)

	// get existing files from args
	files, err := utils.CheckFiles(args)
	if err != nil {
		l.Fatalln(1, "while checking files:", err)
	}

	// read all files and pass it to cleaner, then write cleaner output to {files}.clean.{files[extension]}
	for _, file := range files {
		l.Debugln("file:", file)

		// read file content
		file_content, err := os.ReadFile(file)
		if err != nil {
			l.Warningln(1, "Could not read file content:", file)
		}

		// remove metadata
		cleaned, err := pvpic.Clean(file_content)
		if err != nil {
			l.Warningln(1, "Could not clean file:", file)
		}
		l.Debugln("cleaned:", cleaned)

		// check if model folder is valid
		if(model_folder[len(model_folder)-1:] != "/"){
			model_folder += "/"
		}
		model_path := model_folder + model + ".json"
		l.Debugln("model_path:", model_path)

		// read model
		model, err := utils.ReadModel(model_path)
		if(err != nil){
			l.Fatalln(1, "while reading model:", err)
		}

		l.Debugln("model of metadata:", model)
		

		// add fake metadata from model to image
		faked, err := pvpic.Fake(cleaned, model)
		if(err != nil){
			l.Fatalln(1, "while adding fake metadata:", err)
		}
		l.Debugln("faked:", faked)

		// get extension of file
		new_name := utils.AddBeforeDot(file, "_clean")
		l.Debugln("new name:", new_name)

		// write cleaned file
		if !Testing {
			err = ioutil.WriteFile(new_name, faked, 0644)
			if err != nil {
				l.Warningln(1, "Could not write file:", err)
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(fakeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fakeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fakeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	fakeCmd.Flags().StringVarP(&model, "model", "m", "none", "select model to use metadata from")
	fakeCmd.Flags().StringVarP(&model_folder, "database", "d", values.HOME + "/.config/pvpic/db/", "database path")
	fakeCmd.MarkFlagRequired("model")
}
