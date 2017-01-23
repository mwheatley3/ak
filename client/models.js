export class User {
    constructor(id, email) {
        this.id = id;
        this.email = email;
    }

    static fromJSON(obj) {
        return new User(obj.id, obj.email);
    }
}
