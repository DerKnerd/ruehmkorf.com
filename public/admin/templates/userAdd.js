import {html} from '../lib/js/jinya-html.js';

export default function userAdd() {
    return html`
        <div class="cosmo-modal__container">
            <form class="cosmo-modal">
                <div class="cosmo-modal__title">Benutzer hinzufügen</div>
                <div class="cosmo-modal__content">
                    <div class="cosmo-input__group">
                        <label class="cosmo-label" for="name">Name</label>
                        <input autocomplete="false" id="name" name="name" placeholder="Name" required type="text"
                               class="cosmo-input">
                        <label class="cosmo-label" for="email">Email</label>
                        <input autocomplete="false" id="email" name="email" placeholder="Email" required type="email"
                               class="cosmo-input">
                        <label class="cosmo-label" for="password">Passwort</label>
                        <input autocomplete="false" id="password" name="password" placeholder="Passwort" required
                               type="password" class="cosmo-input">
                        <div class="cosmo-input__group is--checkbox">
                            <input type="checkbox" id="activated" class="cosmo-checkbox">
                            <label for="activated">Aktiv</label>
                        </div>
                    </div>
                </div>
                <div class="cosmo-modal__button-bar">
                    <button class="cosmo-button" data-action="cancelAdd" type="button">Abbrechen</button>
                    <button class="cosmo-button" type="submit">Benutzer speichern</button>
                </div>
            </form>
        </div>`;
}