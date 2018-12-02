import React, { Component } from "react";
import { Link } from "react-router-dom";
import "./cart.css";

class Cart extends Component {
  constructor() {
    super();
    this.state = {
      cart: [
        {
          id: 1,
          name: "Cereals",
          rate: 10.95,
          inventory: 53,
          quantity: 10,
          subtotal: 110
        },
        {
          id: 2,
          name: "Bread",
          rate: 3.99,
          inventory: 11,
          quantity: 10,
          subtotal: 110
        },
        {
          id: 3,
          name: "Milk",
          rate: 3.16,
          inventory: 20,
          quantity: 10,
          subtotal: 110
        }
      ],
      total: 0,
      itemCount: 0
    };
  }
  render() {
    let item_list = this.state.cart.map(item => {
      return (
        <tr key={item.id}>
          <td className="item name col-lg-3">
            <Link to="">{item.name}</Link>
          </td>
          <td className="item rate col-lg-3">${item.rate}</td>
          <td className="item quantity col-lg-3">{item.quantity}</td>
          <td className="item subtotal col-lg-3">{item.subtotal}</td>
        </tr>
      );
    });
    return (
      <React.Fragment>
        <div className="container">
          <div id="title">
            <img
              alt="cart"
              src="data:image/png;base64,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"
            />
            <span className="page-header">Cart</span>
          </div>
          <table className="table table-hover">
            <thead>
              <tr>
                <th className="item name-header col-lg-3">Item</th>
                <th className="item rate-header col-lg-3">Rate</th>
                <th className="item quantity-header col-lg-3">Quantity</th>
                <th className="item subtotal-header col-lg-3">Subtotal</th>
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
              Total Items: {this.state.itemCount}
            </div>
            <div className="item-total col-lg-2">
              Order total:{" "}
              <span className="item-total-value">${this.state.total}</span>
            </div>
          </div>
          <div className="total-row row">
            <span className="col-lg-10" />
            <span className="col-lg-2">
              <button
                onClick={this.addToCart}
                className="btn btn-block btn-login rounded-0"
              >
                Place Order
              </button>
            </span>
          </div>
        </div>
      </React.Fragment>
    );
  }
}

export default Cart;
