const path = require("path");

module.exports = {
  entry: "./src/index.js",
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /(node_modules|bower_components)/,
        use: {
          loader: "babel-loader",
          options: {
            presets: ["react"]
          }
        }
      },
      {
        test: /\.(scss|sass)$/,
        use: ["babel-loader", "css-loader", "sass-loader"]
      }
    ]
  },
  output: {
    filename: "index.js",
    path: path.resolve(__dirname, "assets")
  },
  devServer: {
    host: "0.0.0.0",
    publicPath: "/assets/",
    contentBase: path.resolve(__dirname, "public"),
    port: 9000
  }
};
