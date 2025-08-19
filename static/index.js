import { setup } from './lib/jinya-alpine-tools.js';

document.addEventListener('DOMContentLoaded', async () => {
  await setup({
    defaultArea: 'server',
    defaultPage: 'config',
    baseScriptPath: '/static/js/',
    routerBasePath: '/',
    openIdClientId: window.relayConfig.openIdClientId,
    openIdUrl: window.relayConfig.openIdUrl,
    openIdCallbackUrl: window.relayConfig.openIdCallbackUrl,
  });
});
