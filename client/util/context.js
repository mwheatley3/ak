import React, { Component, PropTypes } from 'react';
import hoist from 'hoist-non-react-statics';

const { element, any } = PropTypes;

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

export function provider(...keys) {
    const ctxtTypes = keys.reduce((obj, k) => {
        obj[k] = any.isRequired;
        return obj;
    }, {});

    return function(Comp) {
        class ContextProvider extends Component {
            static displayName = 'ContextProvider(' + (Comp.displayName || Comp.name) + ')';
            static contextTypes = ctxtTypes;

            render() {
                const props = Object.assign({}, this.props, this.context);
                return <Comp { ...props } />;
            }
        }

        hoist(ContextProvider, Comp);

        return ContextProvider;
    };
}
