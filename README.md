## README.md

---

# Multiply By 10 - NUMAflow Go SDK Project

This project utilizes the NUMAflow Go SDK to implement a User Defined Function (UDF) that multiplies input values by 10. The solution is deployed on a Kubernetes cluster, and for local development/testing, port forwarding is required.

---

### Prerequisites

1. [Go](https://golang.org/dl/) installed and configured.
2. Access to a Kubernetes cluster with `kubectl` command-line tool installed.
3. [NUMAflow Go SDK](https://numaflow.numaproj.io/user-guide/user-defined-functions/map/map/) installed (replace with actual SDK link).

---

### Getting Started

#### 1. Clone the Repository

```bash
git clone https://github.com/shubhamdixit863/numaflowmapudf.git
cd multiply-by-10-numaflow
```

#### 2. Build and Deploy

Ensure your `kubectl` context is set to the correct Kubernetes cluster. Then run the provided deployment script:

```bash
./deploy.sh
```

#### 3. Set Up Port Forwarding

To forward traffic from your local machine to the `multiply-by-10` pod in the Kubernetes cluster, execute the following command:

```bash
kubectl port-forward multiply-by-10-in-0-efy9w 8444:8443
```

With this setup, the UDF can be accessed locally on `http://localhost:8444/`.

---

### Testing the UDF

With the port-forwarding set up, you can test the UDF by making a request to the service:

```bash
curl -X POST http://localhost:8444/multiply -d '{"value": 5}'
```

Expected Response:
```json
{
  "result": 50
}
```

---

### Troubleshooting

1. **Port Forwarding Issues:** Ensure that the pod name `multiply-by-10-in-0-efy9w` is correct. Pod names might change based on deployment configurations. Use `kubectl get pods` to list available pods.

2. **UDF Response Issues:** Check the pod logs for any runtime errors:
```bash
kubectl logs multiply-by-10-in-0-efy9w
```

---

### Contributing

Please read our [CONTRIBUTING.md](./CONTRIBUTING.md) file for details on our code of conduct, and the process for submitting pull requests to the project.

---

### License

This project is licensed under the MIT License - see the [LICENSE.md](./LICENSE.md) file for details.

---

### Acknowledgments

- The NUMAflow team for their robust SDK.
- All contributors to this project.

---

Replace placeholders such as repository URLs, links, and other specific details with the appropriate values related to your project.
