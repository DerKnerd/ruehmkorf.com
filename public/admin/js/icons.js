import {compileTemplate} from "./template.js";
import {unmarkSubMenuLinks} from "./navigation.js";

export async function init() {
    unmarkSubMenuLinks();
    document.querySelector('[data-sublink=icons]').classList.add('is--active');

    await compileTemplate('icons.js', document.getElementById('rcContent'));

    document.querySelector('form').addEventListener('submit', async (e) => {
        e.preventDefault();

        const touchiconInput = document.querySelector('#touchIcon');
        const logoInput = document.querySelector('#logo');
        const faviconInput = document.querySelector('#favicon');

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