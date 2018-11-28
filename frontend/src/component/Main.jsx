import React, { Component } from "react";
import { Route } from "react-router-dom";
import Navbar from "./Navbar";
import SignUp from "./SignUp";
import Login from "./Login";

//Create a Main Component
class Main extends Component {
  render() {
    return (
      <div>
        {/*Render Different Component based on Route*/}
        <Route path="/" exact component={Navbar} />
        <Route path="/login" exact component={Login} />
        <Route path="/signup" exact component={SignUp} />
      </div>
    );
  }
}

// Export The Main Component
export default Main;
