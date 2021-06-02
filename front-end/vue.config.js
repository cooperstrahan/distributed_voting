module.exports = {
    runtimeCompiler: true,
    devServer: {
        proxy: 'http://localhost:8040/register'
    },
    // headers: {
    //     "Access-Control-Allow-Origin": "*",
    // }
}