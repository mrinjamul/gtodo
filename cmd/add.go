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
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mrinjamul/gtodo/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a new todo",
	Long:    `Add will create a new todo item to the list`,
	Run:     addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: gtodo add [task]")
		log.Fatalln("Too short argument")
	}
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		file := []byte("[]")
		err = ioutil.WriteFile(dataFile, file, 0644)
	}
	var todoName string
	for _, x := range args {
		todoName += x + " "
	}
	item := todo.Item{Text: todoName}
	item.SetPriority(priority)
	items = append(items, item)
	err = todo.SaveItems(dataFile, items)
	fmt.Println("+ new todo added")
	if err != nil {
		fmt.Println(err)
	}
}

var priority int

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")
}
