import React, { Component } from 'react';
import Form, { TextInput } from './common/form';

export default class Coffee extends Component {
    static propTypes = {
    };

    render() {
        return (
            <div>Coffee
              <Form>
                twitter handle: <TextInput />
              </Form>
            </div>
        );
    }
}
