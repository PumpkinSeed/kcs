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
				"create": {
					Command: "kubectl create deployment nginx --image=nginx ",
					Description: "start a single instance of nginx",
				},
				"explain": {
					Command: "kubectl explain pods,svc",
					Description: "get the documentation for pod and svc manifests",
				},
			},
		},
		"get": {
			Name: "Viewing, Finding Resources",
			Commands: map[string]CommandDescriptor{
				"services": {
					Command: "kubectl get services",
					Description: "list all services in the namespace",
				},
				"pods-all": {
					Command: "kubectl get pods --all-namespaces",
					Description: "list all pods in all namespaces",
				},
				"pods-wide": {
					Command: "kubectl get pods -o wide",
					Description: "list all pods in the namespace, with more details",
				},
				"deployment": {
					Command: "kubectl get deployment my-dep",
					Description: "list a particular deployment",
				},
				"pods": {
					Command: "kubectl get pods",
					Description: "list all pods in the namespace",
				},
				"pod-yaml": {
					Command: "kubectl get pod my-pod -o yaml ",
					Description: "get a pod's YAML",
				},
				"pod-yaml-export": {
					Command: "kubectl get pod my-pod -o yaml --export",
					Description: "get a pod's YAML without cluster specific information",
				},
				"describe-nodes": {
					Command: "kubectl describe nodes my-node",
					Description: "describe commands with verbose output for nodes",
				},
				"describe-pods": {
					Command: "kubectl describe pods my-pod",
					Description: "describe commands with verbose output for pods",
				},
				"services-sort": {
					Command: "kubectl get services --sort-by=.metadata.name",
					Description: "list Services Sorted by Name",
				},
				"pods-sort": {
					Command: "kubectl get pods --sort-by='.status.containerStatuses[0].restartCount'",
					Description: "list pods Sorted by Restart Count",
				},
				"pv-sort": {
					Command: "kubectl get pv -n test --sort-by=.spec.capacity.storage",
					Description: "list PersistentVolumes in test namespace sorted by capacity",
				},
				"pods-selector": {
					Command: "kubectl get pods --selector=app=cassandra -o jsonpath='{.items[*].metadata.labels.version}'",
					Description: "get the version label of all pods with label app=cassandra",
				},
				"node-selector": {
					Command: "kubectl get node --selector='!node-role.kubernetes.io/master'",
					Description: "get all worker nodes (use a selector to exclude results that have a label named 'node-role.kubernetes.io/master')",
				},
				"pods-field-selector": {
					Command: "kubectl get pods --field-selector=status.phase=Running",
					Description: "get all running pods in the namespace",
				},
				"nodes-external-ip": {
					Command: `kubectl get nodes -o jsonpath='{.items[*].status.addresses[?(@.type=="ExternalIP")].address}'`,
					Description: "get ExternalIPs of all nodes",
				},
				"pods-rc": {
					Command: `sel=${$(kubectl get rc my-rc --output=json | jq -j '.spec.selector | to_entries | .[] | "\(.key)=\(.value),"')%?} && echo $(kubectl get pods --selector=$sel --output=jsonpath={.items..metadata.name})`,
					Description: "list Names of Pods that belong to Particular RC",
				},
				"pods-labels": {
					Command: "kubectl get pods --show-labels",
					Description: "show labels for all pods (or any other Kubernetes object that supports labelling)",
				},
				"nodes-ready": {
					Command: `JSONPATH='{range .items[*]}{@.metadata.name}:{range @.status.conditions[*]}{@.type}={@.status};{end}{end}' && kubectl get nodes -o jsonpath="$JSONPATH" | grep "Ready=True"`,
					Description: "check which nodes are ready",
				},
				"pod-secrets": {
					Command: "kubectl get pods -o json | jq '.items[].spec.containers[].env[]?.valueFrom.secretKeyRef.name' | grep -v null | sort | uniq",
					Description: "list all Secrets currently in use by a pod",
				},
				"events": {
					Command: "kubectl get events --sort-by=.metadata.creationTimestamp",
					Description: "list Events sorted by timestamp",
				},
				"diff": {
					Command: "kubectl diff -f ./my-manifest.yaml",
					Description: "compares the current state of the cluster against the state that the cluster would be in if the manifest was applied",
				},
			},
		},
		"update": {
			Name: "Updating Resources",
			Commands: map[string]CommandDescriptor{
				"set": {
					Command: "kubectl set image deployment/frontend www=image:v2",
					Description: `rolling update "www" containers of "frontend" deployment, updating the image`,
				},
				"rollout": {
					Command: "kubectl rollout",
					Args: []ArgumentDescriptor{
						{
							Argument:    "history deployment/frontend",
							Description: "check the history of deployments including the revision",
						},
						{
							Argument:    "undo deployment/frontend",
							Description: "rollback to the previous deployment",
						},
						{
							Argument:    "undo deployment/frontend --to-revision=2",
							Description: "rollback to a specific revision",
						},
						{
							Argument:    "status -w deployment/frontend",
							Description: `watch rolling update status of "frontend" deployment until completion`,
						},
						{
							Argument:    "restart deployment/frontend",
							Description: `rolling restart of the "frontend" deployment`,
						},
					},
				},
				"replace-json": {
					Command: "cat pod.json | kubectl replace -f -",
					Description: "replace a pod based on the JSON passed into std",
				},
				"replace-force": {
					Command: "kubectl replace --force -f ./pod.json",
					Description: "force replace, delete and then re-create the resource. Will cause a service outage",
				},
				"expose": {
					Command: "kubectl expose rc nginx --port=80 --target-port=8000",
					Description: "create a service for a replicated nginx, which serves on port 80 and connects to the containers on port 8000",
				},
				"pods-image": {
					Command: `kubectl get pod mypod -o yaml | sed 's/\(image: myimage\):.*$/\1:v4/' | kubectl replace -f -`,
					Description: "update a single-container pod's image version (tag) to v4",
				},
				"label": {
					Command: "kubectl label pods my-pod new-label=awesome",
					Description: "add a Label",
				},
				"annotate": {
					Command: "kubectl annotate pods my-pod icon-url=http://goo.gl/XXBTWq",
					Description: "add an annotation",
				},
				"autoscale": {
					Command: "kubectl autoscale deployment foo --min=2 --max=10",
					Description: `auto scale a deployment "foo"`,
				},
			},
		},
		"patch": {
			Name: "Patching Resources",
			Commands: map[string]CommandDescriptor{
				"node": {
					Command: `kubectl patch node k8s-node-1 -p '{"spec":{"unschedulable":true}}'`,
					Description: "partially update a node",
				},
				"pod": {
					Command: `kubectl patch pod valid-pod -p '{"spec":{"containers":[{"name":"kubernetes-serve-hostname","image":"new image"}]}}'`,
					Description: "update a container's image; spec.containers[*].name is required because it's a merge key",
				},
				"pod-array": {
					Command: `kubectl patch pod valid-pod --type='json' -p='[{"op": "replace", "path": "/spec/containers/0/image", "value":"new image"}]'`,
					Description: "update a container's image using a json patch with positional arrays",
				},
				"deployment": {
					Command: `kubectl patch deployment valid-deployment  --type json   -p='[{"op": "remove", "path": "/spec/template/spec/containers/0/livenessProbe"}]'`,
					Description: "disable a deployment livenessProbe using a json patch with positional arrays",
				},
				"sa": {
					Command: `kubectl patch sa default --type='json' -p='[{"op": "add", "path": "/secrets/1", "value": {"name": "whatever" } }]'`,
					Description: "add a new element to a positional array",
				},
			},
		},
		"edit": {
			Name: "Editing Resources",
			Commands: map[string]CommandDescriptor{
				"default": {
					Command: "kubectl edit svc/docker-registry",
					Description: "edit the service named docker-registry",
				},
				"nano": {
					Command: `KUBE_EDITOR="nano" kubectl edit svc/docker-registry`,
					Description: "use an alternative editor",
				},
			},
		},
		"scale": {
			Name: "Scaling Resources",
			Commands: map[string]CommandDescriptor{
				"replicas": {
					Command: `kubectl scale --replicas=3 rs/foo`,
					Description: `scale a replicaset named 'foo' to 3`,
				},
				"replicas-specific": {
					Command: `kubectl scale --replicas=3 -f foo.yaml`,
					Description: `scale a resource specified in "foo.yaml" to 3`,
				},
				"replicas-condition": {
					Command: `kubectl scale --current-replicas=2 --replicas=3 deployment/mysql`,
					Description: `if the deployment named mysql's current size is 2, scale mysql to 3`,
				},
				"replicas-multiple": {
					Command: `kubectl scale --replicas=5 rc/foo rc/bar rc/baz`,
					Description: `scale multiple replication controllers`,
				},
			},
		},
		"delete": {
			Name: "Deleting Resources",
			Commands: map[string]CommandDescriptor{
				"file": {
					Command: `kubectl delete -f ./pod.json`,
					Description: `delete a pod using the type and name specified in pod.json`,
				},
				"name": {
					Command: `kubectl delete pod,service baz foo`,
					Description: `delete pods and services with same names "baz" and "foo"`,
				},
				"label": {
					Command: `kubectl delete pods,services -l name=myLabel`,
					Description: `delete pods and services with label name=myLabel`,
				},
				"namespace": {
					Command: `kubectl -n my-ns delete pod,svc --all`,
					Description: `delete all pods and services in namespace my-ns`,
				},
				"pattern": {
					Command: `kubectl get pods  -n mynamespace --no-headers=true | awk '/pattern1|pattern2/{print $1}' | xargs  kubectl delete -n mynamespace pod`,
					Description: `delete all pods matching the awk pattern1 or pattern2`,
				},
			},
		},
		"logs": {
			Name: "Interacting with running Pods",
			Commands: map[string]CommandDescriptor{
				"pod": {
					Command: "kubectl logs my-pod",
					Description: "dump pod logs (stdout)",
				},
				"label": {
					Command: "kubectl logs -l name=myLabel",
					Description: "dump pod logs, with label name=myLabel (stdout)",
				},
				"prevoius": {
					Command: "kubectl logs my-pod --previous",
					Description: "dump pod logs (stdout) for a previous instantiation of a container",
				},
				"container": {
					Command: "kubectl logs my-pod -c my-container",
					Description: "dump pod container logs (stdout, multi-container case)",
				},
				"label-container": {
					Command: "kubectl logs -l name=myLabel -c my-container ",
					Description: "dump pod logs, with label name=myLabel (stdout)",
				},
				"container-prevoius": {
					Command: "kubectl logs my-pod -c my-container --previous",
					Description: "dump pod container logs (stdout, multi-container case) for a previous instantiation of a container",
				},
				"file": {
					Command: "kubectl logs -f my-pod",
					Description: "stream pod logs (stdout)",
				},
				"file-container": {
					Command: "kubectl logs -f my-pod -c my-container",
					Description: "stream pod container logs (stdout, multi-container case)",
				},
				"file-label": {
					Command: "kubectl logs -f -l name=myLabel --all-containers",
					Description: "stream all pods logs with label name=myLabel (stdout)",
				},
				"run": {
					Command: "kubectl run -i --tty busybox --image=busybox -- sh",
					Description: "run pod as interactive shell",
				},
				"run-namespace": {
					Command: "kubectl run nginx --image=nginx --restart=Never -n mynamespace",
					Description: "run pod nginx in a specific namespace",
				},
				"run-file": {
					Command: "kubectl run nginx --image=nginx --restart=Never --dry-run -o yaml > pod.yaml",
					Description: "run pod nginx and write its spec into a file called pod.yaml",
				},
				"attach": {
					Command: "kubectl attach my-pod -i",
					Description: "attach to Running Container",
				},
				"port-forward": {
					Command: "kubectl port-forward my-pod 5000:6000",
					Description: "listen on port 5000 on the local machine and forward to port 6000 on my-pod",
				},
				"exec": {
					Command: "kubectl exec my-pod -- ls",
					Description: "run command in existing pod (1 container case)",
				},
				"exec-multiple": {
					Command: "kubectl exec my-pod -c my-container -- ls",
					Description: "run command in existing pod (multi-container case)",
				},
				"top": {
					Command: "kubectl top pod POD_NAME --containers",
					Description: "show metrics for a given pod and its containers",
				},
			},
		},
		"master-api": {
			Name: "Interacting with Nodes and Cluster",
			Commands: map[string]CommandDescriptor{
				"cordon": {
					Command: "kubectl cordon my-node",
					Description: "mark my-node as unschedulable",
				},
				"drain": {
					Command: "kubectl drain my-node ",
					Description: "drain my-node in preparation for maintenance",
				},
				"uncordon": {
					Command: "kubectl uncordon my-node",
					Description: "mark my-node as schedulable",
				},
				"top": {
					Command: "kubectl top node my-node",
					Description: "show metrics for a given node",
				},
				"cluster-info": {
					Command: "kubectl cluster-info",
					Args: []ArgumentDescriptor{
						{
							Argument: "dump",
							Description: "dump current cluster state to stdout",
						},
						{
							Argument: "dump --output-directory=/path/to/cluster-state",
							Description: "dump current cluster state to /path/to/cluster-state",
						},
					},
					Description: "display addresses of the master and services",
				},
				"taint": {
					Command: "kubectl taint nodes foo dedicated=special-user:NoSchedule",
					Description: "if a taint with that key and effect already exists, its value is replaced as specified",
				},
				"api-resources": {
					Command: "kubectl api-resources",
					Args: []ArgumentDescriptor{
						{
							Argument: "--namespaced=true",
							Description: "all namespaced resources",
						},
						{
							Argument: "--namespaced=false",
							Description: "all non-namespaced resources",
						},
						{
							Argument: "-o name",
							Description: "all resources with simple output (just the resource name)",
						},
						{
							Argument: "-o wide",
							Description: `all resources with expanded (aka "wide") output`,
						},
						{
							Argument: "--verbs=list,get",
							Description: `all resources that support the "list" and "get" request verbs`,
						},
						{
							Argument: "--api-group=extensions",
							Description: `all resources in the "extensions" API group`,
						},
					},
					Description: "list all supported resource types along with their shortnames, API group, whether they are namespaced, and Kind",
				},
			},
		},
	},
}

