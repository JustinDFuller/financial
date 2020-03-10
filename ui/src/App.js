import React, { Component } from "react";
import Chart from "./Chart";
import Form from "./Form";
import "./App.css";

class App extends Component {
  render() {
    return [<Chart />, <Form />];
  }
}

export default App;
