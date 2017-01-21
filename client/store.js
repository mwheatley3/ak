import { provider as contextProvider } from './util/context';
import Value from './util/value';

export const withStore = contextProvider('store');

export default class Store {
    _inited = false;

    constructor(api) {
        this.auth = new AuthStore(api);
    }

    init() {
        if (this._inited) {
            return;
        }

        this._inited = true;
        this.auth.init();
    }
}

// manage auth state
class AuthStore {
    user = new Value();

    constructor(api) {
        this.api = api;
    }

    init() {
        // see if we get a user, but if we get an error don't worry about it
        this.user.trackPromise(this.api.getUser().catch(err => void 0));
    }

    get loggedIn() {
        return !!this.user.value;
    }

    login(username, password) {
        this.user.trackPromise(this.api.login(username, password));
    }

    logout() {
        this.user.setValue(null);
        this.api.logout().catch(err => console.err("logout error", err));
    }
}
