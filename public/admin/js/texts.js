import {compileTemplate} from "./template.js";
import {alert} from "./dialogs.js";
import {toggleTab} from "./tabs.js";
import {unmarkSubMenuLinks} from "./navigation.js";

export async function init() {
    unmarkSubMenuLinks();
    document.querySelector('[data-sublink=texts]').classList.add('is--active');

    const result = await fetch('/admin/settings');
    const data = await result.json();

    await compileTemplate('texts.js', document.getElementById('rcContent'), data);

    document.querySelector('[data-action=german]').addEventListener('click', () => toggleTab('texts', 'german'));
    document.querySelector('[data-action=english]').addEventListener('click', () => toggleTab('texts', 'english'));

    document.querySelector('form').addEventListener('submit', async (e) => {
        e.preventDefault();
        const cookiesDe = document.getElementById('cookiesDe').value;
        const cookiesEn = document.getElementById('cookiesEn').value;
        const dataProtectionDe = document.getElementById('dataProtectionDe').value;
        const dataProtectionEn = document.getElementById('dataProtectionEn').value;

        const result = await fetch('/admin/settings', {
            method: 'PUT',
            body: JSON.stringify({
                cookiesDe,
                cookiesEn,
                dataProtectionDe,
                dataProtectionEn,
            }),
        });
        if (result.status !== 204) {
            await alert('Speichern fehlgeschlagen', 'Speichern der Texte ist leider fehlgeschlagen.');
        }
    });
}