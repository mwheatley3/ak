import React, { Component } from 'react';

export default class Keri extends Component {
    static propTypes = {
    };

    onSubmit = e => {
        if (e) {
            e.preventDefault();
        }
    }

    render() {
        return (
            <div>Keri's Desk
              <img src="http://s3.amazonaws.com/workingwheatleys/public/keri_desk.jpg"/>
            </div>
        );
    }
}
