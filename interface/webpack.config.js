/* chcę dodać do package.json linijkę:
"build": "webpack --config webpack.config.js"
ale nie wiem jak, żeby nie popsuć tego co jest:
"build": "vue-cli-service build" */
const path = require('path')
module.exports = {
  mode: 'development',
  entry: './src/app.scss',
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'index.min.js'
  },
  module: {
    rules: [
      {
        test: /\.scss$/,
        exclude: /node_modules/,
        use: [
          {
            loader: 'file-loader',
            options: { outputPath: '/src', name: 'app.css' }
          },
          'sass-loader'
        ]
      }
    ]
  }
}
