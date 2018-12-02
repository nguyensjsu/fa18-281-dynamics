import React, { Component } from "react";
import "./orders.css";

class Orders extends Component {
  render() {
    return (
      <React.Fragment>
        <div className="container">
          <div id="title">
            <span className="page-header">Orders</span>
          </div>
          <hr />
          <div className="card">
            <div className="card-header order-header">
              <span className="order-status">Order Placed</span>
              <span className="order-total">Total: $135</span>
              <span className="order-number">ORDER # 113-2127190-1337014</span>
            </div>
            <div className="card-body">
              <div>
                <ol>
                  <li>Cereals</li>
                  <li>Bread</li>
                  <li>Milk</li>
                </ol>
              </div>
            </div>
          </div>
        </div>
      </React.Fragment>
    );
  }
}

export default Orders;
