import React from "react";
import {render} from "react-dom";
import {createStore, combineReducers} from "redux";
import {Provider} from "react-redux";
import {Router, Route, browserHistory} from "react-router";
import {syncHistoryWithStore, routerReducer} from "react-router-redux";
import * as reducers from "./reducers";
import App from "./containers/app";
import Bulma from "./containers/bulma";
import BootstrapSass from "./containers/bootstrap-sass";
import Milligram from "./containers/milligram";

const store = createStore(combineReducers({
    ...reducers,
    routing: routerReducer
}));

const history = syncHistoryWithStore(browserHistory, store);

render((
    <Provider store={store}>
        <Router history={history}>
            <Route path="/" component={App}>
                <Route path="/bulma" component={Bulma} />
                <Route path="/bootstrap-sass" component={BootstrapSass} />
                <Route path="/milligram" component={Milligram} />
            </Route>
        </Router>
    </Provider>
), document.getElementById("root"));
