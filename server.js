var path = require("path");
var webpack = require("webpack");
var DevServer = require("webpack-dev-server");
var config = require("./webpack.config");

config.entry.app = [
    "webpack-dev-server/client?http://localhost:3000",
    "webpack/hot/only-dev-server",
    path.join(__dirname, "react"),
];
config.plugins = [new webpack.HotModuleReplacementPlugin()];

var compiler = webpack(config);
var server = new DevServer(compiler, {
    publicPath: config.output.publicPath,
    hot: true,
});

server.listen(3000, "localhost");
