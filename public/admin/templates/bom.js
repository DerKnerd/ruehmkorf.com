import {html} from '../lib/js/jinya-html.js';

export default function bom({descriptionDe, infoTextDeDe, infoTextEnDe, descriptionEn, infoTextDeEn, infoTextEnEn}) {
    return html`
        <form>
            <div class="cosmo-tab" data-control="bom">
                <div class="cosmo-tab__links">
                    <a data-action="german" class="cosmo-tab__link is--active">Deutsch</a>
                    <a data-action="english" class="cosmo-tab__link">Englisch</a>
                </div>
                <div data-tab="german" class="cosmo-tab__content">
                    <div class="cosmo-input__group">
                        <label for="descriptionDe" class="cosmo-label is--textarea">Beschreibung</label>
                        <textarea class="cosmo-textarea" id="descriptionDe"
                                  rows="30">${descriptionDe}</textarea>
                        <label for="infoTextDeDe" class="cosmo-label is--textarea">Infotext Deutsch</label>
                        <textarea class="cosmo-textarea" id="infoTextDeDe"
                                  rows="30">${infoTextDeDe}</textarea>
                        <label for="infoTextEnDe" class="cosmo-label is--textarea">Infotext Englisch</label>
                        <textarea class="cosmo-textarea" id="infoTextEnDe"
                                  rows="30">${infoTextEnDe}</textarea>
                    </div>
                </div>
                <div data-tab="english" class="cosmo-tab__content rc-hidden">
                    <div class="cosmo-input__group">
                        <label for="descriptionEn" class="cosmo-label is--textarea">Beschreibung</label>
                        <textarea class="cosmo-textarea" id="descriptionEn"
                                  rows="30">${descriptionEn}</textarea>
                        <label for="infoTextDeEn" class="cosmo-label is--textarea">Infotext Deutsch</label>
                        <textarea class="cosmo-textarea" id="infoTextDeEn"
                                  rows="30">${infoTextDeEn}</textarea>
                        <label for="infoTextEnEn" class="cosmo-label is--textarea">Infotext Englisch</label>
                        <textarea class="cosmo-textarea" id="infoTextEnEn"
                                  rows="30">${infoTextEnEn}</textarea>
                    </div>
                </div>
            </div>
            <div class="cosmo-button__container">
                <button class="cosmo-button" type="submit">Speichern</button>
            </div>
        </form>`;
}