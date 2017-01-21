import React from 'react';
import { render } from 'react-dom';
import routes from './client/routes';

"use strict";

import Store from './client/store';
import API from './client/api';
import { injector } from './client/util/context';
const api = new API();
const store = new Store(api);
store.init();


const App = injector( { store: store } );

render(
    <App>{ routes }</App>,
    document.getElementById('root')
);
