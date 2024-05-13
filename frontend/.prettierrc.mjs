// .prettierrc.mjs
/** @type {import("prettier").Config} */

export default {
  overrides: [
    {
      files: ["*.js", "*.ts", "*.tsx", "*.jsx", "src/**/*.tsx", "src/**/*.ts"],
      options: {
        arrowParens: "avoid",
        singleQuote: true,
        bracketSpacing: true,
        endOfLine: "lf",
        semi: false,
        tabWidth: 2,
        trailingComma: "none",
      },
    },
  ],
};
