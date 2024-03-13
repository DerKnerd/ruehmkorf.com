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
        <dl class="cosmo-list is--key-value">
            <dt>Name</dt>
            <dd>${name}</dd>
            <dt>URL</dt>
            <dd>${url}</dd>
            <dt>Öffentlich</dt>
            <dd>${active ? 'Ja' : 'Nein'}</dd>
        </dl>
        <h4>Icon</h4>
        <img src="/admin/profile/icon?id=${id}" alt="Icon">
        <h4>Header</h4>
        <img src="/admin/profile/header?id=${id}" alt="Header">`;
}