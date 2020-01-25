package kcs

type cheatSheetPrinter func(string)

type CheatSheet struct {
	Categories map[string]Category
}

func (c *CheatSheet) Print(category, command string) {

}

type Category struct {
	Name string
	Description string
	Commands map[string]CommandDescriptor
}

type CommandDescriptor struct {
	Command string
	Args []ArgumentDescriptor
	Description string
}

type ArgumentDescriptor struct {
	Argument string
	Description string
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
			},
		},
	},
}

