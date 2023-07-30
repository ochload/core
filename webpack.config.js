const path = require('path');

module.exports = {
    entry: './static/js/App.js',
    mode: 'development',
    output: {
        path: path.resolve(__dirname, 'static/pages/assets'),
        filename: 'app.bundle.js',
    },
    module: {
        rules: [
            {
                test: /\.js?$/,
                exclude: /node_module/,
                use: 'babel-loader'
            },
            {
                test: /\.css$/i,
                use: ["style-loader", "css-loader", "postcss-loader"],
            }
        ]
    }
};
