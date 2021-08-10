import {hideSubmenus, unmarkMainMenuLinks, unmarkSubMenuLinks} from "./navigation.js";
import {init} from "./content.js";

document.querySelectorAll('a.cosmo-menu-bar__main-item').forEach(link => link.addEventListener('click', async (e) => {
    const target = e.target.getAttribute('data-target');
    hideSubmenus();
    document.querySelector(`[data-submenu=${target}]`).classList.remove('rc-hidden');
    const content = await import((`./${target}.js`));
    if (content) {
        await content.init();
    }

    unmarkMainMenuLinks();
    e.target.classList.add('cosmo-menu-bar__main-item--active');

    unmarkSubMenuLinks();
    const defaultSubLink = e.target.getAttribute('data-default-sublink');
    document.querySelector(`[data-sublink=${defaultSubLink}]`).classList.add('cosmo-menu-bar__sub-item--active');
}));

await init();
