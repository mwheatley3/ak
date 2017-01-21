import React, { Component, PropTypes } from 'react';
import { withRouter } from 'react-router';
import { withStore } from '../../store';

const { object, node } = PropTypes;

export function requireAuth(Comp) {
    class RequireAuthHOC extends Component {
      render() {
        return (
          <RequireAuth>
              <Comp { ...this.props } />
          </RequireAuth>
      );
    }
  }
  return RequireAuthHOC;
}

@withRouter
@withStore
export class RequireAuth extends Component {
    static propTypes = {
        location: object.isRequired,
        store: object.isRequired,
        router: object.isRequired,
        children: node.isRequired,
    };

    constructor(...args) {
        super(...args);

        this.state = {
            authed: false,
        };
    }

    checkAuth() {
        const { location, router, store } = this.props;

        if (!store.auth.user.loaded) {
            return;
        }

        if (store.auth.loggedIn) {
            this.setState({ authed: true });
        } else {
            const next = router.createPath({ pathname: location.pathname, query: location.query });
            router.push({ pathname: '/login', query: { next } });
        }
    }

    componentWillMount() {
        this.checkAuth();
    }

    componentWillUnmount() {
        this.checkAuth();
    }

    render() {
        if (!this.state.authed) {
            return null;
        }

        return this.props.children;
    }
}
