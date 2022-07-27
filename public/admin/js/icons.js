import {compileTemplate} from "./template.js";
import {unmarkSubMenuLinks} from "./navigation.js";

export async function init() {
    unmarkSubMenuLinks();
    document.querySelector('[data-sublink=icons]').classList.add('cosmo-menu-bar__sub-item--active');

    await compileTemplate('icons.js', document.getElementById('rcContent'));

    const touchiconInput = document.querySelector('#touchicon');
    touchiconInput.addEventListener('change', (e) => {
        const target = e.target;
        document.querySelector('[for=touchicon].cosmo-picker__name').textContent = target.files.item(0).name;
    });

    const logoInput = document.querySelector('#logo');
    logoInput.addEventListener('change', (e) => {
        const target = e.target;
        document.querySelector('[for=logo].cosmo-picker__name').textContent = target.files.item(0).name;
    });

    const faviconInput = document.querySelector('#favicon');
    faviconInput.addEventListener('change', (e) => {
        const target = e.target;
        document.querySelector('[for=favicon].cosmo-picker__name').textContent = target.files.item(0).name;
    });

    document.querySelector('form').addEventListener('submit', async (e) => {
        e.preventDefault();
        if (faviconInput.files.length > 0) {
            await fetch('/admin/settings/favicon', {method: 'POST', body: faviconInput.files.item(0)});
        }
        if (touchiconInput.files.length > 0) {
            await fetch('/admin/settings/touchicon', {method: 'POST', body: touchiconInput.files.item(0)});
        }
        if (logoInput.files.length > 0) {
            await fetch('/admin/settings/logo', {method: 'POST', body: logoInput.files.item(0)});
        }
        await compileTemplate('icons.js', document.getElementById('rcContent'));
    });
}