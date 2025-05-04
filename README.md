# 🌍 hello-world-k8s

Go server that displays a **"Hello World!"** message, deployed to **Kubernetes**. Serves as a great learning point for setting up Kubernetes for a project. Includes a **Deployment** and a **Service** to expose the app's pod(s).

---

## 🚀 Run the app (Minikube, macOS):

### 1️⃣ Clone the repository:
```bash
git clone https://github.com/Robert076/hello-world-k8s
```

---

### 2️⃣ Run Kubernetes setup:
```bash
kubectl apply -f kubernetes/
```

---

## 🌐 Access the App

### ✅ Option 1: Port forward the service
```bash
kubectl port-forward service/hello-world-service 8080:8080
```

Visit 👉 [http://localhost:8080](http://localhost:8080)

---

### 🔁 Option 2: Use Minikube service (a bit longer)

1. Get the name of the service:
```bash
kubectl get svc
```

2. Get the URL for the service:
```bash
minikube service <the_name_from_previous_step> --url
```

Visit the provided URL in your browser.

---

💡 **Tip:** This setup is perfect for practicing Kubernetes basics like deployments, services, and local development with Minikube.
