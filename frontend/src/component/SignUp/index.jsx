import React, { Component } from "react";
import { Redirect } from "react-router";
import { Link } from "react-router-dom";
import "./index.css";

class SignUp extends Component {
  constructor(props) {
    super(props);
    this.state = {
      firstname: "",
      lastname: "",
      email: "",
      password: "",
      isowner: false,
      regFlag: false
    };
  }

  firstnameChangeHandler = e => {
    this.setState({
      firstname: e.target.value
    });
  };
  lastnameChangeHandler = e => {
    this.setState({
      lastname: e.target.value
    });
  };
  emailChangeHandler = e => {
    this.setState({
      email: e.target.value
    });
  };
  passwordChangeHandler = e => {
    this.setState({
      password: e.target.value
    });
  };

  render() {
    return (
      <div>
        {this.state.regFlag && <Redirect to={{ pathname: "/login" }} />}
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
                    onChange={this.emailChangeHandler}
                    type="text"
                    className="form-control form-control-lg rounded-0"
                    placeholder="Email address"
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
                  onClick={this.submitRegister} // to be implemented
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
