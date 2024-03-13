import {html} from '../lib/js/jinya-html.js';

export default function userDetails({name, email, activated}) {
    return html`
        <h1 class="cosmo-title">${name}</h1>
        <div class="cosmo-toolbar">
            <div class="cosmo-toolbar__group">
                <button class="cosmo-button" type="button" data-action="editUser">Bearbeiten</button>
                <button class="cosmo-button" type="button" data-action="deleteUser">Löschen</button>
            </div>
        </div>
        <dl class="cosmo-list is--key-value">
            <dt>Name</dt>
            <dd>${name}</dd>
            <dt>Email</dt>
            <dd>${email}</dd>
            <dt>Aktiv</dt>
            <dd>${activated ? 'Ja' : 'Nein'}</dd>
        </dl>`;
}