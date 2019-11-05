# Assignment used for Software Engineer position
This task is intended for candidates applying for a Software Engineer position at the Visma Machine Learning team. The assignment is built around the technologies and stack used in the production environments of the team, and the problem is a toy version of some of the tasks we face.


![Interview](./interview-gopher.png)


## The problem
Image pre-processing is an important task in many machine learning applications, but can be an intensive task when trying to keep request times low and throughput high. In the VML team we use a micro-service architecture running on Kubernetes, along with autoscaling implementations to acheive image processing systems that scale well.

For this task we want to implement a small service that scales and grayscales incoming images to some normalized settings, so they could be used as input for a machine learning system.

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
Firstly you need to build the Docker image and push it to some image registry that your cluster has access to. To deploy the service on a Kubernetes cluster use the manifests in `manifests`. The service and deployment in `api.yaml` should spin up the api and get you an external ip to connect to.

## Your Task
We would like you to extend this system with actual image pre-processing, and make the system more "production ready".

### Part 1 - The mandatory part
Create you own micro-service to do the image processing, seperate from the API service. You can choose to write it in whatever programming language you are comfortable with. We expect it to:

* Contain an API definition in a seperate protobuf file
* Handle requests through gRPC
* Take as part of the request:
    * The image originally sent to the API
    * Parameters for scaling the image or not
    * Parameters for grayscaling the image or not

Alongside your code you should also write a manifest YAML to setup a deployment and service. The deployment should specify that your container should be replicated at least in 2 pods, which should be load-balanced through the service.

You will also have to add code in the existing API code in `pkg/api/api.go`, so that it will forward the image to your service. You should not have to change the existing API and protobuf. The API should satisfy the following:

* The API should be configurable, such that the desired image size and grayscale options should be supplied at startup. How is up to you. You can use 1024x768 as the target size.
* The API should check the size of the incoming image, and decide what scaling options should be forwarded to your processing service

__A quick question:__
*What considerations should you have when loadbalancing gRPC requests?*

### Part 2 - the optional extras
Beyond the mandatory part we would like you to extend your service in some way. Choose one or more from below:

* Add an option to supply the API with a URL for downloading images (This is already present in the protobuf, but not implemented) - Whether you want to download the image content in the API and send bytes to your service - or send the URL and download in the service is up to you.
* Add an option for storing the resulting image in cload storage, like S3/GCS
* Add stats for number of requests per second, request times etc., to benchmark the performance of your system (Do what you find practical. The API has the basic scaffolding for exporting Prometheus metrics, but you would need to deploy a setup for that)
* Add autoscaling to your service
* More features you think could be relevant...
