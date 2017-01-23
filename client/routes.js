import React from 'react';
import { Router, Route, IndexRoute, browserHistory } from 'react-router';

import Container from 'client/components/container';
import Home from 'client/components/home';
import Sentiments from 'client/components/sentiments';
import Keri from 'client/components/keri';
import Login from 'client/components/login';
import { requireAuth } from 'client/components/hoc/require_auth';

export default (
    <Router history={ browserHistory }>
        <Route path="/" component={ Container }>
            <IndexRoute component={ Home } />
            <Route path="sentiments" component={ requireAuth(Sentiments) } />
            <Route path="keri" component={ requireAuth(Keri) } />
            <Route path="login" component={ Login } />
        </Route>
    </Router>
);
