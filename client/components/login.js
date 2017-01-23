import css from './login.css';
import React, { Component, PropTypes } from 'react';

import Notification from './common/notification';
import Form, { TextInput, PasswordInput, Submit } from './common/form';

import style from '../util/style';

import { withRouter } from 'react-router';
import { withStore } from 'client/store';
const { object } = PropTypes;


@withRouter
@style(css)
@withStore
export default class Login extends Component {
    static propTypes = {
        location: object.isRequired,
        router: object.isRequired,
        store: object.isRequired,
    };

    onSubmit = e => {
      const { store } = this.props;
        e.preventDefault();
        if (store.auth.user.loading) {
            return;
        }

        const { email, password } = this.refs;
        store.auth.login(email.value, password.value);
        this.checkAuth();
    }

    componentWillMount() {
        this.checkAuth();
    }

    componentDidUpdate() {
        this.checkAuth();
    }

    checkAuth() {
        const { router, location, store } = this.props;
        if (store.auth.loggedIn) {
            const next = location.query.next || '/';
            router.push(next);
        }
    }

    render() {
        const { store } = this.props;

        if (store.auth.loggedIn) {
            return null;
        }

        const isUnauthorized = store.auth.user.error && store.auth.user.error.statusCode === 401 || false;
        const loading = store.auth.user.loading;

        return (
            <div className="login">
                <Notification open={ isUnauthorized && !loading } type="danger">Invalid email or password</Notification>
                <Form className="box" onSubmit={ this.onSubmit }>
                    <h1 className="title">Login</h1>
                    <TextInput ref="email" placeholder="Email" />
                    <PasswordInput ref="password" placeholder="Password" />
                    <Submit loading={ false } disabled={ false }>Login</Submit>
                </Form>
            </div>
        );
    }
}
