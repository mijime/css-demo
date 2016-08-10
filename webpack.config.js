var path = require("path");
var webpack = require("webpack");

module.exports = {
    entry: {
        app: [path.join(__dirname, "react")],
    },
    output: {
        path: path.join(__dirname, "build/Release/assets"),
        filename: "[name].bundle.js",
    },
    module: {
        loaders: [{
            test: /\.js$/,
            loaders: ["babel", "eslint"]
        }, {
            test: /\.css$/,
            loaders: ["css", "postcss"]
        }, {
            test: /\.s[ac]ss$/,
            loaders: ["css", "postcss", "sass"]
        }, {
            test: /\.jade$/,
            loaders: ["jade-react"]
        }, {
            test: /\.woff2?(\?v=[0-9]\.[0-9]\.[0-9])?$/,
            loaders: ["url?limit=10000&mimetype=application/font-woff"]
        }, {
            test: /\.(ttf|eot|svg)(\?v=[0-9]\.[0-9]\.[0-9])?$/,
            loaders: ["file"]
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
