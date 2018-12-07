import React, { Component } from "react";
import { Route } from "react-router-dom";
import Navbar from "./Navbar";
import SignUp from "./SignUp";
import Login from "./Login";
import Inventory from "./Inventory";
import ItemDetails from "./ItemDetails";
import Cart from "./Cart";
import Orders from "./Orders";
import HomePage from "./HomePage";

//Create a Main Component
class Main extends Component {
  render() {
    return (
      <div>
        {/*Render Different Component based on Route*/}
        <Route path="/" component={Navbar} />
        <Route path="/" exact component={HomePage} />
        <Route path="/login" exact component={Login} />
        <Route path="/signup" exact component={SignUp} />
        <Route path="/inventory" exact component={Inventory} />
        <Route path="/itemdetails" exact component={ItemDetails} />
        <Route path="/cart" exact component={Cart} />
        <Route path="/orders" exact component={Orders} />
      </div>
    );
  }
}

// Export The Main Component
export default Main;
