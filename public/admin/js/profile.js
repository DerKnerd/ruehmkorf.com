import {unmarkSubMenuLinks} from "./navigation.js";
import {compileTemplate} from "./template.js";
import {unmarkListLinks} from "./lists.js";
import {alert, confirm} from "./dialogs.js";

async function selectProfile(id) {
    if (!id) {
        return;
    }

    unmarkListLinks();

    const element = document.querySelector(`[data-profile-id="${id}"]`);
    element.classList.add('is--active');

    const profile = await (await fetch(`/admin/profile?id=${id}`)).json();

    await compileTemplate('profileDetails.js', document.getElementById('profileContent'), profile);

    document.querySelector('[data-action=deleteProfile]').addEventListener('click', async () => {
        const result = await confirm('Nachricht löschen', `Soll das Profil ${profile.name} wirklich gelöscht werden?`, 'Profil löschen', 'Profil behalten');
        if (result) {
            await fetch(`/admin/profile?id=${id}`, {method: 'DELETE'});
            await init();
        }
    });

    document.querySelector('[data-action=editProfile]').addEventListener('click', () => showEditModal(profile));
}

async function showEditModal(profile) {
    const container = document.createElement('div');
    await compileTemplate('profileEdit.js', container, profile);
    document.body.appendChild(container);

    container.querySelector('form').addEventListener('submit', async (e) => {
        e.preventDefault();
        const name = document.getElementById('name').value;
        const url = document.getElementById('url').value;
        const active = document.getElementById('active').checked;
        const headerImage = document.getElementById('headerImage');
        const icon = document.getElementById('icon');
        const result = await fetch(`/admin/profile?id=${profile.id}`, {
            body: JSON.stringify({
                name,
                url,
                active,
            }),
            method: 'PUT',
        });
        if (result.status !== 204) {
            await alert('Speichern fehlgeschlagen', 'Beim Speichern ist ein unbekannter Fehler aufgetreten');
        } else {
            if (headerImage.files.length > 0) {
                const fileUploadResult = await fetch(`/admin/profile/header?id=${profile.id}`, {
                    method: 'POST',
                    body: headerImage.files.item(0)
                });
                if (fileUploadResult.status !== 204) {
                    await alert('Speichern fehlgeschlagen', 'Beim Speichern des Header Bildes ist ein unbekannter Fehler aufgetreten');
                    return
                }
            }
            if (icon.files.length > 0) {
                const fileUploadResult = await fetch(`/admin/profile/icon?id=${profile.id}`, {
                    method: 'POST',
                    body: icon.files.item(0)
                });
                if (fileUploadResult.status !== 204) {
                    await alert('Speichern fehlgeschlagen', 'Beim Speichern des Icons ist ein unbekannter Fehler aufgetreten');
                    return
                }
            }

            document.body.removeChild(container);
            await init();
            await selectProfile(profile.id);
        }
    });
    container.querySelector('[data-action=cancelAdd]').addEventListener('click', () => document.body.removeChild(container));
}

async function showAddModal() {
    const container = document.createElement('div');
    await compileTemplate('profileAdd.js', container);
    document.body.appendChild(container);

    container.querySelector('form').addEventListener('submit', async (e) => {
        e.preventDefault();
        const headerImage = document.getElementById('headerImage');
        const icon = document.getElementById('icon');
        const name = document.getElementById('name').value;
        const url = document.getElementById('url').value;
        const active = document.getElementById('active').checked;
        const result = await fetch(`/admin/profile`, {
            body: JSON.stringify({
                name,
                url,
                active,
            }),
            method: 'POST',
        });
        if (result.status !== 201) {
            await alert('Speichern fehlgeschlagen', 'Beim Speichern ist ein unbekannter Fehler aufgetreten');
        } else {
            const id = await result.text();
            if (headerImage.files.length > 0) {
                const fileUploadResult = await fetch(`/admin/profile/header?id=${id}`, {
                    method: 'POST',
                    body: headerImage.files.item(0)
                });
                if (fileUploadResult.status !== 204) {
                    await alert('Speichern fehlgeschlagen', 'Beim Speichern des Header Bildes ist ein unbekannter Fehler aufgetreten');
                    return
                }
            }
            if (icon.files.length > 0) {
                const fileUploadResult = await fetch(`/admin/profile/icon?id=${id}`, {
                    method: 'POST',
                    body: icon.files.item(0)
                });
                if (fileUploadResult.status !== 204) {
                    await alert('Speichern fehlgeschlagen', 'Beim Speichern des Icons ist ein unbekannter Fehler aufgetreten');
                    return
                }
            }

            document.body.removeChild(container);
            await init();
            await selectProfile(id);
        }
    });
    container.querySelector('[data-action=cancelAdd]').addEventListener('click', () => document.body.removeChild(container));
}

export async function init() {
    unmarkSubMenuLinks();
    document.querySelector('[data-sublink=profile]').classList.add('is--active');
    const profiles = await (await fetch('/admin/profile')).json();
    await compileTemplate('profilesList.js', document.getElementById('rcContent'), {profiles});

    await selectProfile(profiles[0]?.id);
    document.querySelectorAll('[data-action=changeProfile]').forEach(link => link.addEventListener('click', async (e) => {
        await selectProfile(e.target.getAttribute('data-profile-id'));
    }));

    document.querySelector('[data-action=addProfile]').addEventListener('click', showAddModal);
}
