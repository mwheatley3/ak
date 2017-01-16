import React, { Component, PropTypes } from 'react';
import Button from './button';
import cx from 'classnames';
import { elementType } from 'react-prop-types';

const { string, bool, oneOf, node } = PropTypes;

export default class Form extends Component {
    render() {
        return (
            <form { ...this.props } />
        );
    }
}

export class Label extends Component {
    static propTypes = {
        className: string,
        horizontal: bool,
    };

    render() {
        const { className, horizontal, ...rest } = this.props;
        const cls = cx('label', className);
        let el = <label className={ cls } { ...rest } />;

        if (horizontal) {
            el = <div className="control-label">{ el }</div>;
        }

        return el;
    }
}

export class Control extends Component {
    static propTypes = {
        component: elementType,
        horizontal: bool,
        className: string,
    };

    static defaultProps = {
        component: 'div',
    };

    render() {
        const { component: Comp, horizontal, className, ...rest } = this.props;
        const cls = cx('control', { 'is-horizontal': horizontal }, className);

        return <Comp className={ cls } { ... rest } />;
    }
}

export class TextInput extends Component {
    static propTypes = {
        className: string,
        icon: node,
        iconPlacement: oneOf(['left', 'right']),
        size: oneOf(['small', 'normal', 'medium', 'large']),
        addOnRight: node,
        addOnLeft: node,
    };

    static defaultProps = {
        iconPlacement: 'right',
        size: 'normal',
    };

    get value() {
        return this.refs.input.value;
    }

    focus() {
        this.refs.input.focus();
    }

    render() {
        const { className, icon, iconPlacement, size, ...rest } = this.props;
        delete rest.type;
        const cls = cx('input', className, {
            ['is-' + size]: size === 'small' || size === 'medium' || size === 'large',
        });
        const ctrlCls = cx({
            'has-icon': !!icon,
            'has-icon-right': !!icon && iconPlacement === 'right',
        });

        return (
            <Control className={ ctrlCls }>
                <input ref="input" className={ cls } type="text" { ...rest } />
                { icon }
            </Control>
        );
    }
}

export class TextAreaInput extends Component {
    static propTypes = {
        className: string,
        autosize: bool,
    };

    static defaultProps = {
        autosize: false,
    };

    get value() {
        return this.refs.input.value;
    }

    focus() {
        this.refs.input.focus();
    }

    render() {
        const { className, ...rest } = this.props;
        const cls = cx('textarea', className);
        const Comp = 'textarea';

        return (
            <Control>
                <Comp ref="input" className={ cls } { ...rest } />
            </Control>
        );
    }
}

export class CheckboxInput extends Component {
    get value() {
        return this.refs.input.value;
    }

    focus() {
        this.refs.input.focus();
    }

    render() {
        return (
            <Control>
                <input ref="input" type="checkbox" { ...this.props } />
            </Control>
        );
    }
}

export class PasswordInput extends Component {
    static propTypes = {
        className: string,
    };

    get value() {
        return this.refs.input.value;
    }

    focus() {
        this.refs.input.focus();
    }

    render() {
        const { className, ...rest } = this.props;
        delete rest.type;
        const cls = cx('input', className);

        return (
            <Control>
                <input ref="input" className={ cls } type="password" { ...this.props } />
            </Control>
        );
    }
}

export class Select extends Component {
    static propTypes = {
        className: string,
        size: oneOf(['small', 'normal', 'medium', 'large']),
    };

    static defaultProps = {
        size: 'normal',
    };

    render() {
        const { className, size, ...rest } = this.props;
        const cls = cx('select', className, {
            ['is-' + size]: size === 'small' || size === 'medium' || size === 'large',
        });

        return (
            <Control>
                <span className={ cls } { ...rest }>
                    <select { ...this.props } />
                </span>
            </Control>
        );
    }
}

export class Submit extends Component {
    render() {
        return (
            <Control>
                <Button component="button" { ...this.props } />
            </Control>
        );
    }
}
