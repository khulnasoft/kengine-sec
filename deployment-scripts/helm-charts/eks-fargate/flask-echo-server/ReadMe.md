# EKS Fargate: Sample helm chart with Khulnasoft agent

### Install

Application source: https://github.com/khulnasoft/Kengine/tree/main/deployment-scripts/sample-applications/flask-echo-server/

**Quick start**

```bash
helm install flask-echo-server flask-echo-server \
    --set managementConsoleIp=40.40.40.40
```

**Detailed setup**

- Create values file
```bash
helm show values flask-echo-server > flask_echo_server_values.yaml
```
- Set Khulnasoft management console ip address or domain name
```yaml
managementConsoleUrl: ""
managementConsolePort: "443"
```
- (Optional) Change image
```yaml
khulnasoftAgentImage:
  name: khulnasoft/khulnasoft_agent_ce
  tag: latest-fargate
```
- Set khulnasoft auth key
Set authentication key when it is enabled in management console
```yaml
# Auth: Get khulnasoft api key from UI -> Settings -> User Management
khulnasoftKey: ""
```
- (Optional) Instance id suffix
Custom Amazon Machine Images might have same hostnames for multiple instances. This can be used to distinguish vm's. 
```yaml
# Suffix cloud instance id to hostnames
instanceIdSuffix: "N"
```
- Install khulnasoft-agent helm chart with values file
```bash
helm install -f flask_echo_server_values.yaml flask-echo-server flask-echo-server
```

### Delete

```bash
helm delete flask-echo-server
```