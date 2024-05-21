# Helm chart for Khulnasoft Agent

### Install

**Quick start**

```bash
helm repo add khulnasoft https://khulnasoft-helm-charts.s3.amazonaws.com/kengine
```

```bash
helm install khulnasoft-agent khulnasoft/khulnasoft-agent \
    --set managementConsoleUrl=40.40.40.40 \
    --set khulnasoftKey="" \
    --set clusterName="prod-cluster" \
    --set global.imageTag="2.0.1" \
    --namespace khulnasoft \
    --create-namespace
```

**Detailed setup**

```bash
helm repo add khulnasoft https://khulnasoft-helm-charts.s3.amazonaws.com/kengine
```

- Create values file
```bash
helm show values khulnasoft/khulnasoft-agent > khulnasoft_agent_values.yaml
```
- (Optional) Edit values file and set docker hub username and password (if using your own registry)
```yaml
imagePullSecret:
  # Specifies whether a image pull secret should be created
  create: true
  registry: "docker.io"
  # registry: "https://index.docker.io/v1/"
  username: ""
  password: ""
  # The name of the imagePullSecret to use.
  # If not set and create is true, a name is generated using the fullname template
  name: ""
```
- Set Khulnasoft management console ip address
```yaml
managementConsoleUrl: ""
```
- Set image tag
```yaml
global:
  # this image tag is used every where for agents
  # to override set tag at agents level
  imageTag: 2.0.1
```
- Set khulnasoft auth key
Set authentication key when it is enabled in management console
```yaml
# Auth (Optional): Get khulnasoft api key from UI -> Settings -> User Management
khulnasoftKey: ""
```
- (Optional) Instance id suffix
Custom Amazon Machine Images might have same hostnames for multiple instances. This can be used to distinguish vm's. 
```yaml
# Suffix cloud instance id to hostnames
instanceIdSuffix: "N"
```
- Set kubernetes cluster name
```yaml
# Set custom name for the cluster and hostname prefix for agent vm's to easily identify in Khulnasoft UI.
# Example: prod-cluster or dev1-cluster
# It will be suffixed with hostname - prod-cluster-aks-agentpool-123456-vmss000001
clusterName: ""
```
- Set container runtime socket path
  By default, docker is disabled and containerd is enabled
```yaml
# Mount container runtime socket path to agent pod. Agent will detect which runtime it is using these files.
mountContainerRuntimeSocket:
  dockerSock: false
  # Change if socket path is not the following
  dockerSockPath: "/var/run/docker.sock"
  containerdSock: true
  # Change if socket path is not the following
  containerdSockPath: "/run/containerd/containerd.sock"
  crioSock: false
  # Change if socket path is not the following
  crioSockPath: "/var/run/crio/crio.sock"
```
- Install khulnasoft-agent helm chart with values file
```bash
helm install -f khulnasoft_agent_values.yaml khulnasoft-agent khulnasoft/khulnasoft-agent \
    --namespace khulnasoft \
    --create-namespace
```
- Wait for pods to start up
```bash
kubectl get daemonset -n khulnasoft
kubectl get pods -n khulnasoft
```

### Delete

```bash
helm uninstall khulnasoft-agent -n khulnasoft
```