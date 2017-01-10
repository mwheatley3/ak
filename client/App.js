import React, { Component, PropTypes } from 'react';

const { element } = PropTypes;

export default class App extends Component {
  static propTypes = {
      children: element.isRequired,
  };
  render() {
    return (
      <div>
        {this.props.children}
      </div>
  );
  }
}
