# Introduction 
Imagine that this is a business application that is supposed to be hosted in a Azure Kubernetes cluster.
This repo holds the application code and a build pipeline for it.

In a build pipeline it will build docker image of the app and push it to centralized repo (ignored in this scenario).

# Tasks
Clean up the solution and make it work.
You can introduce as many of the "good" practices as you see fit.
Make the pipeline production ready on enterprise level.

# Build and Test
To locally buil the app run `go run main.go ../sample_data/data.json` from the `app` folder. 
Then it can be accessed locally using http://localhost:3000/

To Build docker image run `docker build -t my-golang-app .` from `app` folder.
Then to access app from docker image `docker run -p 127.0.0.1:80:3001 my-golang-app`
And then test it http://localhost:80/