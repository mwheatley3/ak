import React, { Component, PropTypes } from 'react';
import cx from 'classnames';

const { node, string, bool, func } = PropTypes;

export default class Notification extends Component {
    static propTypes = {
        children: node,
        type: string,
        open: bool,
        onCloseClick: func,
        className: string,
    };

    static defaultProps = {
        open: true,
    };

    render() {
        const { type, onCloseClick, children, open, className, ...rest } = this.props;

        if (!open) {
            return null;
        }

        const cls = cx('notification', { ['is-' + type]: !!type }, className);

        return (
            <div className={ cls } { ...rest }>
                { onCloseClick ? <button className="delete" onClick={ onCloseClick }/> : null }
                { children }
            </div>
        );
    }
}
