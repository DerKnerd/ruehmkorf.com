import {html} from '../lib/js/jinya-html.js';

export default function userEdit({name, email, activated}) {
    return html`
        <div class="cosmo-modal__container">
            <form class="cosmo-modal">
                <div class="cosmo-modal__title">Benutzer bearbeiten</div>
                <div class="cosmo-modal__content">
                    <div class="cosmo-input__group">
                        <label class="cosmo-label" for="name">Name</label>
                        <input autocomplete="false" id="name" name="name" placeholder="Name" required type="text"
                               class="cosmo-input" value="${name}">
                        <label class="cosmo-label" for="email">Email</label>
                        <input autocomplete="false" id="email" name="email" placeholder="Email" type="email"
                               class="cosmo-input" value="${email}">
                        <label class="cosmo-label" for="password">Passwort</label>
                        <input autocomplete="false" id="password" name="password" placeholder="Passwort" type="password"
                               class="cosmo-input">
                        <div class="cosmo-input__group is--checkbox">
                            <input type="checkbox" id="activated" class="cosmo-checkbox" ${activated ? 'checked' : ''}>
                            <label for="activated">Aktiv</label>
                        </div>
                    </div>
                </div>
                <div class="cosmo-modal__button-bar">
                    <button class="cosmo-button" data-action="cancelEdit" type="button">Abbrechen</button>
                    <button class="cosmo-button" type="submit">Benutzer speichern</button>
                </div>
            </form>
        </div>`;
}