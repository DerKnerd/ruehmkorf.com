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
    element.classList.add('cosmo-list__item--active');

    const profile = await (await fetch(`/admin/profile?id=${id}`)).json();

    await compileTemplate('profileDetails.hbs', document.getElementById('profileContent'), profile);

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
    await compileTemplate('profileEdit.hbs', container, profile);
    document.body.appendChild(container);

    const headerImageInput = container.querySelector('#headerImage');
    headerImageInput.addEventListener('change', (e) => {
        const target = e.target;
        container.querySelector('[for=headerImage].cosmo-picker__name').textContent = target.files.item(0).name;
    });

    const iconInput = container.querySelector('#icon');
    iconInput.addEventListener('change', (e) => {
        const target = e.target;
        container.querySelector('[for=icon].cosmo-picker__name').textContent = target.files.item(0).name;
    });

    container.querySelector('form').addEventListener('submit', async (e) => {
        e.preventDefault();
        const name = document.getElementById('name').value;
        const url = document.getElementById('url').value;
        const active = document.getElementById('active').checked;
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
            if (headerImageInput.files.length > 0) {
                const fileUploadResult = await fetch(`/admin/profile/header?id=${profile.id}`, {
                    method: 'POST',
                    body: headerImageInput.files.item(0)
                });
                if (fileUploadResult.status !== 204) {
                    await alert('Speichern fehlgeschlagen', 'Beim Speichern des Header Bildes ist ein unbekannter Fehler aufgetreten');
                    return
                }
            }
            if (iconInput.files.length > 0) {
                const fileUploadResult = await fetch(`/admin/profile/icon?id=${profile.id}`, {
                    method: 'POST',
                    body: iconInput.files.item(0)
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
    await compileTemplate('profileAdd.hbs', container);
    document.body.appendChild(container);

    const headerImageInput = container.querySelector('#headerImage');
    headerImageInput.addEventListener('change', (e) => {
        const target = e.target;
        container.querySelector('[for=headerImage].cosmo-picker__name').textContent = target.files.item(0).name;
    });

    const iconInput = container.querySelector('#icon');
    iconInput.addEventListener('change', (e) => {
        const target = e.target;
        container.querySelector('[for=icon].cosmo-picker__name').textContent = target.files.item(0).name;
    });

    container.querySelector('form').addEventListener('submit', async (e) => {
        e.preventDefault();
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
            if (headerImageInput.files.length > 0) {
                const fileUploadResult = await fetch(`/admin/profile/header?id=${id}`, {
                    method: 'POST',
                    body: headerImageInput.files.item(0)
                });
                if (fileUploadResult.status !== 204) {
                    await alert('Speichern fehlgeschlagen', 'Beim Speichern des Header Bildes ist ein unbekannter Fehler aufgetreten');
                    return
                }
            }
            if (iconInput.files.length > 0) {
                const fileUploadResult = await fetch(`/admin/profile/icon?id=${id}`, {
                    method: 'POST',
                    body: iconInput.files.item(0)
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
    document.querySelector('[data-sublink=profile]').classList.add('cosmo-menu-bar__sub-item--active');
    const profiles = await (await fetch('/admin/profile')).json();
    await compileTemplate('profilesList.hbs', document.getElementById('rcContent'), {profiles});

    await selectProfile(profiles[0]?.id);
    document.querySelectorAll('[data-action=changeProfile]').forEach(link => link.addEventListener('click', async (e) => {
        await selectProfile(e.target.getAttribute('data-profile-id'));
    }));

    document.querySelector('[data-action=addProfile]').addEventListener('click', showAddModal);
}
