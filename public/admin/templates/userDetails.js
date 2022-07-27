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
        <dl class="cosmo-key-value-list">
            <dt class="cosmo-key-value-list__key">Name</dt>
            <dd class="cosmo-key-value-list__value">${name}</dd>
            <dt class="cosmo-key-value-list__key">Email</dt>
            <dd class="cosmo-key-value-list__value">${email}</dd>
            <dt class="cosmo-key-value-list__key">Aktiv</dt>
            <dd class="cosmo-key-value-list__value">${activated ? 'Ja' : 'Nein'}</dd>
        </dl>`;
}