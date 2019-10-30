# Assignment used for Data Engineer position
This assignment ... bla.. bla...

## The problem
Image pre-processing is an important task in many machine learning applications, but can be a heavy task...

## The sample code
We have provided the code for a simple gRPC API for the image service, along with YAML manifests for deploying the API to a Kubernetes cluster. The code currently just echos the image sent in the request, no pre-processing done yet

### Local testing
The sample is written in Go, with the API defined by the Protocol Buffer found in `proto`. You can start the API locally by running:

```go run cmd/service/service.go```

And test it out with the client:

```go run cmd/client/client.go test.png```

### Deploy on Kubernetes
To deploy the service use the manifests in `manifests`

## Your Task
We would like you to extend this system with actual image pre-processing. 

### Part 1 - The mandatory part
Create you own micro-service to do the image processing, seperate from the API service. You can choose to write it in whatever programming language you are comfortable with. We expect it to:

* Contain an API definition in a seperate protobuf file
* Handle requests through gRPC
* Take as part of the request:
..* Parameters for scaling the image
..* Parameters for Gray scaling the image or not

### Part 2 - the optional extras
