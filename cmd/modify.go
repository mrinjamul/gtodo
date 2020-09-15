/*
Copyright © 2020 Injamul Mohammad Mollah

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"log"
	"sort"
	"strconv"

	"github.com/mrinjamul/gtodo/todo"
	"github.com/spf13/cobra"
)

// modifyCmd represents the modify command
var modifyCmd = &cobra.Command{
	Use:     "modify",
	Aliases: []string{"mod", "ed"},
	Short:   "edit a todo",
	Long:    ``,
	Run:     modifyRun,
}

func modifyRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatalln("Too short argument")
	} else if len(args) == 1 {
		log.Fatalln("todo name is empty")
	} else if len(args) > 2 {
		log.Fatalln("Too much arguments")
	}
	items, err := todo.ReadItems(dataFile)
	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid label\ninvalid syntax")
	}

	text := args[1]
	if i > 0 && i <= len(items) {
		items[i-1].Text = text
		sort.Sort(todo.ByPri(items))
		todo.SaveItems(dataFile, items)
	} else {
		log.Println(i, "doesn't match any items")
	}

}

func init() {
	rootCmd.AddCommand(modifyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
