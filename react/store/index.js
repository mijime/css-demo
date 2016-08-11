import {combineReducers, createStore} from "redux";
import {routerReducer} from "react-router-redux";
import reducers from "../reducers";

/**
  * @param {Object} state is initial state
  * @returns {Store} return to redux store
  */
export default (state) => {
    return createStore(combineReducers({
        ...reducers,
        routing: routerReducer
    }), state);
};
