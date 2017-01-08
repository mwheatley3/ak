var path = require('path');
var autoprefixer = require('autoprefixer');
var postcssnested = require('postcss-nested');

module.exports = {
  entry: './index.js',
  output: {
    path: path.join(__dirname, 'public'),
    filename: 'bundle.js',
  },
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
};
