export function hideSubmenus() {
    const subMenus = document.querySelectorAll('[data-submenu]');
    subMenus.forEach(menu => menu.classList.add('rc-hidden'));
}

export function unmarkMainMenuLinks() {
    const links = document.querySelectorAll('.cosmo-menu__row.is--main .is--active');
    links.forEach(link => link.classList.remove('is--active'));
}

export function unmarkSubMenuLinks() {
    const links = document.querySelectorAll('.cosmo-menu__row.is--sub .is--active');
    links.forEach(link => link.classList.remove('is--active'));
}

export async function navigatePage(section, sublink) {
    window.location.hash = `${section}/${sublink}`;
}
