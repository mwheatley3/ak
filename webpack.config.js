var path = require('path');
var autoprefixer = require('autoprefixer');
var postcssnested = require('postcss-nested');
var base = __dirname;

module.exports = {
  __publicDir: path.join(base, 'public'),

  output: {
      path: path.join(base, 'public'),
      publicPath: '/public/',
      filename: '[name].js',
  },

  entry: './index.js',

  postcss: function() {
        return [autoprefixer, postcssnested];
    },

  module: {
    preLoaders: [
      { test: /\.js?$/, loader: 'eslint-loader', exclude: /node_modules/ },
    ],
    loaders: [
      { test: /\.js?$/, exclude: /node_modules/, loader: 'babel' },
      { test: /\.css$/, loader: 'style/useable!css?importLoaders=1!postcss' },
      { test: /\.json$/, loader: 'json-loader' },
      { test: /\.woff($|\?)/, loader: 'url-loader?limit=10000' },
      { test: /\.woff2($|\?)/, loader: 'url-loader?limit=10000' },
      { test: /\.otf($|\?)/, loader: 'url-loader?limit=10000' },
      { test: /\.ttf($|\?)/, loader: 'url-loader?limit=10000' },
      { test: /\.eot($|\?)/, loader: 'url-loader?limit=10000' },
      { test: /\.svg($|\?)/, loader: 'url-loader?limit=10000' },
      { test: /\.png$/, loader: 'url-loader?limit=10000' },
      { test: /\.jpg$/, loader: 'url-loader?limit=10000' },
      { test: /\.gif$/, loader: 'url-loader?limit=10000' },
    ],
  },

  resolve: {
    extensions: ['', '.webpack.js', '.web.js', '.js'],
    alias: {
        client: path.join(base, 'client'),
    },
  },
};
