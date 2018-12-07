import React, { Component } from "react";

class Wallet extends Component {
  constructor() {
    super();
    this.state = {
      wallet: 20,
      showPaymentOptions: false
    };
  }

  enablePaymentOptions = () => {
    this.setState({ showPaymentOptions: true });
  };

  amountChangeHandler = e => {
    this.setState({ wallet: e.target.value });
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
                <form method="POST">
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
                      onChange={this.amountChangeHandler}
                      type="number"
                      className="form-control form-control-lg rounded-0"
                      placeholder="Amount"
                    />
                  </div>
                  <button
                    onClick={this.addMoney} // to be implemented
                    className="btn btn-lg btn-block btn-login rounded-0"
                  >
                    Add Money
                  </button>
                </form>
              </div>
            </div>
          ) : (
            ""
          )}
        </div>
      </React.Fragment>
    );
  }
}

export default Wallet;
