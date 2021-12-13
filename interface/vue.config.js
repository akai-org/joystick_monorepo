module.exports = {
  css: {
    loaderOptions: {
      sass: {
        prependData: '@import "./src/app.scss";'
      }
    }
  }
}
