import React, { Component } from "react";
import axios from "axios";

class Wallet extends Component {
  constructor() {
    super();
    this.state = {
      wallet: 0,
      add_money: 0,
      showPaymentOptions: false
    };
  }

  enablePaymentOptions = () => {
    this.setState({ showPaymentOptions: true });
  };

  addAmountChangeHandler = e => {
    this.setState({ add_money: e.target.value });
  };

  addMoney = e => {
    let PAYMENT_HOST_ELB = "payments-1051217824.us-west-2.elb.amazonaws.com";
    let PORT = 3000;
    let data = {
      username: sessionStorage.getItem("username"),
      wallet_amount: parseInt(this.state.add_money)
    };
    console.log("add money:", data);
    axios
      .put(`http://${PAYMENT_HOST_ELB}:${PORT}/wallet/add`, data)
      .then(response => {
        console.log("response from PUT ADD MONEY", response);
        this.setState({
          wallet: response.data.wallet_amount
        });
      });
  };

  render() {
    return (
      <React.Fragment>
        <div className="container">
          <div id="title">
            <h2>Shayona Wallet</h2>
            <hr />
          </div>
          <span style={{ margin: 20, fontSize: 25, fontWeight: 500 }}>
            Wallet:{" "}
            <span style={{ fontWeight: 700 }}>${this.state.wallet}</span>
          </span>
          {!this.state.showPaymentOptions ? (
            <span style={{ margin: 20 }}>
              <button
                onClick={this.enablePaymentOptions}
                className="btn btn-primary"
              >
                Add Money
              </button>
            </span>
          ) : (
            ""
          )}
          {this.state.showPaymentOptions ? (
            <div
              style={{ margin: 20, width: 650 }}
              className="login-form card rounded-0"
            >
              <div className="card-header card-title">
                <h3
                  style={{
                    textAlign: "center"
                  }}
                >
                  Add Card Information
                </h3>
              </div>
              <div className="card-body">
                <div>
                  <div className="form-group">
                    <input
                      type="number"
                      className="form-control form-control-lg rounded-0"
                      placeholder="Card Number"
                    />
                  </div>
                  <div className="form-row">
                    <div className="col form-group">
                      <input
                        type="number"
                        className="form-control form-control-lg rounded-0"
                        placeholder="Expiration Month"
                      />
                    </div>
                    <div className="col form-group">
                      <input
                        type="number"
                        className="form-control form-control-lg rounded-0"
                        placeholder="Expiration Year"
                      />
                    </div>
                    <div className="col form-group">
                      <input
                        type="number"
                        className="form-control form-control-lg rounded-0"
                        placeholder="CVV"
                      />
                    </div>
                  </div>
                  <div className="form-group">
                    <input
                      type="text"
                      className="form-control form-control-lg rounded-0"
                      placeholder="Name on the Card"
                    />
                  </div>
                  <div className="form-group">
                    <input
                      type="password"
                      className="form-control form-control-lg rounded-0"
                      placeholder="Billing Address"
                    />
                  </div>
                  <div className="form-group">
                    <input
                      onChange={this.addAmountChangeHandler}
                      type="number"
                      className="form-control form-control-lg rounded-0"
                      placeholder="Amount"
                    />
                  </div>
                  <button
                    onClick={this.addMoney}
                    className="btn btn-lg btn-block btn-login rounded-0"
                  >
                    Add Money
                  </button>
                </div>
              </div>
            </div>
          ) : (
            ""
          )}
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
        if (response.status === 204) {
          // user wallet doesn't exist
          let data = {
            username: sessionStorage.getItem("username"),
            wallet_amount: this.state.wallet
          };
          axios
            .post(`http://${PAYMENT_HOST_ELB}:${PORT}/wallet`, data)
            .then(response => {
              console.log("Status Code POST Wallet:", response.status);
              console.log("response from POST Wallet:", response);
              this.setState({
                wallet: response.data.wallet_amount
              });
            });
        } else if (response.status === 200) {
          this.setState({
            wallet: response.data[0].wallet_amount
          });
        }
      })
      .catch(err => {
        console.log(err);
      });
  }
}

export default Wallet;
