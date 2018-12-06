package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"github.com/unrolled/render"
	"io/ioutil"
	//"log"
	"math"
	"net/http"
	"strings"
	"time"
	"github.com/rs/cors"
)

//var server1 = "http://internal-RIAKelb-829969978.us-west-1.elb.amazonaws.com:8098" // set in environment
//var server2 = "http://internal-RIAKelb-829969978.us-west-1.elb.amazonaws.com:8098"  // set in environment
//var server3 = "http://internal-RIAKelb-829969978.us-west-1.elb.amazonaws.com:8098"  // set in environment
//var server4 = "http://internal-RIAKelb-829969978.us-west-1.elb.amazonaws.com:8098"
//var server5 = "http://internal-RIAKelb-829969978.us-west-1.elb.amazonaws.com:8098"
//var elboregon = "http://internal-RIAKelb-829969978.us-west-1.elb.amazonaws.com:8098"
var elbcart = "http://internal-RIAKelb-829969978.us-west-1.elb.amazonaws.com:8098"

//var elb = "http://riak-cali-131891034.us-west-1.elb.amazonaws.com:80"
var debug = true

var tr = &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}

// Create a new client
func NewClient(server string) *Client {
	return &Client{
		Endpoint: server,
		Client:   &http.Client{Transport: tr},
	}
}

// // Create a new server
// func NewServer() *negroni.Negroni {
// 	formatter := render.New(render.Options{
// 		IndentJSON: true,
// 	})
// 	n := negroni.Classic()
// 	mx := mux.NewRouter()
// 	initRoutes(mx, formatter)
// 	n.UseHandler(mx)
// 	return n
// }

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
		corsObj := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
        AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
    })
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.Use(corsObj)
	n.UseHandler(mx)
	return n
}

// Ping the API to check if its working.
func (c *Client) Ping() (string, error) {
	resp, err := c.Get(c.Endpoint + "/ping")

	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return "Ping Error!", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if debug {
		fmt.Println("[RIAK DEBUG] GET: " + c.Endpoint + "/ping => " + string(body))
	}
	return string(body), nil
}

// Create a new order
func (c *Client) CreateOrder(key, reqbody string) (Cart, error) {
	var ord_nil = Cart{}

	resp, err := c.Post(c.Endpoint+"/buckets/carttype/keys/"+key+"?returnbody=true",
		"application/json", strings.NewReader(reqbody))

	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return ord_nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if debug {
		fmt.Println("[RIAK DEBUG] POST: " + c.Endpoint + "/buckets/carttype/keys/" + key + "?returnbody=true => " + string(body))
	}

	var place Cart

	err = json.Unmarshal(body, &place)

	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return ord_nil, err
	}
	return place, err
}

// View order of specific key
func (c *Client) GetOrder(key string) Cart {
	var ord_nil = Cart{}
	resp, err := c.Get(c.Endpoint + "/buckets/carttype/keys/" + key)

	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return ord_nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if debug {
		fmt.Println("[RIAK DEBUG] GET: " + c.Endpoint + "/buckets/carttype/keys/" + key + " => " + string(body))
	}

	var ord = Cart{}

	if err := json.Unmarshal(body, &ord); err != nil {
		fmt.Println("RIAK DEBUG] JSON unmarshaling failed: %s", err)
		return ord_nil
	}
	return ord
}

// Get keys of all objects stored in database.
func (c *Client) GetKeys() ([]string, error) {

	var keys_nil []string
	resp, err := c.Get(c.Endpoint + "/buckets/carttype/keys?keys=true")

	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return keys_nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if debug {
		fmt.Println("[RIAK DEBUG] GET: " + c.Endpoint + "/buckets/carttype/keys/ " + string(body))
	}

	var all_keys Keys
	err = json.Unmarshal(body, &all_keys)
	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return keys_nil, err
	}

	fmt.Println(all_keys)

	return all_keys.Keys, err
}

// Update order for updating completing order.
func (c *Client) UpdateOrder(cartEdit Cart) (Cart, error) {
	var ord_nil = Cart{}
	reqbody, _ := json.Marshal(cartEdit)

	// fmt.Println("Id is: ", cartEdit.Id)

	req_body := string(reqbody)

	req, _ := http.NewRequest("PUT", c.Endpoint+"/buckets/carttype/keys/"+cartEdit.Id+"?returnbody=true", strings.NewReader(req_body))
	req.Header.Add("Content-Type", "application/json")
	fmt.Println(req)
	resp, err := c.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if debug {
		fmt.Println("[RIAK DEBUG] GET: " + c.Endpoint + "/buckets/carttype/keys/" + cartEdit.Id + " => " + string(body))
	}

	var update Cart

	err = json.Unmarshal(body, &update)
	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return ord_nil, err
	}
	return update, err
}

// Clear the cart of current order session.
func (c *Client) ClearCart(reqbody string) error {
	req, err := http.NewRequest("DELETE", c.Endpoint+"/buckets/carttype/keys/"+reqbody, nil)
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return err
	}

	_, err = c.Do(req)
	if err != nil {
		fmt.Println("[RIAK DEBUG] " + err.Error())
		return err
	}

	return nil
}

// Initialize our server and test ping.
func init() {
	// Riak KV Setup

	

	c6 := NewClient(elbcart)
	msg1, err := c6.Ping()
	if err != nil {
		fmt.Println("[INIT DEBUG] " + err.Error())
	} else {
		fmt.Println("ELB ping working ", msg1)
	}

}

