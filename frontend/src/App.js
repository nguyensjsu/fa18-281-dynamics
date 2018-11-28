import React, { Component } from "react";
import { BrowserRouter } from "react-router-dom";
import Main from "./component/Main";

class App extends Component {
  render() {
    return (
      <BrowserRouter>
        <div>
          <Main />
        </div>
      </BrowserRouter>
    );
  }
}

export default App;
