const js = require('@eslint/js');
const prettier = require('eslint-config-prettier');
const globals = require('globals');

module.exports = [
  {
    ...js.configs.recommended,
    ...prettier,
    ignore: ['/vendor', '/themes', '/static/lib'],
    languageOptions: {
      ecmaVersion: 'latest',
      sourceType: 'module',
      globals: {
        ...globals.browser,
        Jodit: 'readonly',
      },
    },
  },
];
