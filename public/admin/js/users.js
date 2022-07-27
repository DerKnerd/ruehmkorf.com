import {unmarkSubMenuLinks} from "./navigation.js";
import {compileTemplate} from "./template.js";
import {unmarkListLinks} from "./lists.js";
import {alert, confirm} from "./dialogs.js";

async function selectUser(id) {
    if (!id) {
        return;
    }

    unmarkListLinks();

    const element = document.querySelector(`[data-user-id="${id}"]`);
    element.classList.add('cosmo-list__item--active');

    const user = await (await fetch(`/admin/user?id=${id}`)).json();

    await compileTemplate('userDetails.js', document.getElementById('userContent'), user);
    document.querySelector('[data-action=deleteUser]').addEventListener('click', async () => {
        const result = await confirm('Benutzer löschen', `Soll der Benutzer ${user.name} wirklich gelöscht werden?`, 'Benutzer löschen', 'Benutzer behalten');
        if (result) {
            await fetch(`/admin/user?id=${id}`, {method: 'DELETE'});
            await init();
        }
    });

    document.querySelector('[data-action=editUser]').addEventListener('click', () => showEditModal(user));
}

async function showEditModal(user) {
    const container = document.createElement('div');
    await compileTemplate('userEdit.js', container, user);
    document.body.appendChild(container);

    container.querySelector('form').addEventListener('submit', async (e) => {
        e.preventDefault();
        const name = document.getElementById('name').value;
        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;
        const activated = document.getElementById('activated').checked;
        const data = {
            ...user,
            name,
            email,
            activated,
        };
        if (password !== '') {
            data.password = password;
        }
        const result = await fetch(`/admin/user?id=${user.id}`, {
            body: JSON.stringify(data),
            method: 'PUT',
        });
        if (result.status !== 204) {
            await alert('Speichern fehlgeschlagen', 'Beim Speichern ist ein unbekannter Fehler aufgetreten');
        } else {
            document.body.removeChild(container);
            await selectUser(user.id);
        }
    });
    container.querySelector('[data-action=cancelEdit]').addEventListener('click', () => document.body.removeChild(container));
}

async function showAddModal() {
    const container = document.createElement('div');
    await compileTemplate('userAdd.js', container);
    document.body.appendChild(container);

    container.querySelector('form').addEventListener('submit', async (e) => {
        e.preventDefault();
        const name = document.getElementById('name').value;
        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;
        const activated = document.getElementById('activated').checked;
        const result = await fetch(`/admin/user`, {
            body: JSON.stringify({
                password,
                name,
                email,
                activated,
            }),
            method: 'POST',
        });
        if (result.status === 409) {
            await alert('Speichern fehlgeschlagen', `Eine Nachricht mit dem Password ${password} existiert bereits`);
        } else if (result.status !== 201) {
            await alert('Speichern fehlgeschlagen', 'Beim Speichern ist ein unbekannter Fehler aufgetreten');
        } else {
            const id = await result.text();
            document.body.removeChild(container);
            await init();
            await selectUser(id);
        }
    });
    container.querySelector('[data-action=cancelAdd]').addEventListener('click', () => document.body.removeChild(container));
}

export async function init() {
    unmarkSubMenuLinks();
    document.querySelector('[data-sublink=users]').classList.add('cosmo-menu-bar__sub-item--active');
    const users = await (await fetch('/admin/user')).json();
    await compileTemplate('usersList.js', document.getElementById('rcContent'), {users});

    await selectUser(users[0]?.id);
    document.querySelectorAll('[data-action=changeUser]').forEach(link => link.addEventListener('click', async (e) => {
        await selectUser(e.target.getAttribute('data-user-id'));
    }));

    document.querySelector('[data-action=addUser]').addEventListener('click', showAddModal);
}
