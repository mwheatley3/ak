import React from 'react';
import { render } from 'react-dom';
import App from './client/App';
import routes from './routes';

render(
    <App>{ routes }</App>,
    document.getElementById('root')
);
