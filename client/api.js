import { get, del, post } from './util/fetch';
import { User } from 'client/models';

export default class API {
    constructor(baseURL) {
        this.baseURL = baseURL;
    }

    login(email, password) {
        return post({
            url: '/api/auth',
            data: { email, password },
            Type: User.fromJSON,
        }).then( data => console.log("login response", data));
    }

    logout() {
        return del({
            url: '/api/auth',
        });
    }

    getUser(userID = 'me') {
        return get({
            url: '/api/users/' + userID,
            Type: User.fromJSON,
        });
    }

    getTweets() {
      return get({
          url: '/api/tweets/',
      });
    }
}
