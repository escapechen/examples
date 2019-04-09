# Scope
## Task
* Setup and run Kubernetes Pod and use it to provide info about running Ingress Pods running in same cluster via Kubernetes API
* expose found info via website
## Ingress GO Piece (main.go)
* the ingress example is based on [KubernetesAPI](https://github.com/kubernetes/client-go/tree/master/examples/in-cluster-client-configuration)
* with some research and trial&error i found clientset.ExtensionsV1beta1().Ingresses("").List
* i simply lack the experience (coming from php arrays) to deal with this nasty result struct. Need feedback on how to pull out useful maps out of this garbled thing. Spew to the rescue to get at least something to show for :-)
* found this fine piece of code to get a webserver going and write some output. Needs wrapping into proper HTML syntax.
* This definetely needs improvement, but hey this is after 12h shared kubernetes/docker/go bootcamp. 
## Docker Piece (Dockerfile)
* checkout this ingress example
* use Dockerfile to build a docker image
`docker build -t ingress:final .`
* find new image and note Image ID
`docker images`
* tag new image for upload and push to private repo, new TAG avoids cache
`docker tag <ID> escapechen/examples:ingress_NR
docker push escapechen/examples:ingress_NR`
## Kubernetes Piece (kubernetes/)
* the _permission*_ Files deal with creating a dedicated service user which has only a very limited view on the cluster (basicall just view ingresses)
* _listService_ File introduces this extra layer between ingresses and application, to allow them to die/jump between nodes etc. Assuming there is some iptables NAT going on in the back. In our case, it points to our app on port 8000 and listens on port 80 to receive data from the ingress nginx.
* _ingress_ is the original ingress file for the initial LAB setup. I extended it to map /ingress to our new service. This way we also have a little bit more input for our Ingress Go Piece.
* _deployment-ng_ File is abstraction to pods. Basically sets up the environment, starts the pod(s) with various images, allows for replicas etc. In our case, it sets up a single pod using our dedicated user, pulls the image from the private github repo using provided credentials and fires up the _ingress GO piece_ which will continously listen on port 8000.
* _missingPiece_ is the secret for the docker registry, which can be created via:
`kubectl create secret docker-registry marcelk-regcred --docker-server=https://index.docker.io/v1/ --docker-username=escapechen --docker-password=NOBODY-KNOWS --docker-email=fancy@email.com`
