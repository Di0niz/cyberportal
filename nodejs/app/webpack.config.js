module.exports = (env) => ({
  entry: {
    bundle: path.resolve(__dirname, 'src/index.js')
  }, 
  
  output: {
    path: path.resolve(__dirname, 'build'),
    filename: 'bundle.js'
  },
  devtool: 'source-map',
  module: {
    loaders: [ 
      {
        test: /\.jsx?$/,
        loader: 'babel-loader',
        exclude: /node_modules/
      }
    ]
  }, 
  
  plugins: [
    new CopyWebpackPlugin([ { context: 'public', from: '**/*' } ])
  ]

  
});