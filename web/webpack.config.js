const path = require("path");
const webpack = require("webpack");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const OptimizeCssAssetsPlugin = require("optimize-css-assets-webpack-plugin");
const UglifyJSPlugin = require("uglifyjs-webpack-plugin");

process.env.NODE_ENV = process.env.NODE_ENV || "development";

module.exports = {
  mode: process.env.NODE_ENV,
  entry: "./src/index.js",
  plugins: [
    new webpack.EnvironmentPlugin(["NODE_ENV"]),
    new HtmlWebpackPlugin({
      title: "atateno.io",
      template: require("html-webpack-template"),
      appMountId: "root",
      minify: process.env.NODE_ENV === "production"
    })
  ],
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
        use: [
          process.env.NODE_ENV === "production"
            ? MiniCssExtractPlugin.loader
            : "style-loader",
          "css-loader",
          "sass-loader"
        ]
      }
    ]
  },
  output: {
    filename: "assets/index.js",
    path: path.resolve(__dirname, "dist")
  },
  devtool: "source-map",
  devServer: {
    host: "0.0.0.0",
    port: 9000
  }
};

if (process.env.NODE_ENV === "production") {
  module.exports.plugins.push(
    new MiniCssExtractPlugin({
      filename: "[name].[hash].css",
      chunkFilename: "[id].[hash].css"
    }),
    new OptimizeCssAssetsPlugin({}),
    new UglifyJSPlugin({
      sourceMap: true
    })
  );
}
