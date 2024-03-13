import {compileTemplate} from "./template.js";
import {alert} from "./dialogs.js";
import {toggleTab} from "./tabs.js";
import {unmarkSubMenuLinks} from "./navigation.js";

export async function init() {
    unmarkSubMenuLinks();
    document.querySelector('[data-sublink=bom]').classList.add('is--active');

    const result = await fetch('/admin/buchstabieromat');
    const data = await result.json();

    await compileTemplate('bom.js', document.getElementById('rcContent'), data);

    document.querySelector('[data-action=german]').addEventListener('click', () => toggleTab('bom', 'german'));
    document.querySelector('[data-action=english]').addEventListener('click', () => toggleTab('bom', 'english'));

    document.querySelector('form').addEventListener('submit', async (e) => {
        e.preventDefault();
        const descriptionDe = document.getElementById('descriptionDe').value;
        const infoTextDeDe = document.getElementById('infoTextDeDe').value;
        const infoTextEnDe = document.getElementById('infoTextEnDe').value;
        const descriptionEn = document.getElementById('descriptionEn').value;
        const infoTextDeEn = document.getElementById('infoTextDeEn').value;
        const infoTextEnEn = document.getElementById('infoTextEnEn').value;

        const result = await fetch('/admin/buchstabieromat', {
            method: 'PUT',
            body: JSON.stringify({
                descriptionDe,
                infoTextDeDe,
                infoTextEnDe,
                descriptionEn,
                infoTextDeEn,
                infoTextEnEn,
            }),
        });
        if (result.status !== 204) {
            await alert('Speichern fehlgeschlagen', 'Speichern der Texte für den Buchstabier-O-Mat ist leider fehlgeschlagen.');
        }
    });
}