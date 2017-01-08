import hoist from 'hoist-non-react-statics';

export default function(...styles) {
    return function(Comp) {
        class Style extends Comp {
            static displayName = 'Style(' + (Comp.displayName || Comp.name) + ')';

            constructor(props, context) {
                for (const i in styles) {
                    styles[i].use();
                }

                super(props, context);
            }
        }

        hoist(Style, Comp);

        return Style;
    };
}
