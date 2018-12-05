import React, { Component } from "react";
import { Link } from "react-router-dom";
import "./index.css";
import logo from "../../images/shayona-logob.png";

//create the Navbar Component
class Navbar extends Component {
  constructor() {
    super();
    this.userLogout = this.userLogout.bind(this);
  }
  userLogout() {
    sessionStorage.removeItem("email");
    this.forceUpdate();
  }
  render() {
    let email = sessionStorage.getItem("email");
    return (
      <div>
        <nav id="homepage-nav" className="navbar navbar-default">
          <div className="container-fluid">
            <div className="navbar-header">
              <Link to="/" alt="homepage">
                <img src={logo} className="navbar-brand App-logo" alt="logo" />
              </Link>
            </div>
            {email ? (
              <ul className="nav navbar-right">
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
                    Hi, {email}
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
