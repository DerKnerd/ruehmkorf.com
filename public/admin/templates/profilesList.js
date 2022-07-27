import {html} from '../lib/js/jinya-html.js';

export default function profilesList({profiles}) {
    return html`
        <div class="cosmo-list">
            <nav class="cosmo-list__items">
                ${profiles.map(({
                                    name,
                                    id
                                }) => `<a class="cosmo-list__item" data-action="changeProfile" data-profile-id="${id}">${name}</a>`)}
                <button class="cosmo-button cosmo-button--full-width" data-action="addProfile">Neues Profil</button>
            </nav>
            <div class="cosmo-list__content" id="profileContent">
            </div>
        </div>`;
}