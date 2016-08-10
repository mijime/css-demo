import {Component} from "react";
import {Link} from "react-router";
import {connect} from "react-redux";
import view from "./view.jade";
import styles from "./style.scss";

class App extends Component {
    render() {
        const {children} = this.props;
        const style = styles.map(s => s[1]).join("");
        const {headerClass} = styles.locals;

        return view({
            style,
            headerClass,
            children,
            Link
        });
    }
}

export default connect(state => state)(App);