// Initializing routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/cart/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/cart/add", newOrderHandler(formatter)).Methods("POST")
	mx.HandleFunc("/cart/view/{id}", getOrderHandler(formatter)).Methods("GET")
	//mx.HandleFunc("/history/{id}", viewCartHandler(formatter)).Methods("GET")
	mx.HandleFunc("/cart/edit", updateCartHandler(formatter)).Methods("PUT")
	mx.HandleFunc("/cart/clear", clearCartHandler(formatter)).Methods("DELETE")
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Println("[FAIL ON ERROR DEBUG] %s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

// Handles the ping call
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		c1 := NewClient(elbcart)

		message, err := c1.Ping()

		if message == "OK" {
			message = "Cart API live!!"
		}

		if err != nil {
			fmt.Println("[HANDLER DEBUG] ", err.Error())
			return
		} else {
			formatter.JSON(w, http.StatusOK, message)
		}
	}
}

// Handle new order request
func newOrderHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var newCart Cart
		uuid, _ := uuid.NewV4()

		decoder := json.NewDecoder(req.Body)
		// fmt.Println(decoder)

		err := decoder.Decode(&newCart)
		if err != nil {
			ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			fmt.Println("[HANDLER DEBUG] ", err.Error())
			return
		}

		newCart.Id = uuid.String()
		cartItems := newCart.Items

		var totalAmount float64

		for i := 0; i < len(cartItems); i++ {
			cartItems[i].ItemSubtotal = calculateAmount(cartItems[i].ItemQuantity, cartItems[i].ItemRate)
			totalAmount += cartItems[i].ItemSubtotal
		}

		totalAmount = math.Ceil(totalAmount*100) / 100
		newCart.CartTotal = totalAmount

		reqbody, _ := json.Marshal(newCart)

		c := NewClient(elbcart)
		val_resp, err := c.CreateOrder(uuid.String(), string(reqbody))

		if err != nil {
			fmt.Println("[HANDLER DEBUG] ", err.Error())
			formatter.JSON(w, http.StatusBadRequest, err)
		} else {
			formatter.JSON(w, http.StatusOK, val_resp)
		}
	}
}

// To view our order pass order id
func getOrderHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		params := mux.Vars(req)
		var uuid string = params["id"]
		// fmt.Println( "Order Params ID: ", uuid )

		if uuid == "" {
			formatter.JSON(w, http.StatusBadRequest, "Invalid Request. Order ID Missing.")
		} else {

			c := NewClient(elbcart)

			ord := c.GetOrder(uuid)

			if ord.Id == "" {
				formatter.JSON(w, http.StatusBadRequest, "")
			} else {
				fmt.Println("Order Details: ", ord)
				formatter.JSON(w, http.StatusOK, ord)
			}
		}
	}
}

//to view  cart with userid, pass user id
func viewCartHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("View Cart handler called.")

		params := mux.Vars(req)
		var uid string = params["id"]

		if uid == "" {
			formatter.JSON(w, http.StatusBadRequest, "User name missing")
		} else {
			c := NewClient(elbcart)

			cart_keys, err := c.GetKeys()
			cart_list := []Cart{}
			for _, item := range cart_keys {
				if c.GetOrder(item).UserName == uid {
					cart_list = append(cart_list, c.GetOrder(item))
				}
			}

			if err != nil {
				fmt.Println("[HANDLER DEBUG] ", err.Error())
				formatter.JSON(w, http.StatusBadRequest, err)
			} else {
				formatter.JSON(w, http.StatusOK, cart_list)
			}
		}

	}
}

func clearCartHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Clear cart Handler called.")

		var orderId Order
		decoder := json.NewDecoder(req.Body)
		// fmt.Println(decoder)

		err := decoder.Decode(&orderId)

		if err != nil {
			ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			fmt.Println("[HANDLER DEBUG] ", err.Error())
			return
		}

		reqbody, _ := json.Marshal(orderId)

		c := NewClient(elbcart)
		err = c.ClearCart(string(reqbody))

		if err != nil {
			fmt.Println("[HANDLER DEBUG] ", err.Error())
			formatter.JSON(w, http.StatusBadRequest, err)
		} else {
			formatter.JSON(w, http.StatusOK, "Cart Items deleted")
		}

	}
}

func updateCartHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var newCart Cart
		decoder := json.NewDecoder(req.Body)
		// fmt.Println(decoder)

		err := decoder.Decode(&newCart)

		if err != nil {
			ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			fmt.Println("[HANDLER DEBUG] ", err.Error())
			return
		}

		var totalAmount float64

		cartItems := newCart.Items
		for i := 0; i < len(cartItems); i++ {
			cartItems[i].ItemSubtotal = calculateAmount(cartItems[i].ItemQuantity, cartItems[i].ItemRate)
			totalAmount += cartItems[i].ItemSubtotal
		}

		totalAmount = math.Ceil(totalAmount*100) / 100

		newCart.CartTotal = totalAmount

		c := NewClient(elbcart)
		val_resp, err := c.UpdateOrder(newCart)

		if err != nil {
			fmt.Println("[HANDLER DEBUG] ", err.Error())
			formatter.JSON(w, http.StatusBadRequest, err)
		} else {
			formatter.JSON(w, http.StatusOK, val_resp)
		}
	}

}

func calculateAmount(count int, ItemRate float64) float64 {
	total := float64(count) * ItemRate
	total = math.Ceil(total*100) / 100
	return total
}
