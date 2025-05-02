# hello-world-k8s
Go server that displays a "Hello World!" message, deployed to Kubernetes. Serves as a good learning point for figuring out how to setup Kubernetes for a project. Contains a deployment, a service to expose the deployment's pod(s).

Run the app (minikube, macOS):

1) Clone the repository:
   
```bash
git clone https://github.com/Robert076/hello-world-k8s
```

2) Run Kubernetes setup

```bash
kubectl apply -f kubernetes/
```

3) Get the name for the service

```bash
kubectl get svc
```

4) Get the URL for the service

```bash
minikube service <the_name_from_previous_step> --url
```

5. Visit localhost:8080/.
