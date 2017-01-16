import React from 'react';
import { Router, Route, IndexRoute, browserHistory } from 'react-router';

import Container from './client/components/container.js';
import Home from './client/components/home.js';
import Coffee from './client/components/coffee.js';
import Login from './client/components/login';

export default (
    <Router history={ browserHistory }>
        <Route path="/" component={ Container }>
            <IndexRoute component={ Home } />
            <Route path="/coffee" component={ Coffee } />
            <Route path="/login" component={ Login } />
        </Route>
    </Router>
);
