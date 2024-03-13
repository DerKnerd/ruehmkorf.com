import {html} from '../lib/js/jinya-html.js';

export default function profilesList({profiles}) {
    return html`
        <div class="cosmo-side-list">
            <nav class="cosmo-side-list__items">
                ${profiles.map(({
                                    name,
                                    id
                                }) => `<a class="cosmo-side-list__item" data-action="changeProfile" data-profile-id="${id}">${name}</a>`)}
                <button class="cosmo-button is--full-width" data-action="addProfile">Neues Profil</button>
            </nav>
            <div class="cosmo-side-list__content" id="profileContent">
            </div>
        </div>`;
}