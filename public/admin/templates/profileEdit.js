import {html} from '../lib/js/jinya-html.js';

export default function profileEdit({name, url, active}) {
    return html`
        <div class="cosmo-modal__container">
            <form class="cosmo-modal">
                <div class="cosmo-modal__title">Profil bearbeiten</div>
                <div class="cosmo-modal__content">
                    <div class="cosmo-input__group">
                        <label class="cosmo-label" for="name">Name</label>
                        <input autocomplete="false" id="name" name="name" placeholder="Name" required type="text"
                               value="${name}" class="cosmo-input">
                        <label class="cosmo-label" for="url">URL</label>
                        <input autocomplete="false" id="url" name="url" placeholder="URL" type="url" class="cosmo-input"
                               value="${url}">
                        <label for="icon" class="cosmo-label">Icon</label>
                        <input class="cosmo-input" type="file" id="icon">
                        <label for="headerImage" class="cosmo-label">Header Bild</label>
                        <input class="cosmo-input" type="file" id="headerImage">
                        <div class="cosmo-checkbox__group">
                            <input type="checkbox" id="active" class="cosmo-checkbox" ${active ? 'checked' : ''}>
                            <label class="cosmo-label" for="active">Aktiv</label>
                        </div>
                    </div>
                </div>
                <div class="cosmo-modal__button-bar">
                    <button class="cosmo-button" data-action="cancelAdd" type="button">Abbrechen</button>
                    <button class="cosmo-button" type="submit">Profil speichern</button>
                </div>
            </form>
        </div>`;
}