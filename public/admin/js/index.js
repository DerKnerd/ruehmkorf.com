import {hideSubmenus, unmarkMainMenuLinks, unmarkSubMenuLinks} from "./navigation.js";
import {initContent} from "./content.js";

document.querySelectorAll('a.cosmo-menu-bar__main-item').forEach(link => link.addEventListener('click', (e) => {
    const target = e.currentTarget.getAttribute('data-target');
    hideSubmenus();
    document.querySelector(`[data-submenu=${target}]`).classList.remove('rc-hidden');

    unmarkMainMenuLinks();
    e.currentTarget.classList.add('cosmo-menu-bar__main-item--active');

    unmarkSubMenuLinks();
    const defaultSubLink = e.currentTarget.getAttribute('data-default-sublink');
    document.querySelector(`[data-sublink=${defaultSubLink}]`).classList.add('cosmo-menu-bar__sub-item--active');
}));

document.querySelector('[data-target=content]').addEventListener('click', async () => {
    const content = await import('./content.js');
    await content.navigateToNews()
});

await initContent();
