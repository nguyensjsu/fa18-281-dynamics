import React, { Component } from "react";
import { Redirect } from "react-router";
import { Link } from "react-router-dom";
import axios from "axios";
import "./index.css";

class SignUp extends Component {
  constructor(props) {
    super(props);
    this.state = {
      firstname: "",
      lastname: "",
      username: "",
      password: "",
      userRegistered: false
    };
  }

  firstnameChangeHandler = e => {
    console.log(this.state.firstname);
    this.setState({
      firstname: e.target.value
    });
  };
  lastnameChangeHandler = e => {
    this.setState({
      lastname: e.target.value
    });
  };
  usernameChangeHandler = e => {
    this.setState({
      username: e.target.value
    });
  };
  passwordChangeHandler = e => {
    this.setState({
      password: e.target.value
    });
  };

  userRegistration = e => {
    e.preventDefault();
    let USERS_GOAPI_ELB =
      "http://Shayona-GOAPI-ELB-1280633407.us-west-2.elb.amazonaws.com";
    let PORT = 3000;
    console.log(this.state.username);
    let data = {
      firstname: this.state.firstname,
      lastname: this.state.lastname,
      username: this.state.username,
      password: this.state.password
    };
    axios
      .post(`${USERS_GOAPI_ELB}:${PORT}/users`, data)
      .then(response => {
        console.log("Status Code : ", response);
        if (response.status === 201) {
          console.log("response data:", response);
          this.setState({
            userRegistered: true
          });
        }
      })
      .catch(err => {
        console.log("error", err);
        this.setState({
          userRegistered: false
        });
      });
  };

  render() {
    return (
      <div>
        {this.state.userRegistered && <Redirect to="/login" />}
        <div className="container">
          <h4
            style={{
              textAlign: "center"
            }}
          >
            Already have an account?{" "}
            <Link to={{ pathname: "/login" }}>Login</Link>
            <br />
          </h4>
          <br />
          <div className="login-form card w-50 rounded-0">
            <div className="card-header card-title">
              <h3
                style={{
                  textAlign: "center"
                }}
              >
                User Sign Up
              </h3>
            </div>
            <div className="card-body">
              <form method="POST">
                <div className="form-row">
                  <div className="col form-group">
                    <input
                      onChange={this.firstnameChangeHandler}
                      type="text"
                      className="form-control form-control-lg rounded-0"
                      placeholder="First Name"
                    />
                  </div>
                  <div className="col form-group">
                    <input
                      onChange={this.lastnameChangeHandler}
                      type="text"
                      className="form-control form-control-lg rounded-0"
                      placeholder="Last Name"
                    />
                  </div>
                </div>
                <div className="form-group">
                  <input
                    onChange={this.usernameChangeHandler}
                    type="text"
                    className="form-control form-control-lg rounded-0"
                    placeholder="Username"
                  />
                </div>
                <div className="form-group">
                  <input
                    onChange={this.passwordChangeHandler}
                    type="password"
                    className="form-control form-control-lg rounded-0"
                    placeholder="Password"
                  />
                </div>
                <h6>Forgot password?</h6>
                <button
                  onClick={this.userRegistration} // to be implemented
                  className="btn btn-lg btn-block btn-login rounded-0"
                >
                  Sign Me Up
                </button>
              </form>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default SignUp;
