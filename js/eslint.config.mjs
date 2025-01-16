import globals from "globals";
import pluginJs from "@eslint/js";
import pluginReact from "eslint-plugin-react";

export default [
  pluginJs.configs.recommended,
  pluginReact.configs.flat.recommended,
  { files: ["**/*.{js,cjs,ts,jsx,tsx}"] },
  { languageOptions: { globals: globals.browser } },
  {
    settings: {
      react: {
        version: "detect",
      },
    },
  },
  {
    "rules": {
      "react/prop-types": [
        "off"
      ],
      "react/jsx-props-no-spreading": [
        "off"
      ],
      "react/jsx-uses-react": [
        "off"
      ],
      "react/react-in-jsx-scope": [
        "off"
      ],
      "react-hooks/exhaustive-deps": [
        "off"
      ],
      "react/no-array-index-key": [
        "off"
      ],
      "no-unused-vars": "off",
      "react/no-unescaped-entities": "off",
      "react/no-children-prop": "off",
    }
  },
];