import {html} from '../lib/js/jinya-html.js';

export default function newsAdd({titleDe, titleEn, public: isPublic, concatTags}) {
    return html`
        <div class="cosmo-modal__container">
            <form class="cosmo-modal">
                <div class="cosmo-modal__title">Nachricht hinzufügen</div>
                <div class="cosmo-modal__content">
                    <div class="cosmo-input__group">
                        <label class="cosmo-label" for="titleDe">Deutscher Titel</label>
                        <input autocomplete="false" id="titleDe" name="titleDe" placeholder="Deutscher Titel" required
                               type="text" class="cosmo-input" value="${titleDe}">
                        <label class="cosmo-label" for="titleEn">Englischer Titel</label>
                        <input autocomplete="false" id="titleEn" name="titleEn" placeholder="Englischer Titel" required
                               type="text" class="cosmo-input" value="${titleEn}">
                        <label class="cosmo-label" for="tags">Tags</label>
                        <input autocomplete="false" id="tags" name="tags" placeholder="Tags" type="text"
                               class="cosmo-input" value="${concatTags}">
                        <label for="heroImage" class="cosmo-label">Hero Bild</label>
                        <input class="cosmo-input" type="file" id="heroImage">
                        <div class="cosmo-input__group is--checkbox">
                            <input type="checkbox" id="public" class="cosmo-checkbox" ${isPublic ? 'checked' : ''}>
                            <label class="cosmo-label" for="public">Öffentlich</label>
                        </div>
                    </div>
                </div>
                <div class="cosmo-modal__button-bar">
                    <button class="cosmo-button" data-action="cancelEdit" type="button">Abbrechen</button>
                    <button class="cosmo-button" type="submit">Nachricht speichern</button>
                </div>
            </form>
        </div>`;
}