package kcs

import (
	"sort"

	"github.com/fatih/color"
)

var (
	verbose         = false
	categoryColor   = color.New(color.FgCyan)
	categoryHiColor = color.New(color.FgHiCyan)
	commandColor    = color.New(color.FgMagenta)
	commandHiColor  = color.New(color.FgHiMagenta)
	argumentColor   = color.New(color.FgYellow)
	argumentHiColor = color.New(color.FgHiYellow)
)

func SetVerbose(v bool) {
	verbose = v
}

type CheatSheet struct {
	Categories map[string]Category
}

func (c CheatSheet) Sort() []Category {
	var keys = make([]string, len(c.Categories))
	var i = 0
	for key := range c.Categories {
		keys[i] = key
		i++
	}

	sort.Strings(keys)
	var categories []Category
	for _, key := range keys {
		categories = append(categories, c.Categories[key])
	}
	return categories
}

func (c CheatSheet) Print(category, command string) {
	var innerC = c
	if category != "" {
		if pickedCategory, ok := c.Categories[category]; ok {
			innerC = CheatSheet{Categories: make(map[string]Category)}
			innerC.Categories[category] = pickedCategory
		}
	}
	categories := innerC.Sort()

	var counter = 0
	for _, singleCategory := range categories {
		counter++
		singleCategory.Print(counter == 1)
	}
}

type Category struct {
	Name        string
	Description string
	Commands    map[string]CommandDescriptor
}

func (c Category) Sort() []CommandDescriptor {
	var keys = make([]string, len(c.Commands))
	var i = 0
	for key := range c.Commands {
		keys[i] = key
		i++
	}

	sort.Strings(keys)
	var commands []CommandDescriptor
	for _, key := range keys {
		commands = append(commands, c.Commands[key])
	}
	return commands
}

func (c Category) Print(first bool) {
	c.header(first)

	sorted := c.Sort()
	for _, cd := range sorted {
		cd.Print()
	}
}

func (c Category) header(first bool) {
	if !first {
		_, _ = categoryColor.Print("\n\n")
	}
	_, _ = categoryColor.Println(charMultiplier(len(c.Name)+2, "="))
	_, _ = categoryHiColor.Println(c.Name + " |")
	_, _ = categoryColor.Println(charMultiplier(len(c.Name)+2, "="))
	if verbose && len(c.Description) > 0 {
		_, _ = categoryColor.Println(c.Description)
	}
	_, _ = categoryColor.Print("\n")
}

type CommandDescriptor struct {
	Command     string
	Args        []ArgumentDescriptor
	Description string
}

func (cd *CommandDescriptor) Print() {
	_, _ = commandHiColor.Printf("$ %s", cd.Command)
	if verbose && cd.Description != "" {
		_, _ = commandColor.Printf(" - %s", cd.Description)
	}
	_, _ = commandColor.Print("\n")
	for _, argument := range cd.Args {
		argument.Print(len(cd.Command) + 3)
	}
}

type ArgumentDescriptor struct {
	Argument    string
	Description string
}

func (ad *ArgumentDescriptor) Print(tabsize int) {
	if verbose {
		tabsize = 4
	}
	tab := charMultiplier(tabsize, " ")
	_, _ = argumentHiColor.Printf("%s%s", tab, ad.Argument)
	if verbose && ad.Description != "" {
		_, _ = argumentColor.Printf(" - %s", ad.Description)
	}
	_, _ = argumentColor.Print("\n")
}
