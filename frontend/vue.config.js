 // '^/api': {
        //   target: 'http://83.69.248.180:9050/api',
        //   changeOrigin: true,
        // }
      //}
const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
    devServer: {
      webSocketServer: false,
      //proxy: 'http://83.69.249.5:8000/'
      proxy: 'https://genmilk-krasnodar.ru:443'
      
      // {
      //   '^/api': {
      //     target: 'http://localhost:8080/',
      //     changeOrigin: true,
      //   }
      // }
      // headers: {"Access-Control-Allow-Origin": "*"}
  }
})
