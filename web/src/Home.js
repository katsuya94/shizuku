// @flow

import React from "react";
import classnames from "classnames";
import bulma from "bulma";

import styles from "./App.scss";

const Home = () => (
  <section className={bulma.hero}>
    <div className={classnames(bulma.heroBody, styles.banner)}>
      <div className={bulma.container}>
        <h1 className={bulma.title}>Adrien Tateno</h1>
        <h2 className={bulma.subtitle}>Developer</h2>
      </div>
    </div>
  </section>
);

export default Home;
