import Promise from 'bluebird';

export const EMPTY = 'empty';
export const LOADING = 'loading';
export const SUCCESS = 'success';
export const ERROR = 'error';

export default class Value {
     state = EMPTY;
     value = void 0;
     error = void 0;

    constructor({ value, error } = {}) {
        if (value) {
            this.setValue(value);
        } else if (error) {
            this.setError(error);
        }
    }

     get empty() {
        return this.state === EMPTY;
    }

     get loaded() {
        return this.state !== EMPTY && this.state !== LOADING;
    }

     get loading() {
        return this.state === LOADING;
    }

     setValue(v) {
        this.value = v;
        this.error = void 0;
        this.state = SUCCESS;
    }

     setError(err) {
        this.value = void 0;
        this.error = err;
        this.state = ERROR;
    }

     trackPromise(p) {
        this.state = LOADING;

        p.then(
            val => this.setValue(val),
            err => this.setError(err)
        );
    }

    onValue(fn) {
        const dispose = () => {
            if (this.state === SUCCESS) {
                fn(this.value);
                dispose();
                return;
            }
        };

        return this;
    }

    asPromise() {
        return new Promise((resolve, reject) => {
            const dispose = () => {
                if (this.state === SUCCESS) {
                    resolve(this.value);
                    dispose();
                    return;
                }

                if (this.state === ERROR) {
                    reject(this.error);
                    dispose();
                    return;
                }
            };
        });
    }

    static trackPromise(p) {
        const v = new Value();
        v.trackPromise(p);

        return v;
    }
}
