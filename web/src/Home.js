// @flow

import React from "react";
import classnames from "classnames";
import { hero, heroBody, container, title, subtitle } from "bulma/bulma.sass";

import { atBanner } from "./App.scss"

const Home = () => (
  <section className={hero}>
    <div className={classnames(heroBody, atBanner)}>
      <div className={container}>
        <h1 className={title}>Adrien Tateno</h1>
        <h2 className={subtitle}>Developer</h2>
      </div>
    </div>
  </section>
);

export default Home;
