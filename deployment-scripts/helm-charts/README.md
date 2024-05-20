### Using kubernetes helm charts in this directory
1. create a namespace for example kengine
```sh 
kubectl create ns kengine
```
2. Install khulnasoft router helm chart 
```sh 
helm install router ./khulnasoft-router --namespace kengine
```
3. Install khulnasoft console helm chart 
```sh
helm install console ./khulnasoft-console --namespace kengine
```
4. wait for all the services to start, access to console ui depends on router service type
5. to customize the installation refer values.yaml of khulnasoft-console and khulnasoft-router