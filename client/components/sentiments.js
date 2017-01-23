import React, { Component } from 'react';
import Form, { TextInput } from './common/form';

export default class Sentiments extends Component {
    static propTypes = {
    };

    onSubmit = e => {
        if (e) {
            e.preventDefault();
        }
    }

    render() {
        return (
            <div>Sentiments
              <Form onSubmit={ this.onSubmit }>
                twitter handle: <TextInput />
              </Form>
            </div>
        );
    }
}
