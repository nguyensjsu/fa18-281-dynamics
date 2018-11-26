## Kubernetes

### Amazon EKS

- DOC: https://aws.amazon.com/eks/

Available Regions:

- US East (N. Virginia)
- US East (Ohio)
- EU (Ireland)
- US West (Oregon)

Cost: 

EKS Control Plane       $0.20/hr
Worker Nodes            Normal EC2 Pricing



### Creating a cluster using Kubeadm (Ubuntu)

DOC: 
- https://kubernetes.io/docs/setup/independent/create-cluster-kubeadm/
- https://github.com/kubernetes-incubator/metrics-server

#### Installing Master Node

Install Docker
```
$ sudo apt-get update
$ sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
$ sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
$ sudo apt-get update

Installing latest version of docker supported by Kubernetes
$ export VERSION=18.06 && curl -sSL get.docker.com | sh
```

Install Kubernetes
```
$ curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
$ sudo echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" >> /etc/apt/sources.list.d/kubernetes.list
$ sudo apt-get update
$ sudo apt-get install -y kubelet kubeadm kubernetes-cni
```

Swap must be disabled
```
$ sudo swapoff -a
```

Start Kubernetes Cluster
```
$ sudo kubeadm init --apiserver-advertise-address=<HOST_IP_ADDRESS> --service-cidr=10.96.0.0/12 --pod-network-cidr=10.244.0.0/16 --feature-gates=CoreDNS=false
```
**NOTE**: Store the join command printed. Each of the worker nodes will need to run that command in order to join the cluster.

Allowing normal user to run kubectl commands
```
$ mkdir -p $HOME/.kube
$ sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/admin.conf
$ sudo chown $(id -u):$(id -g) $HOME/.kube/admin.conf
$ export KUBECONFIG=$HOME/admin.conf
$ echo "export KUBECONFIG=$HOME/admin.conf" | tee -a ~/.bashrc
```

Deploying Flannel Networking Layer
```
$ kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/bc79dd1505b0c8681ece4de4c0d86c5cd2643275/Documentation/kube-flannel.yml
```

#### Adding a worker node (To join the cluster)

Install Docker
```
$ sudo apt-get update
$ sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
$ sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
$ sudo apt-get update

Installing latest version of docker supported by Kubernetes
$ export VERSION=18.06 && curl -sSL get.docker.com | sh
```

Install Kubernetes
```
$ curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
$ sudo echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" >> /etc/apt/sources.list.d/kubernetes.list
$ sudo apt-get update
$ sudo apt-get install -y kubelet kubeadm kubectl
$ sudo apt-mark hold kubelet kubeadm kubectl
```

Turn off swap
```
$ sudo swapoff -a
```

Join the cluster
```
$ JOIN_COMMAND (FROM_MASTER_KUBEADM_INIT)
```

### Horizontal Pod Autoscaling

DOC: 
- https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/
- https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/

Metrics Server is essential for keeping track of the pods resource utilization. In the master node:
```
$ git clone https://github.com/kubernetes-incubator/metrics-server
$ kubectl create -f metrics-server/deploy/1.8+/
```

Walkthrough
```
$ kubectl run php-apache --image=k8s.gcr.io/hpa-example --requests=cpu=200m --expose --port=80
$ kubectl autoscale deployment php-apache --cpu-percent=50 --min=1 --max=10
$ kubectl run -i --tty load-generator --image=busybox /bin/sh

In the container:
/# while true; do wget -q -O- http://php-apache.default.svc.cluster.local; done

In another terminal:
$ watch kubectl get hpa

```
You should see it scale up to 10 pods within a minute or two. Once you are satisfied that it works, you can scale down by stopping the load. In the container: 
```
/# <Ctrl> + C
```
