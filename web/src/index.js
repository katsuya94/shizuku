// @flow

import React from "react";
import ReactDOM from "react-dom";
import { Switch, Route, BrowserRouter } from "react-router-dom";
import "bulma/bulma.sass";

import Home from "./Home";
import Dashboard from "./Dashboard";

const root = document.getElementById("root");
if (root) {
  ReactDOM.render(
    <BrowserRouter>
      <Switch>
        <Route exact path="/" component={Home} />
        <Route path="/dashboard" component={Dashboard} />
      </Switch>
    </BrowserRouter>,
    root
  );
}
