import {hideSubmenus, unmarkMainMenuLinks, unmarkSubMenuLinks} from "./navigation.js";

document.querySelectorAll('a.cosmo-menu-bar__main-item').forEach(link => link.addEventListener('click', async (e) => {
    e.preventDefault();
    const target = e.target.getAttribute('data-target');
    hideSubmenus();
    document.querySelector(`[data-submenu=${target}]`).classList.remove('rc-hidden');

    unmarkMainMenuLinks();
    e.target.classList.add('cosmo-menu-bar__main-item--active');

    unmarkSubMenuLinks();
    const defaultSubLink = e.target.getAttribute('data-default-sublink');
    document.querySelector(`[data-sublink=${defaultSubLink}]`).classList.add('cosmo-menu-bar__sub-item--active');

    const content = await import((`./${target}.js`));
    if (content) {
        await content.init();
    }
}));

async function navigateHash(hash) {
    const hashParts = hash.split('/');
    if (hashParts.length === 2) {
        const base = hashParts[0];
        const page = hashParts[1];

        const content = await import((`./${page}.js`));
        if (content) {
            await content.init();
        }
    }

    return hashParts;
}

window.addEventListener('hashchange', async (e) => {
    const urlParts = e.newURL.split('#');
    if (urlParts.length === 2) {
        await navigateHash(urlParts[1]);
    }
});

const parts = await navigateHash(window.location.hash.replace('#', ''));
if (parts.length === 2) {
    hideSubmenus();
    document.querySelector(`[data-submenu=${parts[0]}]`).classList.remove('rc-hidden');

    unmarkMainMenuLinks();
    document.querySelector(`[data-target=${parts[0]}]`).classList.add('cosmo-menu-bar__main-item--active');

    unmarkSubMenuLinks();
    document.querySelector(`[data-sublink=${parts[1]}]`).classList.add('cosmo-menu-bar__sub-item--active');
} else {
    hideSubmenus();
    document.querySelector('[data-submenu=content]').classList.remove('rc-hidden');

    unmarkMainMenuLinks();
    document.querySelector('[data-target=content]').classList.add('cosmo-menu-bar__main-item--active');

    unmarkSubMenuLinks();
    document.querySelector('[data-sublink=news]').classList.add('cosmo-menu-bar__sub-item--active');
    window.location.hash = 'content/news';
}