// @flow

import React from "react";
import { Link, Switch, Route, BrowserRouter } from "react-router-dom";
import classnames from "classnames";
import bulma from "bulma";

import Home from "./Home";
import Dashboard from "./Dashboard";

const App = () => (
  <BrowserRouter>
    <div>
      <nav className={bulma.navbar} role="navigation">
        <div className={bulma.navbarBrand}>
          <Link to="/" className={bulma.navbarItem}>
            <span
              className={classnames(bulma.isSize3, bulma.hasTextWeightBold)}
            >
              at.io
            </span>
          </Link>
        </div>
      </nav>
      <Switch>
        <Route exact path="/" component={Home} />
        <Route path="/dashboard" component={Dashboard} />
      </Switch>
    </div>
  </BrowserRouter>
);

export default App;
