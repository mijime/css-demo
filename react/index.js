import render from "./router/client";
import renderToString from "./router/server";

if (typeof window !== "undefined") {
    render(document.getElementById("root"));
} else {
    global.main = renderToString;
}
