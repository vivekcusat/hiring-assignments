# Assignment used for Data Engineer position
This task is intended for candidates applying for a Data Engineer position at the Visma Machine Learning team. The assignment is build around the technologies and stack used in the production environments of the team, and the problem a toy version of some of the tasks we face.

## The problem
Image pre-processing is an important task in many machine learning applications, but can be an intensive task when trying to keep request times low and throughput high. In the VML team we use a micro-service architecture running on Kubernetes, along with autoscaling implementations to acheive image processing systems that scale well.

For this task we want to implement a small service that scales and grayscales incomming images to some nomalized settings, so they could be used as input for a machine learning system.

We also want to be able to horizontally scale the service to account for changing amounts of requests throughout a day.

## The sample code
We have provided the code for a simple gRPC API for the image service, along with YAML manifests for deploying the API to a Kubernetes cluster. The code currently just echos the image sent in the request, no pre-processing done yet.

### Prerequisits
You need to have a working setup for writing code in Go, building container images with Docker, and deploying to a Kubernetes cluster with eg. `kubectl`

### Local testing
The sample is written in Go, with the API defined by the Protocol Buffer found in `proto`. You can start the API locally by running:

```go run cmd/service/service.go```

And test it out with the client:

```go run cmd/client/client.go```

This will take the image in `test.jpg`, send it to the service and and save the output to `out.jpg`

### Deploy on Kubernetes
To deploy the service on a Kubernetes cluster use the manifests in `manifests`. The service and deployment in `api.yaml` should spin up the api and get you an external ip to connect to.

## Your Task
We would like you to extend this system with actual image pre-processing, and make the system more "production ready". 

### Part 1 - The mandatory part
Create you own micro-service to do the image processing, seperate from the API service. You can choose to write it in whatever programming language you are comfortable with. We expect it to:

* Contain an API definition in a seperate protobuf file
* Handle requests through gRPC
* Take as part of the request:
..* Parameters for scaling the image or not
..* Parameters for Grayscaling the image or not

Alongside your code you should also write a manifest YAML to setup a deployment and service. The deployment should specify that your container should be replicated a least in 2 pods, which should be load-balanced through the service.

You will also have to add code in the existing API code in `pkg/api/api.go`, so that it will forward the image to your service. You should not have to change the existing API and protobuf. The API should:

* The API should be configurable, such that the desired image size and grayscale options should be supplied at startup. How is up to you
* The API should check the size of the incomming image, and decide what scaling options should be forwarded to your processing service

__A quick question:__
*What considerations should you have when loadbalancing gRPC requests?*

### Part 2 - the optional extras
Beyond the mandatory part we would like you to extend your service in some way. Choose one or more from below:

* Add an option to supply the API with a URL for downloading images (This is already present in the protobuf, but not implemented)
* Add an option for storing the resulting image in cload storage, like S3/GCS
* Add stats for number of requests, request times to benchmark the performance of your system
* Add autoscaling to your service
* More features you think could be relevant...
