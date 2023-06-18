# Overview

A case study for a simple API backend using `Golang`'s `gin` framework, `postgresql` and `redis`.
It also includes a `docker-compose` file to run the project in a containerized environment.

## Pre-requisites

The setup instructions assume that you have `debian` based linux distro with `git`, `docker` and `docker compose` plugin installed.

###  Docker Installation

To install `docker` with the `compose` plugin, run the following commands:

```bash
apt-get update
apt-get -y install \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg

echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null

apt-get update
apt-get -y install docker-ce docker-ce-cli containerd.io docker-compose-plugin
```

## Running the project

To run the project with the default config, `git clone` the project locally, pull the docker images and run the containers.

```bash
git clone https://github.com/rjbasitali/at-case-study && cd at-case-study

docker compose pull
docker compose up -d
```

Config variables can be changed by editing the `.env` file.
