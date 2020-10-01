/*Package cmd ...
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
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/mrinjamul/gtodo/todo"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove a todo",
	Long:    `Remove will remove a todo item from the list by Label(index)`,
	Run:     removeRun,
}

// remove subcommad func
func removeRun(cmd *cobra.Command, args []string) {
	// remove only done tasks func
	if doOpt {
		response := todo.ConfirmPrompt("Do you want to remove all undone task(s)?")
		if response {
			items, err := todo.ReadItems(dataFile)
			var undoneItems []todo.Item
			if err != nil {
				fmt.Println("Something went wrong!")
				os.Exit(0)
			}
			for i, item := range items {
				if !items[i].Done {
					undoneItems = append(undoneItems, item)
				}
				if items[i].Done {
					text := items[i].Text
					fmt.Println("- " + "\"" + strconv.Itoa(i) + ". " + text + "\"" + " has been removed")
				}
			}
			sort.Sort(todo.ByPri(undoneItems))
			todo.SaveItems(dataFile, undoneItems)
		}
	} else { // Remove one by one
		if len(args) == 0 {
			fmt.Println("Usage: gtodo remove [tasks id]")
			log.Fatalln("Too short argument")
		}
		items, err := todo.ReadItems(dataFile)
		i, err := strconv.Atoi(args[0])

		if err != nil {
			log.Fatalln(args[0], "is not a valid label\ninvalid syntax")
		}
		if i > 0 && i <= len(items) {
			text := items[i-1].Text
			items = todo.RemoveItem(items, i-1)
			fmt.Println("- " + "\"" + strconv.Itoa(i) + ". " + text + "\"" + " has been removed")
			sort.Sort(todo.ByPri(items))
			todo.SaveItems(dataFile, items)
		} else {
			log.Println(i, "doesn't match any items")
		}
	}
}

// flag variables
var (
	doOpt bool
)

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	removeCmd.Flags().BoolVarP(&doOpt, "done", "d", false, "remove only done tasks")
}
