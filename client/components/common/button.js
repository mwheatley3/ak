// http://bulma.io/documentation/elements/button/

import React, { Component, PropTypes, Children } from 'react';
import cx from 'classnames';
import { elementType } from 'react-prop-types';
import { Control } from './form';

const { string, bool, oneOf, node } = PropTypes;

export default class Button extends Component {
    static propTypes = {
        component: elementType,
        className: string,
        loading: bool,
        disabled: bool,
        inverted: bool,
        outlined: bool,
        type: string,
        size: oneOf(['small', 'normal', 'medium', 'large']),
    };

    static defaultProps = {
        component: 'a',
        type: 'primary',
        loading: false,
        disabled: false,
        inverted: false,
        outlined: false,
        size: 'normal',
    };

    render() {
        const { component: Comp, className, type, loading, disabled, inverted, outlined, size, ...rest } = this.props;

        const cls = cx(className, 'button', {
            ['is-' + type]: !!type,
            'is-loading': loading,
            'is-disabled': disabled,
            'is-inverted': inverted,
            'is-outlined': outlined,
            ['is-' + size]: !!size && size !== 'normal',
        });

        return <Comp className={ cls } { ...rest } />;
    }
}

export class ButtonGroup extends Component {
    static propTypes = {
        children: node,
        className: string,
        condensed: bool,
    };

    static defaultProps = {
        condensed: false,
    }

    render() {
        const { children, condensed, className } = this.props;
        const cls = cx(className, {
            'has-addons': condensed,
            'is-grouped': !condensed,
        });

        return (
            <Control className={ cls }>{
                Children.map(children, ch => condensed ? ch : <Control>{ ch }</Control>)
            }</Control>
        );
    }
}
