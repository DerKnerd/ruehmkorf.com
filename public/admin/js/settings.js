import {navigatePage} from "./navigation.js";

export async function init() {
    document.querySelectorAll('[data-sublink]').forEach(link => link.addEventListener('click', async (e) => {
        e.preventDefault();
        await navigatePage('settings', e.target.getAttribute('data-sublink'));
    }));

    await navigatePage('settings', 'icons');
}
