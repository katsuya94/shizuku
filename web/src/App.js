// @flow

import React from "react";
import { Link, Switch, Route, BrowserRouter } from "react-router-dom";
import classnames from "classnames";
import { navbar, navigation, navbarBrand, navbarItem, isSize3, hasTextWeightBold, section } from "bulma/bulma.sass";

import Home from "./Home";
import Dashboard from "./Dashboard";

const App = () => (
  <BrowserRouter>
    <div>
      <nav className={navbar} role="navigation">
        <div className={navbarBrand}>
          <Link to="/" className={navbarItem}>
            <span className={classnames(isSize3, hasTextWeightBold)}>at.io</span>
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
