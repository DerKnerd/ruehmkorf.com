import {html} from '../lib/js/jinya-html.js';

export default function downloadEdit({nameDe, nameEn, selfDestructDays, public: isPublic}) {
    return html`
        <div class="cosmo-modal__container">
            <form class="cosmo-modal">
                <div class="cosmo-modal__title">Download bearbeiten</div>
                <div class="cosmo-modal__content">
                    <div class="cosmo-input__group">
                        <label class="cosmo-label" for="nameDe">Deutscher Name</label>
                        <input autocomplete="false" id="nameDe" name="nameDe" placeholder="Deutscher Name" required
                               value="${nameDe}" type="text" class="cosmo-input">
                        <label class="cosmo-label" for="nameEn">Englischer Name</label>
                        <input autocomplete="false" id="nameEn" name="nameEn" placeholder="Englischer Titel" required
                               value="${nameEn}" type="text" class="cosmo-input">
                        <label class="cosmo-label" for="date">Datum</label>
                        <input autocomplete="false" id="date" name="date" placeholder="Datum" required type="date"
                               class="cosmo-input">
                        <label class="cosmo-label" for="selfDestructDays">Löschen nach X Tagen</label>
                        <input autocomplete="false" id="selfDestructDays" name="selfDestructDays" class="cosmo-input"
                               placeholder="Löschen nach X Tagen" type="number" value="${selfDestructDays}">
                        <label for="heroImage" class="cosmo-label">Vorschau Bild</label>
                        <div class="cosmo-input cosmo-input--picker">
                            <input class="cosmo-input" type="file" id="previewImage">
                        </div>
                        <div class="cosmo-checkbox__group">
                            <input type="checkbox" id="public" class="cosmo-checkbox" ${isPublic ? 'Ja' : 'Nein'}>
                            <label class="cosmo-label" for="public">Öffentlich</label>
                        </div>
                    </div>
                </div>
                <div class="cosmo-modal__button-bar">
                    <button class="cosmo-button" data-action="cancelEdit" type="button">Abbrechen</button>
                    <button class="cosmo-button" type="submit">Download speichern</button>
                </div>
            </form>
        </div>`;
}