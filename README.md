# devops-sample-apps

This repository contains sample applications and deployment resources for DevOps demonstrations, including Golang and PHP apps, Dockerfiles, and Helm charts.

## golang

Application in runtime needs p12 file(filename: file.p12) next to application binary.

## php

Application to run on production needs env `APP_ENV=prod` and file `config` next to index.php, 
repository contains `config.prod` and `config.dev`, for production purposes `config.prod` needs to be renamed to `config`. 

## Manual Installation

For step-by-step instructions on how to manually install the applications from ECR using Helm, see [INSTALL.md](./INSTALL.md).
