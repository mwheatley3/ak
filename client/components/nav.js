import React, { Component, PropTypes } from 'react';
import css from './nav.css';
import { withRouter, Link } from 'react-router';
import style from '../util/style';

const { object } = PropTypes;

@withRouter
@style(css)
export default class Nav extends Component {
    static propTypes = {
        router: object.isRequired,
    };

    render() {
        return (
            <nav className="nav">
                <div className="container">
                    <div className="nav-left">
                        <Link className="nav-item" to="/">nav left</Link>
                        <Link className="nav-item" to="/login">Login</Link>
                        <Link className="nav-item" to="/coffee">Coffee</Link>
                    </div>
                </div>
            </nav>
        );
    }
}
