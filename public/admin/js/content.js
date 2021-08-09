import {unmarkSubMenuLinks} from "./navigation.js";

export async function navigateToNews() {
    const news = await import('./news.js');
    await news.initNews();
}

export async function init() {
    document.querySelector('[data-sublink=news]').addEventListener('click', async () => await navigateToNews());
    document.querySelector('[data-sublink=downloads]').addEventListener('click', (e) => {
        unmarkSubMenuLinks();
        e.currentTarget.classList.add('cosmo-menu-bar__sub-item--active');
    });

    await navigateToNews();
}
