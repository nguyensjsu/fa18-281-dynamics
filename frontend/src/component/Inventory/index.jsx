import React, { Component } from "react";
import { Link } from "react-router-dom";
import "./index.css";

class Inventory extends Component {
  constructor() {
    super();
    this.state = {
      inventory: [
        {
          id: 1,
          name: "Cereals",
          rate: 10.95,
          inventory: 53,
          quantity: 0,
          subtotal: 0
        },
        {
          id: 2,
          name: "Bread",
          rate: 3.99,
          inventory: 11,
          quantity: 0,
          subtotal: 0
        },
        {
          id: 3,
          name: "Milk",
          rate: 3.16,
          inventory: 20,
          quantity: 0,
          subtotal: 0
        }
      ],
      total: 0,
      itemCount: 0
    };
  }

  quantityChangeHandler = (e, id) => {
    let inventory = this.state.inventory;
    for (let item of inventory) {
      if (item.id === id) {
        if (e.target.value <= item.inventory) {
          item.quantity = parseInt(e.target.value);
          item.subtotal = parseFloat((item.quantity * item.rate).toFixed(2));
        }
      }
    }
    this.setState({ inventory: inventory });
    this.totalChangeHandler();
    this.itemCountChangeHandler();
  };

  totalChangeHandler() {
    let inventory = this.state.inventory;
    let total = 0;
    for (let item of inventory) {
      if (item.subtotal === undefined) item.subtotal = 0;
      total += item.subtotal;
    }
    this.setState({ total: parseFloat(total.toFixed(2)) });
  }

  itemCountChangeHandler() {
    let inventory = this.state.inventory;
    let itemCount = 0;
    for (let item of inventory) {
      itemCount += item.quantity;
    }
    this.setState({ itemCount: parseFloat(itemCount.toFixed(2)) });
  }

  render() {
    let item_list = this.state.inventory.map(item => {
      return (
        <tr key={item.id}>
          <td className="item name col-lg-4">
            <Link to="">{item.name}</Link>
          </td>
          <td className="item rate col-lg-2">${item.rate}</td>
          <td className="item inventory col-lg-2">{item.inventory}</td>
          <td className="item quantity col-lg-2">
            <input
              value={item.quantity}
              min="0"
              max={item.inventory}
              onChange={e => this.quantityChangeHandler(e, item.id)}
              type="number"
            />
          </td>
          <td className="item subtotal col-lg-2">{item.subtotal}</td>
        </tr>
      );
    });
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
              Order total: ${this.state.total}
            </div>
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
