module.exports = {
  // ...
  module: {
    rules: [
      {
        test: /\.(ts|tsx)$/,
        exclude: /node_modules/,
        use: {
          loader: "babel-loader",
        },
      },
    ],
  },
  // ...
  resolve: {
    extensions: [".js", ".jsx", ".ts", ".tsx"],
  },
};
