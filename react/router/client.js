import React from "react";
import {render} from "react-dom";
import {Router, browserHistory} from "react-router";
import {Provider} from "react-redux";
import {syncHistoryWithStore} from "react-router-redux";
import createRoutes from "./routes";
import createStore from "../store";

export default (element) => {
    const store = createStore(window["--app-initial"]);
    const history = syncHistoryWithStore(browserHistory, store);

    return render((<Provider store={store}>
                <Router history={history}>
                    {createRoutes({store, first: {time: true}})}
                </Router>
            </Provider>), element);
};
