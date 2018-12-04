Cart API 

To add to cart data to be passed:

{  
	"userId":"pavan001",
	   "items":[
      {"name":"Username", "count":4, "rate":3.1222}
   ]
}


The reponse is:

{
    "id": "4dc4d509-87b4-4f6c-b46e-aa08a216d4b0",
    "userId": "pavan001",
    "items": [
        {
            "name": "Username",
            "count": 4,
            "rate": 3.1222,
            "amount": 12.49
        }
    ],
    "total": 12.49
}

