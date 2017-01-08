import bcss from 'bulma/css/bulma.css';
import facss from 'font-awesome/css/font-awesome.css';
import css from './container.css';

import React, { Component, PropTypes } from 'react';
import style from '../util/style';

import Nav from './nav';

const { node } = PropTypes;

@style(bcss, facss, css)
export default class Container extends Component {
    static propTypes = {
        children: node,
    };

    render() {
        const { children } = this.props;

        return (
            <div className="container-container">
                <Nav />
                <div className="container">
                    { children }
                </div>
            </div>
        );
    }
}
