import React, { Component, PropTypes } from 'react';

const { any, element } = PropTypes;

export function injector(ctxt) {
    const ctxtTypes = {};

    for (let k in ctxt) {
        ctxtTypes[k] = any.isRequired;
    }

    return class ContextInjector extends Component {
        static propTypes = {
            children: element.isRequired,
        };

        static childContextTypes = ctxtTypes;

        getChildContext() {
            return ctxt;
        }

        render() {
            return this.props.children;
        }
    };
}

// provider creates a function that can decorate react components to provide context
export function provider(...keys) {
  // transform inputs into an object that represents the context
    const ctxtTypes = keys.reduce((obj, k) => {
        obj[k] = any.isRequired;
        return obj;
    }, {});

    return function(Comp) {
        class ContextProvider extends Component {
            static displayName = 'ContextProvider(' + (Comp.displayName || Comp.name) + ')';
            // use the default naming convention to place context on component
            static contextTypes = ctxtTypes;

            // combine context and props into props so that the context can be accessed via props
            render() {
                const props = Object.assign({}, this.props, this.context);
                return <Comp { ...props } />;
            }
        }

        // hoist(ContextProvider, Comp);

        return ContextProvider;
    };
}
