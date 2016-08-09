import {Component} from "react";
import {connect} from "react-redux";
import view from "./view.jade";
import styles from "./style.scss";

class Bulma extends Component {
    render() {
        const style = styles.map(s => s[1]).join("");
        const {children} = this.props;

        return view({
            style,
            children
        });
    }
}

export default connect(state => state)(Bulma);
