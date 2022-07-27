import {unmarkSubMenuLinks} from "./navigation.js";
import {compileTemplate} from "./template.js";
import {unmarkListLinks} from "./lists.js";
import {toggleTab} from "./tabs.js";
import {alert, confirm} from "./dialogs.js";

export async function selectDownload(slug) {
    if (!slug) {
        return;
    }

    unmarkListLinks();
    document.querySelector(`[data-download-slug="${slug}"]`).classList.add('cosmo-list__item--active');

    const download = await (await fetch(`/admin/download?slug=${slug}`)).json();
    await compileTemplate('downloadDetails.js', document.getElementById('downloadContent'), download);

    document.querySelector('[data-action=german]').addEventListener('click', () => toggleTab('download', 'german'));
    document.querySelector('[data-action=english]').addEventListener('click', () => toggleTab('download', 'english'));
    document.querySelector('[data-action=preview]').addEventListener('click', () => toggleTab('download', 'preview'));

    document.querySelector('[data-action=deleteDownload]').addEventListener('click', async () => {
        const result = await confirm('Nachricht löschen', `Soll der Download ${download.slug} wirklich gelöscht werden?`, 'Download löschen', 'Download behalten');
        if (result) {
            await fetch(`/admin/download?slug=${slug}`, {method: 'DELETE'});
            await init();
        }
    });

    document.getElementById('saveDownload').addEventListener('click', async () => {
        const descriptionDe = document.getElementById('descriptionDe').value;
        const descriptionEn = document.getElementById('descriptionEn').value;

        const response = await fetch(`/admin/download?slug=${slug}`, {
            method: 'PUT',
            body: JSON.stringify({
                ...download,
                descriptionDe,
                descriptionEn,
            }),
            headers: {
                'Content-Type': 'application/json',
            },
        });
        if (response.status !== 204) {
            await alert('Speichern fehlgeschlagen', 'Beim Speichern ist ein unbekannter Fehler aufgetreten.');
        }
    });
    document.querySelector('[data-action=editDownload]').addEventListener('click', () => showEditModal(download));

    document.querySelector('[data-action=uploadFile]').addEventListener('click', async () => {
        const fileInput = document.createElement('input');
        fileInput.type = 'file';
        fileInput.addEventListener('change', async (e) => {
            if (fileInput.files.length > 0) {
                document.getElementById('uploadStatus').classList.remove('rc-hidden');
                const file = fileInput.files.item(0);
                const progressElement = document.getElementById('uploadStatusProgress');
                progressElement.setAttribute('max', file.size.toString());
                progressElement.setAttribute('value', '0');
                const uploadStatusElement = document.getElementById('uploadStatusText');
                const length = 10 * 1024 * 1024;
                let body;
                for (let offset = 0; offset <= file.size; offset += length) {
                    const percent = (offset / file.size) * 100;
                    uploadStatusElement.textContent = `${percent.toFixed(2)} % hochgeladen`;
                    progressElement.setAttribute('value', offset.toString());
                    body = file.slice(offset, length + offset);
                    const result = await fetch(`/admin/download/file/chunk?slug=${slug}&index=${offset.toString().padStart(10, '0')}`, {
                        body,
                        method: 'POST',
                    });
                    if (result.status !== 204) {
                        await alert('Upload fehlgeschlagen', 'Beim Upload ist ein unbekannter Fehler aufgetreten.');
                        break;
                    }
                }
                uploadStatusElement.textContent = 'Upload wird abgeschlossen';
                const result = await fetch(`/admin/download/file/finish?slug=${slug}`, {method: 'POST'});
                if (result.status !== 204) {
                    await alert('Upload fehlgeschlagen', 'Beim Upload ist ein unbekannter Fehler aufgetreten.');
                }
                document.getElementById('uploadStatus').classList.add('rc-hidden');
            }
        });
        fileInput.click();
    });
}

