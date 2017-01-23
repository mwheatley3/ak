import React, { Component, PropTypes } from 'react';
import css from './nav.css';
import { withRouter, Link } from 'react-router';
import style from '../util/style';
import Button from 'client/components/common/button';
import { withStore } from 'client/store';

const { object } = PropTypes;

@withRouter
@style(css)
@withStore
export default class Nav extends Component {
    static propTypes = {
        router: object.isRequired,
        store: object.isRequired,
    };

    onLogoutClick() {
        const { store, router } = this.props;

        store.auth.logout();
        router.push('/login');
    }

    render() {
      const { store } = this.props;
        return (
            <nav className="nav">
                <div className="container">
                    <div className="nav-left">
                        <Link className="nav-item" to="/">nav left</Link>
                        <Link className="nav-item" to="/sentiments">Sentiments</Link>
                        <Link className="nav-item" to="/keri">Keri</Link>
                    </div>
                    <div className="nav-right">{
                        store.auth.loggedIn ? [
                            <span key="1" className="nav-item">{ store.auth.user.value.email }</span>,
                            <span key="2" className="nav-item"><Button onClick={ () => this.onLogoutClick() }>Log Out</Button></span>,
                        ] : <span key="3" className="nav-item"><Button component={ Link } to="/login">Log In</Button></span>
                    }</div>
                </div>
            </nav>
        );
    }
}
