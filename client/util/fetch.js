import 'whatwg-fetch';
import Promise from 'bluebird';

export function urlEncode(params) {
    const p = [];

    for (const k in params) {
        p.push(encodeURIComponent(k) + '=' + encodeURIComponent(params[k]));
    }

    return p.join('&');
}

export function pathString(url, data) {
    data = urlEncode(data);

    if (!data) {
        return url;
    }

    return url + '?' + data;
}

export function handleResponse(resp, body, Type) {
    if (resp.status >= 400) {
        let msg = body;
        let data;

        if (body && body.error) {
            msg = body.error.message;
            data = body.error.data;
        }

        throw new FetchError(resp.status, msg, data);
    }

    return Type ? Type(body.data) : null;
}

function call({ url, options = {}, Type }) {
    if (!options.headers) {
        options.headers = {};
    }

    options.headers.Accept = 'application/json';
    options.credentials = 'same-origin';

    return Promise.resolve(fetch(url, options)
        .then(resp => (
            (resp.headers.get('Content-Type') === 'application/json' ? resp.json() : resp.text())
                .then(body => handleResponse(resp, body, Type))
        )));
}

export function get({ url, data = {}, Type = null }) {
    url = pathString(url, data);

    return call({ url, Type, options: {
        method: 'get',
    } });
}

export function post({ url, data = {}, Type = null }) {
    return call({ url, Type, options: {
        method: 'post',
        headers: {
            'Content-Type': 'application/json',
        },
        body: handleBody(data),
    } });
}

export function del({ url, data = {}, Type = null }) {
    return call({ url, Type, options: {
        method: 'delete',
        headers: {
            'Content-Type': 'application/json',
        },
        body: handleBody(data),
    } });
}

function handleBody(data) {
    if (data instanceof Blob) {
        return data;
    }

    return JSON.stringify(data);
}

export class FetchError extends Error {
    constructor(code, message, data) {
        super(message);

        this.statusCode = code;
        this.data = data;
    }
}
