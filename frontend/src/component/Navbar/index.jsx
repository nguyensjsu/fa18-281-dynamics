import React, { Component } from "react";
import { Link, Redirect } from "react-router-dom";
import "./index.css";
import logo from "../../images/shayona-logob.png";
import storelogo from "../../images/store-logo.png";
import walleticon from "../../images/wallet-icon.png";

//create the Navbar Component
class Navbar extends Component {
  constructor() {
    super();
    this.userLogout = this.userLogout.bind(this);
  }
  userLogout() {
    sessionStorage.removeItem("username");
    this.forceUpdate();
  }
  render() {
    let username = sessionStorage.getItem("username");
    if (username === undefined) {
      return <Redirect to="/login" />;
    }
    return (
      <div>
        <nav id="homepage-nav" className="navbar navbar-default">
          <div className="container-fluid">
            <div className="navbar-header">
              <Link to="/" alt="homepage">
                <img
                  src={storelogo}
                  className="navbar-brand App-logo"
                  alt="store logo"
                />
                <img src={logo} className="navbar-brand App-logo" alt="logo" />
              </Link>
            </div>
            {username ? (
              <ul className="nav navbar-right">
                <li className="nav-item">
                  <Link
                    to="/inventory"
                    style={{
                      borderRadius: 5,
                      color: "#0067db",
                      marginRight: 20,
                      backgroundColor: "transparent"
                    }}
                    className="btn btn-primary"
                  >
                    Buy Groceries
                  </Link>
                </li>
                <li className="nav-item">
                  <Link
                    to="/wallet"
                    style={{
                      background: `url(${walleticon}) 10px 10px no-repeat`,
                      borderRadius: 5,
                      color: "#0067db",
                      marginRight: 20,
                      paddingLeft: 30
                    }}
                    className="btn btn-primary"
                  >
                    Shayona Wallet
                  </Link>
                </li>
                <li className="nav-item dropdown">
                  <button
                    id="login-btn"
                    className="btn btn-secondary dropdown-toggle"
                    type="button"
                    data-toggle="dropdown"
                    aria-haspopup="true"
                    aria-expanded="false"
                    style={{ color: "#0067db" }}
                  >
                    Hi, {username}
                  </button>
                  <div
                    className="dropdown-menu"
                    aria-labelledby="dropdownMenuButton"
                  >
                    <button onClick={this.userLogout} className="dropdown-item">
                      Logout
                    </button>
                  </div>
                </li>
              </ul>
            ) : (
              <ul className="nav navbar-right">
                <Redirect to="/" />
                <Link to="/login">
                  <button
                    id="login-btn"
                    className="btn btn-primary"
                    style={{ color: "#0067db" }}
                  >
                    Login
                  </button>
                </Link>
              </ul>
            )}
          </div>
        </nav>
      </div>
    );
  }
}

export default Navbar;
