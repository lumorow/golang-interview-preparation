# Docker
Docker is a tool for building packages, delivering and running applications in "containers"

## What is a Docker container?

A container is a basic unit of software that encloses code and all its dependencies
to ensure that the application runs transparently, quickly and reliably, regardless of the environment. Container
Docker can be built using a Docker image. This is an executable software package,
containing everything necessary to run the application, for example, system programs, libraries, code, runtime environments
and settings.

## Components of the Docker architecture

The main components of the Docker architecture are:

- `server`, contains the Docker service, images and containers. The service contacts `Registry`,
  images - metadata for applications running in Docker containers.
- `client`, used to run various actions on the Docker server.
- `registry`, used to store images. There are public ones available to everyone, for example, `Docker Hub` and `Docker Cloud`.

## The most important Docker commands

The most important Docker commands are:
- `build`, building an image for Docker
- `create`, creating a new container
- `kill`. forced stop of the container
- `dockerd`, start the Docker service
- `commit`, creating a new image from changes in the container
## Docker container state

To determine the status, you need to run the command:

> docker ps -a

This command will list all available containers with their status on the server.
From this list you need to select the required container and find out its status.

## Docker image, Docker run

A Docker image is a set of files combined with settings that can be used to create instances of
which run in separate containers as isolated processes. The image is built using
instructions for obtaining the full executable version of the application, depending on the version of the server kernel.
The `Docker run` command is used to create such instances, called containers, run by
using Docker images. When running one image, the user can create multiple containers.

## Differences between virtualization and containerization

`Virtualization` allows you to run multiple operating systems on one physical server.
`Containerization` runs on the same operating system in which the applications are packaged
into containers and run on one server/virtual machine.

## Advanced Docker commands

The most important of them:

- `docker -version`: find out the installed Docker version;
- `docker ps`: list all running containers along with additional information about them;
- `docker ps -a`: list all containers, including stopped ones, along with additional information about them;
- `docker exec`: enter the container and execute a command in it;
- `docker build`: build an image from a Dockerfile;
- `docker rm`: delete the container with the specified identifier;
- `docker rmi`: delete the image with the specified identifier;
- `docker info`: get extended information about the installed Docker, for example, how many containers, images are running, kernel version, available RAM, etc.;
- `docker cp`: save the file from the container to the local system;
- `docker history`: show the history of the image with the specified name.

## Docker and Docker Compose

Docker is used to manage the individual containers (services) that make up an application.

Docker Compose is used to simultaneously manage multiple containers that make up an application.
This tool offers the same capabilities as Docker, but allows you to work with more complex applications.

## Resources

- [50 questions about Docker](https://habr.com/ru/companies/slurm/articles/528206/)
- [20 common Docker interview questions](https://itsecforu.ru/2020/01/15/%F0%9F%90%B3-20-%D1%80%D0%B0%D1%81%D0%BF%D1%80%D0%BE%D1%81%D1%82%D1%80%D0%B0%D0%BD%D0%B5%D0%BD%D0%BD%D1%8B%D1%85-%D0%B2%D0%BE%D0%BF%D1%80%D0%BE%D1%81%D0%BE%D0%B2-%D0%BF%D1%80%D0%BE-docker-%D0%BD%D0%B0/)
- [Docker Guide](https://medium.com/nuances-of-programming/%D1%80%D1%83%D0%BA%D0%BE%D0%B2%D0%BE%D0%B4%D1%81%D1%82%D0%B2%D0%BE-%D0%BF%D0%BE-docker-%D1%87%D0%B0%D1%81%D1%82%D1%8C-1-%D0%BE%D0%B1%D1%80%D0%B0%D0%B7-%D0%BA%D0%BE%D0%BD%D1%82%D0%B5%D0%B9%D0%BD%D0%B5%D1%80-%D1%81%D0%BE%D0%BF%D0%BE%D1%81%D1%82%D0%B0%D0%B2%D0%BB%D0%B5%D0%BD%D0%B8%D0%B5-%D0%BF%D0%BE%D1%80%D1%82%D0%BE%D0%B2-%D0%B8-%D0%BE%D1%81%D0%BD%D0%BE%D0%B2%D0%BD%D1%8B%D0%B5-%D0%BA%D0%BE%D0%BC%D0%B0%D0%BD%D0%B4%D1%8B-4997f2968925#:~:text=%D0%BD%D0%B0%D1%81%D1%82%D1%80%D0%BE%D0%B5%D0%BD%D0%BD%D0%BE%D0%B5%20%D1%81%D0%BE%D0%BF%D0%BE%D1%81%D1%82%D0%B0%D0%B2%D0%BB%D0%B5%D0%BD%D0%B8%D0%B5%20%D0%BF%D0%BE%D1%80%D1%82%D0%BE%D0%B2-,%D0%92%D1%8B%D0%B2%D0%BE%D0%B4%D1%8B,%D0%9F%D1%80%D0%B8%D0%BB%D0%BE%D0%B6%D0%B5%D0%BD%D0%B8%D0%B5%20%D1%80%D0%B0%D0%B1%D0%BE%D1%82%D0%B0%D0%B5%D1%82%20%D0%B2%D0%BD%D1%83%D1%82%D1%80%D0%B8%20%D0%BA%D0%BE%D0%BD%D1%82%D0%B5%D0%B9%D0%BD%D0%B5%D1%80%D0%B0%20Docker!)

## README.md

- eng [English](https://github.com/lumorow/golang-interview-preparation/blob/main/Docker/README.md)
- ru [Русский](https://github.com/lumorow/golang-interview-preparation/blob/main/Docker/readme/README.ru.md)