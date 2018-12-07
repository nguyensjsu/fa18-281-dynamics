import React, { Component } from "react";
import banner1 from "../../images/banner-1.jpg";

class HomePage extends Component {
  render() {
    return (
      <React.Fragment>
        <div>
          <img src={banner1} alt="grocery store" />
        </div>
      </React.Fragment>
    );
  }
}

export default HomePage;
