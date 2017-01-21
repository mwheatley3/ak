import React, { Component } from 'react';
import Form, { TextInput } from './common/form';

export default class Coffee extends Component {
    static propTypes = {
    };

    onSubmit = e => {
        if (e) {
            e.preventDefault();
        }
    }

    render() {
        return (
            <div>Coffee
              <Form onSubmit={ this.onSubmit }>
                twitter handle: <TextInput />
              </Form>
            </div>
        );
    }
}
