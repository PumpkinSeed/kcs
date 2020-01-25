# kubectl-cheatsheet

### Config

- `k config view`
- `k config get-contexts`
- `k config current-context`
- `k config use-context NAMEOFCONTEXT`

### Apply

- `k apply -f file.yaml`
- `k apply -f dir/`
- `k apply -f https://domain.com/file.yaml`
- `k create deployment nginx --image=nginx`

### Viewing, Finding Resources

- `k get services`
- `k get pods --all-namespaces`
- `k get pods -o wide`

