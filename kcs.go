package kcs

import (
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	verbose = true
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
	_, _ = categoryHiColor.Println(charMultiplier(len(c.Name), "-"))
	_, _ = categoryHiColor.Println(c.Name)
	_, _ = categoryHiColor.Println(charMultiplier(len(c.Name), "-"))
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
		_, _ = categoryColor.Print("\n")
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

var Instance = CheatSheet{
	Categories: map[string]Category{
		"config": {
			Name: "Kubectl Context and Configuration",
			Description: "Set which Kubernetes cluster kubectl communicates with and modifies configuration information. See Authenticating Across Clusters with kubeconfig documentation for detailed config file information.",
			Commands: map[string]CommandDescriptor{
				"view": {
					Command: "kubectl config view",
					Args: []ArgumentDescriptor{
						{
							Argument: `-o jsonpath='{.users[?(@.name == "e2e")].user.password}'`,
							Description: "get the password for the e2e user",
						},
						{
							Argument: `-o jsonpath='{.users[].name}'`,
							Description: "display the first user",
						},
						{
							Argument: `-o jsonpath='{.users[*].name}'`,
							Description: "get a list of users",
						},
					},
					Description: "Show Merged kubeconfig settings.",
				},
				"get-contexts": {
					Command: "kubectl config get-contexts",
					Description: "display list of contexts",
				},
				"current-context": {
					Command: "kubectl config current-contexts",
					Description: "display the current-context",
				},
				"use-context": {
					Command: "kubectl config use-context my-cluster-name",
					Description: "set the default context to my-cluster-name",
				},
				"set-credentials": {
					Command: "kubectl config set-credentials kubeuser/foo.kubernetes.com --username=kubeuser --password=kubepassword",
					Description: "add a new cluster to your kubeconf that supports basic auth",
				},
				"set-context": {
					Command: "kubectl config set-context",
					Args: []ArgumentDescriptor{
						{
							Argument: "--current --namespace=ggckad-s2",
							Description: "permanently save the namespace for all subsequent kubectl commands in that context",
						},
						{
							Argument: "gce --user=cluster-admin --namespace=foo",
							Description: "set a context utilizing a specific username and namespace",
						},
					},
				},
				"unset": {
					Command: "kubectl config unset users.fo",
					Description: "delete user foo",
				},
			},
		},
		"apply": {
			Name: "Apply",
			Description: "apply manages applications through files defining Kubernetes resources. It creates and updates resources in a cluster through running kubectl apply. This is the recommended way of managing Kubernetes applications on production.",
			Commands: map[string]CommandDescriptor{
				"apply": {
					Command: "kubectl apply",
					Args: []ArgumentDescriptor{
						{
							Argument: "-f ./my-manifest.yaml",
							Description: "create resource(s)",
						},
						{
							Argument: "-f ./my1.yaml -f ./my2.yaml",
							Description: "create from multiple files",
						},
						{
							Argument: "-f ./dir",
							Description: "create resource(s) in all manifest files in dir",
						},
						{
							Argument: "-f https://git.io/vPieo",
							Description: "create resource(s) from url",
						},
					},
				},
			},
		},
	},
}

