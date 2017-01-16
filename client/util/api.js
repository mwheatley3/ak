import { get, del, post } from './fetch';

export default class API {
    constructor(baseURL) {
        this.baseURL = baseURL;
    }

    login(email, password) {
      console.log("api login");
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
}
