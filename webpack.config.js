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
            loaders: ["css", "postcss"]
        }, {
            test: /\.s[ac]ss$/,
            loaders: ["css", "postcss", "sass"]
        }, {
            test: /\.jade$/,
            loaders: ["jade-react"]
        }, {
            test: /\.woff(\?v=\d+\.\d+\.\d+)?$/,
            loader: "url?limit=10000&mimetype=application/font-woff"
        }, {
            test: /\.woff2(\?v=\d+\.\d+\.\d+)?$/,
            loader: "url?limit=10000&mimetype=application/font-woff"
        }, {
            test: /\.ttf(\?v=\d+\.\d+\.\d+)?$/,
            loader: "url?limit=10000&mimetype=application/octet-stream"
        }, {
            test: /\.eot(\?v=\d+\.\d+\.\d+)?$/,
            loader: "file"
        }, {
            test: /\.svg(\?v=\d+\.\d+\.\d+)?$/,
            loader: "url?limit=10000&mimetype=image/svg+xml"
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
