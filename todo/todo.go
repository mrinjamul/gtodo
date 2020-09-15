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
package todo

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

// Item : structure
type Item struct {
	Text     string
	Priority int
	position int
	Done     bool
}

var version string = "v0.2.0"

// GetVersion : returns version info
func GetVersion() string {
	// file, _ := ioutil.ReadFile("version.txt")
	// var version string
	// version = string(file)
	// version = strings.TrimSuffix(version, "\n")
	return version
}

// SaveItems : save data
func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

// ReadItems : read data
func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}
	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}
	for i := range items {
		items[i].position = i + 1
	}
	return items, nil
}

// SetPriority : sets priority to todo
func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

// PrettyP : prettify priority list
func (i *Item) PrettyP() string {
	if i.Priority == 1 {
		return "(1)"
	}
	if i.Priority == 3 {
		return "(3)"
	}

	return " "
}

// Label : index lists
func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
}

// ByPri implements sort.Interface for []Item base on
// the priority & position field.
type ByPri []Item

func (s ByPri) Len() int      { return len(s) }
func (s ByPri) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPri) Less(i, j int) bool {
	if s[i].Done != s[j].Done {
		return s[j].Done
	}

	if s[i].Priority != s[j].Priority {
		return s[i].Priority < s[j].Priority
	}
	return s[i].position < s[j].position
}

// PrettyDone : prettify
func (i *Item) PrettyDone() string {
	if i.Done {
		return "X"
	}
	return ""
}

// RemoveItem removes todo from list
func RemoveItem(slice []Item, s int) []Item {
	return append(slice[:s], slice[s+1:]...)
}
