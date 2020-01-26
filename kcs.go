package kcs

import (
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	verbose = false
	categoryColor = color.New(color.FgCyan)
	categoryHiColor = color.New(color.FgHiCyan)
	commandColor = color.New(color.FgMagenta)
	commandHiColor = color.New(color.FgHiMagenta)
	argumentColor = color.New(color.FgYellow)
	argumentHiColor = color.New(color.FgHiYellow)
)

func SetVerbose(v bool) {
	verbose = v
}

type CheatSheet struct {
	Categories map[string]Category
}

func (c *CheatSheet) Print(category, command string) {
	width, _, _ := terminal.GetSize(0)
	_, _ = categoryColor.Println(charMultiplier(width, "="))

	var categories = c.Categories
	if category != "" {
		if pickedCategory, ok := c.Categories[category]; ok {
			categories = make(map[string]Category)
			categories[category] = pickedCategory
		}
	}

	var counter = 0
	for _, singleCategory := range categories {
		counter++
		singleCategory.Print(len(categories)==counter)
	}
}

type Category struct {
	Name string
	Description string
	Commands map[string]CommandDescriptor
}

func (c *Category) Print(last bool) {
	//_, _ = categoryHiColor.Println(charMultiplier(len(c.Name), "-"))
	_, _ = categoryHiColor.Println(c.Name + " |")
	_, _ = categoryColor.Println(charMultiplier(len(c.Name)+2, "="))
	if verbose {
		_, _ = categoryColor.Println(c.Description)
	}
	_, _ = categoryColor.Print("\n")
	for _, cd := range c.Commands { // @TODO alphabetic order
		cd.Print()
	}
	if !last {
		width, _, _ := terminal.GetSize(0)
		_, _ = categoryColor.Print("\n")
		_, _ = categoryColor.Println(charMultiplier(width, "="))
		//_, _ = categoryColor.Print("\n")
	}
}

type CommandDescriptor struct {
	Command string
	Args []ArgumentDescriptor
	Description string
}

func (cd *CommandDescriptor) Print() {
	_, _ = commandHiColor.Printf("$ %s", cd.Command)
	if verbose && cd.Description != "" {
		_, _ = commandColor.Printf(" - %s", cd.Description)
	}
	_, _ = commandColor.Print("\n")
	for _, argument := range cd.Args {
		argument.Print(len(cd.Command)+3)
	}
}

type ArgumentDescriptor struct {
	Argument string
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



