export function hideSubmenus() {
    const subMenus = document.querySelectorAll('[data-submenu]');
    subMenus.forEach(menu => menu.classList.add('rc-hidden'));
}

export function unmarkMainMenuLinks() {
    const links = document.querySelectorAll('.cosmo-menu-bar__main-item--active');
    links.forEach(link => link.classList.remove('cosmo-menu-bar__main-item--active'));
}

export function unmarkSubMenuLinks() {
    const links = document.querySelectorAll('.cosmo-menu-bar__sub-item--active');
    links.forEach(link => link.classList.remove('cosmo-menu-bar__sub-item--active'));
}

export async function navigatePage(section, sublink) {
    window.location.hash = `${section}/${sublink}`;
}
