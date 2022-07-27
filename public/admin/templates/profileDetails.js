import {html} from '../lib/js/jinya-html.js';

export default function profileDetails({name, url, active, id}) {
    return html`
        <h1 class="cosmo-title">${name}</h1>
        <div class="cosmo-toolbar">
            <div class="cosmo-toolbar__group">
                <button class="cosmo-button" type="button" data-action="editProfile">Bearbeiten</button>
                <button class="cosmo-button" type="button" data-action="deleteProfile">Löschen</button>
            </div>
        </div>
        <dl class="cosmo-key-value-list">
            <dt class="cosmo-key-value-list__key">Name</dt>
            <dd class="cosmo-key-value-list__value">${name}</dd>
            <dt class="cosmo-key-value-list__key">URL</dt>
            <dd class="cosmo-key-value-list__value">${url}</dd>
            <dt class="cosmo-key-value-list__key">Öffentlich</dt>
            <dd class="cosmo-key-value-list__value">${active ? 'Ja' : 'Nein'}</dd>
        </dl>
        <h4>Icon</h4>
        <img src="/admin/profile/icon?id=${id}" alt="Icon">
        <h4>Header</h4>
        <img src="/admin/profile/header?id=${id}" alt="Header">`;
}