# Manual Installation Guide

This guide explains how to manually install the Golang and PHP sample applications using Helm charts stored as OCI artifacts in Amazon ECR.

## Prerequisites
- AWS CLI installed and configured with access to ECR
- Helm v3.8.0+ installed
- Sufficient permissions to deploy to your Kubernetes cluster

## 1. Authenticate Helm to Amazon ECR

First, authenticate your Helm client to the ECR registry:

```sh
aws ecr get-login-password --region eu-north-1 | \
  helm registry login --username AWS --password-stdin <aws_account_id>.dkr.ecr.eu-north-1.amazonaws.com
```

Replace `<aws_account_id>` with your AWS account ID.

---

## 2. Install the Golang Application Chart

Install the Helm chart directly from ECR:

```sh
helm upgrade --install my-golang-app \
  oci://<aws_account_id>.dkr.ecr.eu-north-1.amazonaws.com/devops-sample-apps-golang \
  --version <version> \
  --namespace <namespace> \
  --create-namespace \
  -f <your-values-file>.yaml
```

- Replace `<version>` with the desired chart version (e.g., `0.1.0`).
- Replace `<namespace>` with your target namespace (e.g., `testing`).
- Replace `<your-values-file>.yaml` with your custom values file if needed.

---

## 3. Install the PHP Application Chart

```sh
helm upgrade --install my-php-app \
  oci://<aws_account_id>.dkr.ecr.eu-north-1.amazonaws.com/devops-sample-apps-php \
  --version <version> \
  --namespace <namespace> \
  --create-namespace \
  -f <your-values-file>.yaml
```

---

## 4. Verify the Installation

Check that the releases are deployed and pods are running:

```sh
helm list -n <namespace>
kubectl get pods -n <namespace>
```

---

## Notes
- Always use specific chart versions for reproducibility.
- Use values files to override image tags, resource limits, and environment variables as needed.
- For production, review and adjust security, resource, and ingress settings in your values files. 
