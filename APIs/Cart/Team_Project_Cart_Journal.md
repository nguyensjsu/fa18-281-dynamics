# PROJECT JOURNAL (WEEKLY UPDATES /PROGRESS)
## CMPE 281 Fall 2018 - Individual NOSQL project and Enhancement to Team project.

###STUDENT NAME: Pavan Kumar Shekar
###STUDENT ID: 012451907

-

Youtube demo for the project can be viewed [here](https://youtu.be/eECJQla6f9s).

## Architecture Diagram


### Week 1 (November 10, 2018 - November 16, 2018)

##### Agenda of the week:
1. Decision on the NoSQL database to use.
2. Decide how the database has to integrate with the API that we will be creating
3. Decide the end points for API.
4. Initial Riak setup

##### Completed:
1. Decided to use Riak and Mongo DB as databases.
2. Team discussion for overview of our application
3. Changed database to Riak. Which is Highly Available and Partition Tolerant.
4. The decision of choosing RIAK was made because CART API needs high availability over consistency.The CRUD operations of POST, PUT, GET and DELETE will be performed.
5. Riak Documentation was searched for.

##### Roadblocks or issues faced:
1. Deciding on which Database would suit ideally for my API. Considering so many options.
2. Gathering documentation for Riak and understanding the concept of Bucket-type and Buckets.
3. Minimal documentation available on the internet.


### Week 2 (November 17, 2018 - November 23, 2018)

##### Agenda of the week:
1. Launch 5 nodes of EC2 instances and create a replica set on Mongo DB and Riak as well.
2. Try CRUD operations with the cluster and check replication.

##### Completed:
1. Launced 5 EC2 instance creating a cluster on Mongo.
2. Tested if replication is working by adding a key value to it.

##### Roadblocks or issues faced:
1. Riak cluster creation was not easy.
2. Lack of examples and documentation.
___

### Week 3 (November 24, 2018 - November 30, 2018)

##### Agenda of the week:
1. Create Riak cluster.
2. Create sample Go API to hit my DB instance.

##### Completed:
1. Partition tolerance test on RIAK and Mongo DB
2. High Availability test of RIAK and Mongo DB


##### Roadblocks or issues faced:
1. Which method to use for partition: Netwrok AC/IP Tables/ Security Groups.
2. Use of ELB with multiple VPCs was not possible.
___

### Week 4 (December 1, 2018 - December 7, 2018)

##### Agenda of the week:
1. Develop API and connect it to RIAK.
2. Perform all crud operations using API calls to database.

##### Completed:
1. Developed a GoAPI for cart that can perform CRUD operations
2. Connected API with RIAK that could do POST, PUT, DELETE, GET calls.

##### Roadblocks or issues faced:
1. Configuring API to connect to RIAK.
2. Getting buckets from databse and then displaying data.
___

### NoSQL Database Project/ Team Project

The aim of the project is to create a cluster of databases ideal for the application and micro service chosen and make sure the replication of data is happeneing and test how the Database behaves during a partition. Also how well the cluster recovers after the removal of the partition.

### Overview of Project:
1. Research on the databases avaialable and choose the right database.
2. Setup the DB instances on different subnets on AWS platform
3. Develop a Go API which connects to the DB cluster to deploy on instance.

### Phase 1
1. What is the API I will be choosing?  
    A Cart API for a Grocery store.  

2. What could be the API Calls?  
    The API calls will include the CRUD operations:
    * Create Cart
    * Add to Cart
    * Update Cart
    * Delete Cart

3. What all could be there in cart?  
    Items in the Grocery store with their price and the number of items is to be present. Also, the total amount of the cart is required.  

### Sample collection/table might look like
Item_Name | Item_Count | Price | Cart_Amount
--- | --- | --- | --- |
Cereal Box| 2 | 7.2 | 14.4
Milk | 4 | 5.25 | 21

___

### About RIAK
* Riak is a NoSQL Database that is built to show High Availablity and Partition tolerance.  
* Riak has concept of Buckets just like Mongo has concept of Collection. Inside Buckets, you add objects.
* You create a Bucket and inside that you can store multiple objects. Each object is associated with a key, that makes it a Key-Value pair.
* Riak doesn't have Master-slave relationship like MongoDB or Redis but rather Coordinator and Members.
* Good thing about having Coordinator-Member relationship is that any node can respond to WRITE request unlike in Master-Slave case where a election is organised if Master goes down.
* In riak, members request to coordinator to join the cluster and then coordinator plans how much dataload they share and then creates the cluster.


### Partition Tolerance of RIAK
* Creation of a cluster with 5 nodes.
* Writes and Reads can be done on any of the Nodes
* In the case of blockage of communication between any of the nodes, say isolation of a few nodes, Reads and Writes can still occur on any node. Though it may contain stale data, it has to be always available. 


### High Availability of RIAK
In the case of a failure of a node or a group of nodes, the database should still be able to function as expected without causing issues to the user.


### Approaching AP with RIAK
* We plan to launch 5 instances of RIAK and make a cluster with them.
* To test High Availability we will bring down 1 node and then maybe increase the count of offline nodes and see if cluster still responds to user querry.
We will also bring down Coordinator node and see if it has any effect on the cluster.
* Partition tolerance will be tested through IP tables. Nodes are isolated from each other with IP table rules.

### Setting up AWS
* Riak KV is available in Amazon marketplace making it easier to setup.
* instances are split among different subnets which are placed in different availablitlity zones.
* These instances are accessed through a public instance which in our case is called a jumpbox instance.

##### Options thought of for creating partition
* Create partition using Network AC. Associate the subnets with different inbound and outbound rules
* Security groups are changed accordignly
* IP tables are used to isolate nodes.

___

## Steps to setup RIAK KV
* After launching Riak instance, make changes in riak.conf to mention range of ports for Erlang to use while communicating with other nodes of same cluster.
* Also, we have to make sure our ports are exposed for inter-node communication as well as from outer world.
* To start the service of Riak do: `sudo riak start`
* From all other instances that we want to join the cluster, we run following command: `sudo riak-admin cluster-join riak@<private ip of coordinator>`
* From coordinator plan the joining of members: `sudo riak-admin cluster plan`
* From coordinatoer commit the changes: `sudo riak-admin cluster commit`
* To view changes that we commited or statuses of member do: `sudo riak-admin member-status`
* To check if things are working, we upload an image and check if it works or not:  
`curl -XPUT http://<ip-of-instance>:8098/riak/images/test.jpg \
    -H "Content-type: image/jpeg" \
    --data-binary @<img-name>.jpg`  
* Above command will upload the image on mentioned url and we can check that in our browser.
* To see list of buckets do:  
`curl -i http://<ip-address>:8098/buckets?bucktes=true`  
* To see list of keys we do:
`curl -i http://<ip-address>:8098/buckets/keys?keys=true`  
* To remove a node from cluster do this in the node that wants to leave:
`sudo riak-admin leave <node-name>`  
* To forcefully make a node leave, do this from coordinator:
`sudo riak-admin force-remove -f <node-name>`

___

## Testing our work
* Testing is done using adding a key value through a node.
* reading this value from another node in the cluster.
* Updating this through a different node
* Reading this again.
* creating a partition.
* Update the value.
* Read from multiple instances.
* Check stale data
* Remove partition
* Check partition recovery

## 2 Ways to figure out replication in RIAK
* Last write wins. The last write done to a node in the cluster is replicated through out.
* Updates are made through vector clocks and this avoids confusion.



