
## Cart API

* cart is stored as nested struct containing userid, cartdid(id) and struct of items, containing Items.Name, Items.Count, Items.Rate and Items.Amount. It also stores the total of the cart in Total
* All the 5 IP addresses of the nodes are given in var. This will be changed to addresses of 2 different load balancers in 2 VPC's. 
* /ping used to check status of nodes
* /order : adds item to cart with userid
* /view/{id} pass the order id, and get the order details.
* /history/{id} pass the user id and get the status of cart 
* /update updates an element in cart 
* /clearcart clears the cart, clears riak (deprecated)

Not yet completly implemented. 

### Add item to cart /order 
Data to be passed into /order for adding items into cart is 
```JSON
{"userId":"user100",
	"items":[
		{
		"name":"starbucks",
		"count":1,
		"rate":3.95
		},
		{
		"name":"peets",
		"count":1,
		"rate":4.95
		}
		]
	
}
```
response is 
```JSON
{
    "id": "bd29ac5a-8801-4c0e-a102-852d7149ee5d",
    "userId": "user100",
    "items": [
        {
            "name": "starbucks",
            "count": 1,
            "rate": 3.95,
            "amount": 3.95
        },
        {
            "name": "peets",
            "count": 1,
            "rate": 4.95,
            "amount": 4.95
        }
    ],
    "total": 8.9
}
```
