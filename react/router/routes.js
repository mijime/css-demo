import React from "react";
import {Route, IndexRoute} from "react-router";
import App from "../containers/app";
import Bulma from "../containers/bulma";
import BootstrapSass from "../containers/bootstrap-sass";
import Milligram from "../containers/milligram";

/**
 * Return route
 *
 * @param {Object} options any
 * @returns {Route} return to route
 */
export default ({store, first}) => {

    /**
      * @param {loader} loader is xxx
      * @returns {store} xxx
      */
    function w(loader) {
        return (nextState, replaceState, done) => {
            if (first.time) {
                first.time = false;
                return done();
            }
            return loader ?
                loader({store, nextState, replaceState, done}) :
                done();
        };
    }

    return (<Route path="/" component={App}>
                <IndexRoute component={Bulma} onEnter={w(Bulma.onEnter)} />
                <Route path="/bulma" component={Bulma} onEnter={w(Bulma.onEnter)} />
                <Route path="/bootstrap-sass" component={BootstrapSass} onEnter={w(BootstrapSass.onEnter)} />
                <Route path="/milligram" component={Milligram} onEnter={w(Milligram.onEnter)} />
            </Route>);
};
