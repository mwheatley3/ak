import { get, del, post } from './util/fetch';

export default class API {
    constructor(baseURL) {
        this.baseURL = baseURL;
    }

    login(email, password) {
        return post({
            url: '/api/auth',
            // data: { email, password },
            // Type: User.fromJSON,
        }).then(val => console.log(val + "something"));
    }

    logout() {
        return del({
            url: '/api/auth',
        });
    }

    getUser(userID = 'me') {
        return get({
            url: '/api/users/' + userID,
        });
    }

    getTweets() {
      return get({
          url: '/api/tweets/',
      });
    }
}
