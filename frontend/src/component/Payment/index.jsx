import React, { Component } from "react";
import { Link } from "react-router-dom";
import axios from "axios";
import uniqid from "uniqid";

class Payment extends Component {
  constructor(props) {
    super(props);
    this.state = {
      cart: this.props.location.state.cart,
      cart_total: this.props.location.state.cart_total,
      item_count: this.props.location.state.item_count,
      wallet: 0
    };
  }

  pay = () => {
    let PAYMENT_HOST_ELB = "payments-1051217824.us-west-2.elb.amazonaws.com";
    let PORT = 3000;
    let data = {
      username: sessionStorage.getItem("username"),
      wallet_amount: this.state.cart_total
    };
    axios
      .put(`http://${PAYMENT_HOST_ELB}:${PORT}/wallet/pay`, data)
      .then(response => {
        console.log("Status Code POST Wallet:", response.status);
        console.log("response from POST Wallet:", response);
        this.setState({
          wallet: response.data.wallet_amount
        });
      });
    console.log("cart", this.state.cart);
    data = {
      username: sessionStorage.getItem("username"),
      items: this.state.cart,
      cart_total: this.state.cart_total
    };
    axios
      .post(`http://${PAYMENT_HOST_ELB}:${PORT}/payment`, data)
      .then(response => {
        console.log("Status Code POST Wallet:", response.status);
        console.log("response from POST Wallet:", response);
      });
  };

  render() {
    let item_list = this.state.cart.map(item => {
      return (
        <tr className="row" key={uniqid()}>
          <td className="item name col">
            <Link to="">{item.item_name}</Link>
          </td>
          <td className="item rate col">${item.item_rate}</td>
          <td className="item quantity col">{item.item_quantity}</td>
          <td className="item subtotal col">{item.item_subtotal}</td>
        </tr>
      );
    });
    return (
      <React.Fragment>
        <div className="container">
          <div id="title">
            <h2>Payment</h2>
            <hr />
          </div>
          <div>
            <h3>Your order</h3>
            <table className="table table-hover">
              <thead>
                <tr className="row">
                  <th className="item name-header col">Item</th>
                  <th className="item rate-header col">Rate</th>
                  <th className="item quantity-header col">Quantity</th>
                  <th className="item subtotal-header col">Subtotal</th>
                </tr>
              </thead>
              <tbody
                className="card-parent table table-hover container grocery-item"
                style={{ margin: "auto" }}
              >
                {item_list}
              </tbody>
            </table>
            <div className="total-row row">
              <span className="col-lg-8" />
              <div className="item-count col-lg-2">
                Total Items: {this.state.item_count}
              </div>
              <div className="item-total col-lg-2">
                Order total:{" "}
                <span className="item-total-value">
                  ${this.state.cart_total}
                </span>
              </div>
            </div>
            <div>
              <span style={{ fontSize: 25, fontWeight: 500 }}>
                Wallet Amount:{" "}
                <span style={{ fontWeight: 700 }}>${this.state.wallet}</span>
              </span>
              <span style={{ margin: 20 }}>
                <button onClick={this.pay} className="btn btn-primary">
                  Pay from Wallet
                </button>
              </span>
            </div>
          </div>
        </div>
      </React.Fragment>
    );
  }

  componentDidMount() {
    let PAYMENT_HOST_ELB = "payments-1051217824.us-west-2.elb.amazonaws.com";
    let PORT = 3000;
    let username = sessionStorage.getItem("username");
    axios
      .get(`http://${PAYMENT_HOST_ELB}:${PORT}/wallet/${username}`)
      .then(response => {
        console.log("Status Code GET Wallet:", response);
        this.setState({ wallet: response.data[0].wallet_amount });
        console.log(this.state.wallet);
      })
      .catch(err => {
        console.log(err);
      });
  }
}

export default Payment;
