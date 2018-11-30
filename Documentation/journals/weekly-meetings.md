# **November 17, 2018**

We have decided to implement a Grocery store application. The following Go APIs will be implemented:

- Users - Add and search for users

- Inventory - Get and update inventory

- Cart - Add, remove, update cart contents

- Payment - Process payments

- Reviews - Add reviews 



Responsibilities:

- Clifton Francis
  - Implement Payment API
  - DB - MongoDB cluster
  - Kubernetes Horizontal Pod Autoscaling research
- Pavan Kumar Shekar
  -Implement Riak Cluster
  -Implement CART API
  -Intigrate with other Microservices (Inventory and Payment)
- Ujjval Soni
  -Implement the Reviews API (which also includes the maintaining sessions)
  - DB - MongoDB cluster (not sure, might change to Riak cluster in the future based on the needs of the project)
  -Prepare a frontend for the to add and view reviews, will be using the inventory frontend as a base and built upon the same.
  -Updating the on the Kanban board for the tasks.
- Prajwal Venkatesh
- Shivam Waghela
  - Implement Users API
  - DB - MongoDB cluster

# **November 29, 2018**

**Location:** Google Hangouts</br>
**Time:** 7:30PM - 8:30PM</br>

**Attendees:**

1. Shivam Waghela
2. Pavan Kumar Shekar
3. Ujjval Soni
4. Clifton Francis
5. Prajwal Venkatesh


## Key Discussion Points
1. API Design
2. Frontend Design
2. Data types and Data format used for API communication 


## Decisions

1. All the team members must complete their personal projects by Dec 2.
2. Everyone should finish the Ping API by Dec 2.
3. Start work on Frontend.
4. Schedule personal meetings with people whose APIs are dependent on each other.
5. After completing the API on instance move on to deploy as docker containers.


## Action Items
1. Design the Frontend with React.
2. Design the API and deploy on ec2 instances.
3. Connect and test your API with the database cluster. 
4. Dockerize the API after successful testing on ec2.
