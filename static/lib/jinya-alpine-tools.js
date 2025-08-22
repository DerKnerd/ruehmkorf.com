import Alpine from './alpine.js';
import PineconeRouter from './pinecone-router.js';
import { head, httpDelete } from './jinya-http.js';

let scriptBasePath = '/static/js/';

async function checkLogin() {
  try {
    await head('/api/authentication/login');
    return true;
  } catch (e) {
    return false;
  }
}

export async function needsLogin(context) {
  if (await checkLogin()) {
    Alpine.store('authentication').loggedIn = true;
    return null;
  } else {
    Alpine.store('authentication').loggedIn = false;
    return context.redirect('/login');
  }
}

export async function needsLogout(context) {
  if (await checkLogin()) {
    Alpine.store('authentication').loggedIn = true;
    return context.redirect('/');
  } else {
    Alpine.store('authentication').loggedIn = false;
    return null;
  }
}

export async function fetchScript({ route }) {
  const [, , area] = route.split('/');
  if (area) {
    await import(`${scriptBasePath}${area}.js`);
    Alpine.store('navigation').navigate({
      area,
    });
  }
}

function setupRouting(baseScriptPath, routerBasePath = '') {
  scriptBasePath = baseScriptPath;

  document.addEventListener('alpine:init', () => {
    window.PineconeRouter.settings.basePath = routerBasePath;
    window.PineconeRouter.settings.templateTargetId = 'app';
    window.PineconeRouter.settings.includeQuery = false;
  });
}

async function setupAlpine(alpine, defaultArea, defaultPage) {
  Alpine.directive('active-route', (el, { expression, modifiers }, { Alpine, effect }) => {
    effect(() => {
      const { page, area } = Alpine.store('navigation');
      if ((modifiers.includes('area') && area === expression) || (!modifiers.includes('area') && page === expression)) {
        el.classList.add('is--active');
      } else {
        el.classList.remove('is--active');
      }
    });
  });
  Alpine.directive('active', (el, { expression }, { Alpine, effect }) => {
    effect(() => {
      if (Alpine.evaluate(el, expression)) {
        el.classList.add('is--active');
      } else {
        el.classList.remove('is--active');
      }
    });
  });

  Alpine.store('loaded', false);
  Alpine.store('authentication', {
    needsLogin,
    needsLogout,
    login() {
      this.loggedIn = true;
      history.replaceState(null, null, location.href.split('?')[0]);
    },
    async logout() {
      try {
        await httpDelete('/api/authentication/login');
      } catch (e) {}
      window.PineconeRouter.context.navigate('/login');
      this.loggedIn = false;
      this.roles = [];
    },
  });
  Alpine.store('navigation', {
    fetchScript,
    area: defaultArea,
    navigate({ area }) {
      this.area = area;
    },
  });
}

export async function setup({ defaultArea, baseScriptPath, routerBasePath = '', afterSetup = () => {} }) {
  window.Alpine = Alpine;

  Alpine.plugin(PineconeRouter);
  await setupAlpine(Alpine, defaultArea);

  setupRouting(baseScriptPath, routerBasePath);

  await afterSetup();
  Alpine.store('authentication').loggedIn = !!(await checkLogin());

  Alpine.start();

  Alpine.store('loaded', true);
}
