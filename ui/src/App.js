import React, { Component } from "react";
import Chart from "./Chart";
import { Form } from "./Form";
import { Nav } from "./Nav";
import "./App.css";

class App extends Component {
  render() {
    return (
      <div>
        <Nav />
        <div className="container mt-2">
          <Form />
        </div>
      </div>
    );
  }
}

export default App;
