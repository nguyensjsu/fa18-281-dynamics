import React, { Component } from "react";
import axios from "axios";
import "./orders.css";

class Orders extends Component {
  constructor() {
    super();
    this.state = {
      orders: []
    };
  }
  render() {
    let orders_list = this.state.orders.map(order => {
      return (
        <div style={{ margin: 10 }} className="card">
          <div className="card-header order-header">
            <span className="order-status">Order Processed</span>
            <span className="order-total">Total: ${order.cart_total}</span>
            <span className="order-number">ORDER # {order._id}</span>
          </div>
          <div className="card-body">
            <div>
              <ol>
                {console.log(order)}
                {order["items"].map(item => {
                  return <li>{item.item_name}</li>;
                })}
              </ol>
            </div>
          </div>
        </div>
      );
    });
    return (
      <React.Fragment>
        <div className="container">
          <div id="title">
            <span className="page-header">Your Orders</span>
          </div>
          <hr />
          {orders_list}
        </div>
      </React.Fragment>
    );
  }

  componentDidMount() {
    let PAYMENT_HOST_ELB =
      "Payments-EKS-2070687438.us-west-2.elb.amazonaws.com";
    let PORT = 3000;
    let username = sessionStorage.getItem("username");
    axios
      .get(`http://${PAYMENT_HOST_ELB}:${PORT}/payments/${username}`)
      .then(response => {
        console.log("Status Code GET Orders:", response);
        this.setState({ orders: response.data });
      })
      .catch(err => {
        console.log(err);
      });
  }
}

export default Orders;
