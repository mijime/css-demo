import React from "react";
import {Provider} from "react-redux";
import {renderToString} from "react-dom/server";
import {match, RouterContext} from "react-router";
import Helmet from "react-helmet";
import createRoutes from "./routes";
import createStore from "../store";

/**
 * Handle HTTP request at Golang server
 *
 * @param   {Object}   options  request options
 * @param   {Function} done     response callback
 * @returns {none} return use callback
 */
export default function({uuid, url}, done) {
    const store = createStore();

    try {
        return match({
            routes: createRoutes({
                store,
                first: {time: false}
            }),
            location: url,
        }, (error, redirectLocation, renderProps) => {
            try {
                if (error) {
                    return done(JSON.stringify({
                        uuid,
                        error,
                    }));

                } else if (redirectLocation) {
                    const redirect = redirectLocation.pathname + redirectLocation.search;
                    return done(JSON.stringify({
                        uuid,
                        redirect,
                    }));

                } else {
                    const app = renderToString(
                                <Provider store={store}>
                                    <RouterContext {...renderProps} />
                                </Provider>);
                    const {title, meta} = Helmet.rewind();
                    const initial = JSON.stringify(store.getState());
                    return done(JSON.stringify({
                        uuid,
                        app,
                        title,
                        meta,
                        initial,
                    }));
                }
            } catch (e) {
                return done(JSON.stringify({
                    uuid,
                    error: e,
                }));
            }
        });
    } catch (error) {
        return done(JSON.stringify({
            uuid,
            error,
        }));
    }
}
