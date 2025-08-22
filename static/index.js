import { setup } from './lib/jinya-alpine-tools.js';

document.addEventListener('DOMContentLoaded', async () => {
  await setup({
    defaultArea: 'profiles',
    baseScriptPath: '/static/js/',
    routerBasePath: '/admin',
  });
});
