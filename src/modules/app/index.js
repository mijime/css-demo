import {Component} from "react";
import view from "./view.jade";
import "./style.css";

export default class App extends Component {
    render() {
        return view();
    }
}
