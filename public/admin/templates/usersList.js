import {html} from '../lib/js/jinya-html.js';

export default function usersList({users}) {
    return html`
        <div class="cosmo-list">
            <nav class="cosmo-list__items">
                ${users.map(({
                                 id,
                                 name
                             }) => `<a class="cosmo-list__item" data-action="changeUser" data-user-id="${id}">${name}</a>`)}
                <button class="cosmo-button cosmo-button--full-width" data-action="addUser">Neuer Benutzer</button>
            </nav>
            <div class="cosmo-list__content" id="userContent">
            </div>
        </div>`;
}