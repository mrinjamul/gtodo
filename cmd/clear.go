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

	"github.com/mrinjamul/gtodo/todo"
	"github.com/spf13/cobra"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all todos",
	Long:  ``,
	Run:   clearRun,
}

func clearRun(cmd *cobra.Command, args []string) {
	items := []todo.Item{}
	if forceOpt == true {
		err := todo.SaveItems(dataFile, items)
		if err != nil {
			fmt.Printf("%v", err)
		}
		fmt.Println("Warning: All todo has been cleared.")
	} else {
		response := todo.ConfirmPrompt("Do you want to clear all todos")
		if response {
			err := todo.SaveItems(dataFile, items)
			if err != nil {
				fmt.Printf("%v", err)
			}
			fmt.Println("Warning: All todo has been cleared.")
		} else {
			fmt.Println("Operation Canceled")
		}

	}
}

var (
	forceOpt bool
)

func init() {
	rootCmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	clearCmd.Flags().BoolVarP(&forceOpt, "force", "f", false, "Force remove all todos")
}
