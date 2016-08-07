var path = require("path");
var webpack = require("webpack");

module.exports = {
    entry: {
        app: [path.join(__dirname, "src")],
    },
    output: {
        path: path.join(__dirname, "build/Release"),
        filename: "[name].bundle.js",
    },
    module: {
        loaders: [{
            test: /\.js$/,
            loaders: ["babel", "eslint"]
        }, {
            test: /\.css$/,
            loaders: ["style", "css", "postcss"]
        }, {
            test: /\.jade$/,
            loaders: ["jade-react"]
        }]
    },
    cache: true,
    devtool: "#source-map",
    plugins: [
        new webpack.optimize.UglifyJsPlugin(),
        new webpack.DefinePlugin({
            "process.env": {
                NODE_ENV: JSON.stringify("production")
            }
        }),
    ]
};
