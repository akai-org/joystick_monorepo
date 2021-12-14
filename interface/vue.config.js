module.exports = {
  css: {
    loaderOptions: {
      sass: {
        prependData: `@import "@/app.scss";`
      }
    }
  }
}
