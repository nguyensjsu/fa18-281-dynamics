import React, { Component } from "react";
import { Link } from "react-router-dom";
import uniqid from "uniqid";
import "./index.css";

class Inventory extends Component {
  render() {
    return (
      <React.Fragment>
        <div className="container">
          <div id="title">
            <h2>Inventory</h2>
          </div>
          <table className="table table-hover">
            <thead>
              <tr>
                <th className="item name-header col-lg-4">Item</th>
                <th className="item rate-header col-lg-2">Rate</th>
                <th className="item inventory-header col-lg-2">
                  Items Available
                </th>
                <th className="item quantity-header col-lg-2">Quantity</th>
                <th className="item subtotal-header col-lg-2">Subtotal</th>
              </tr>
            </thead>
            <tbody
              key={uniqid()}
              className="card-parent table table-hover container grocery-item"
              style={{ margin: "auto" }}
            >
              <tr>
                <td className="item name col-lg-4">
                  <Link to="">Cereals</Link>
                </td>
                <td className="item rate col-lg-2">$10.95</td>
                <td className="item inventory col-lg-2">55</td>
                <td className="item quantity col-lg-2">
                  <input
                    placeholder="0"
                    min="0"
                    max="10"
                    onChange={this.quantityChangeHandler}
                    type="number"
                  />
                </td>
                <td className="item subtotal col-lg-2">$123</td>
              </tr>
              <tr>
                <td className="item name col-lg-4">
                  <Link to="">Cereals</Link>
                </td>
                <td className="item rate col-lg-2">$10.95</td>
                <td className="item inventory col-lg-2">55</td>
                <td className="item quantity col-lg-2">
                  <input
                    placeholder="0"
                    min="0"
                    max="10"
                    onChange={this.quantityChangeHandler}
                    type="number"
                  />
                </td>
                <td className="item subtotal col-lg-2">$123</td>
              </tr>
              <tr>
                <td className="item name col-lg-4">
                  <Link to="">Cereals</Link>
                </td>
                <td className="item rate col-lg-2">$10.95</td>
                <td className="item inventory col-lg-2">55</td>
                <td className="item quantity col-lg-2">
                  <input
                    placeholder="0"
                    min="0"
                    max="10"
                    onChange={this.quantityChangeHandler}
                    type="number"
                  />
                </td>
                <td className="item subtotal col-lg-2">$123</td>
              </tr>
            </tbody>
          </table>
          <div className="total-row row">
            <span className="col-lg-8" />
            <div className="item-count col-lg-2">Item Count: 10</div>
            <div className="item-total col-lg-2">Total: $111</div>
          </div>
          <div className="total-row row">
            <span className="col-lg-10" />
            <span className="col-lg-2">
              <button
                onClick={this.addToCart}
                className="btn btn-block btn-login rounded-0"
              >
                Add to Cart
              </button>
            </span>
          </div>
        </div>
      </React.Fragment>
    );
  }
}

export default Inventory;