async function showAddModal() {
    const container = document.createElement('div');
    await compileTemplate('downloadAdd.js', container);
    document.body.appendChild(container);

    container.querySelector('form').addEventListener('submit', async (e) => {
        e.preventDefault();
        const previewImage = document.getElementById('previewImage');
        const nameDe = document.getElementById('nameDe').value;
        const nameEn = document.getElementById('nameEn').value;
        const slug = document.getElementById('slug').value;
        const selfDestructDays = parseInt(document.getElementById('selfDestructDays').value);
        const date = new Date(Date.parse(document.getElementById('date').value));
        const publicChecked = document.getElementById('public').checked;
        const result = await fetch(`/admin/download`, {
            body: JSON.stringify({
                slug,
                nameDe,
                nameEn,
                public: publicChecked,
                date: date.toISOString(),
                selfDestructDays,
            }),
            method: 'POST',
        });
        if (result.status === 409) {
            await alert('Speichern fehlgeschlagen', `Ein Download mit dem Slug ${slug} existiert bereits`);
        } else if (result.status !== 201) {
            await alert('Speichern fehlgeschlagen', 'Beim Speichern ist ein unbekannter Fehler aufgetreten');
        } else {
            if (previewImage.files.length > 0) {
                const fileUploadResult = await fetch(`/admin/download/preview?slug=${slug}`, {
                    method: 'POST',
                    body: previewImage.files.item(0),
                });
                if (fileUploadResult.status !== 204) {
                    await alert('Speichern fehlgeschlagen', 'Beim Speichern des Vorschau Bildes ist ein unbekannter Fehler aufgetreten');
                    return
                }
            }
            document.body.removeChild(container);
            await init();
            await selectDownload(slug);
        }
    });
    container.querySelector('[data-action=cancelAdd]').addEventListener('click', () => document.body.removeChild(container));
}

async function showEditModal(download) {
    const container = document.createElement('div');
    await compileTemplate('downloadEdit.js', container, download);
    document.body.appendChild(container);

    const date = new Date(Date.parse(download.date));
    container.querySelector('#date').value = `${date.getFullYear().toString().padStart(2, '0')}-${(date.getMonth() + 1).toString().padStart(2, '0')}-${date.getDate().toString().padStart(2, '0')}`;

    container.querySelector('form').addEventListener('submit', async (e) => {
        e.preventDefault();
        const previewImage = document.getElementById('previewImage');
        const nameDe = document.getElementById('nameDe').value;
        const nameEn = document.getElementById('nameEn').value;
        const selfDestructDays = parseInt(document.getElementById('selfDestructDays').value);
        const date = new Date(Date.parse(document.getElementById('date').value));
        const publicChecked = document.getElementById('public').checked;
        const result = await fetch(`/admin/download?slug=${download.slug}`, {
            body: JSON.stringify({
                ...download,
                nameDe,
                nameEn,
                public: publicChecked,
                date: date.toISOString(),
                selfDestructDays,
            }),
            method: 'PUT',
        });
        if (result.status === 409) {
            await alert('Speichern fehlgeschlagen', `Ein Download mit dem Slug ${download.slug} existiert bereits`);
        } else if (result.status !== 204) {
            await alert('Speichern fehlgeschlagen', 'Beim Speichern ist ein unbekannter Fehler aufgetreten');
        } else {
            if (previewImage.files.length > 0) {
                const fileUploadResult = await fetch(`/admin/download/preview?slug=${download.slug}`, {
                    method: 'POST',
                    body: previewImage.files.item(0),
                });
                if (fileUploadResult.status !== 204) {
                    await alert('Speichern fehlgeschlagen', 'Beim Speichern des Vorschau Bildes ist ein unbekannter Fehler aufgetreten');
                    return
                }
            }
            document.body.removeChild(container);
            await init();
            await selectDownload(download.slug);
        }
    });
    container.querySelector('[data-action=cancelEdit]').addEventListener('click', () => document.body.removeChild(container));
}

export async function init() {
    unmarkSubMenuLinks();
    document.querySelector('[data-sublink=downloads]').classList.add('cosmo-menu-bar__sub-item--active');
    const downloads = await (await fetch('/admin/download')).json();
    await compileTemplate('downloadsList.js', document.getElementById('rcContent'), {downloads});

    await selectDownload(downloads[0]?.slug);
    document.querySelectorAll('[data-action=changeDownload]').forEach(link => link.addEventListener('click', async (e) => {
        await selectDownload(e.target.getAttribute('data-download-slug'));
    }));

    document.querySelector('[data-action=addDownload]').addEventListener('click', showAddModal);
}