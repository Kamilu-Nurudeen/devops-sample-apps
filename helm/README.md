# Helm Charts

This directory contains Helm charts for deploying the Golang and PHP sample applications to Kubernetes.

## Structure
- `golang-app/`: Helm chart for the Go application.
- `php-app/`: Helm chart for the PHP application.

Each chart includes templates for Kubernetes resources such as Deployments, Services, and Ingress.

## Packaging and Publishing
To package a chart:
```sh
helm package ./golang-app
helm package ./php-app
```

To push a packaged chart to Amazon ECR as an OCI artifact:
```sh
helm push devops-sample-apps-golang-<version>.tgz oci://<aws_account_id>.dkr.ecr.<region>.amazonaws.com/
helm push devops-sample-apps-php-<version>.tgz oci://<aws_account_id>.dkr.ecr.<region>.amazonaws.com/
```

## Deployment
After pulling the chart from ECR, you can install it with:
```sh
helm install my-golang-app oci://<aws_account_id>.dkr.ecr.<region>.amazonaws.com/devops-sample-apps-golang --version <version>
helm install my-php-app oci://<aws_account_id>.dkr.ecr.<region>.amazonaws.com/devops-sample-apps-php --version <version>
```
 
