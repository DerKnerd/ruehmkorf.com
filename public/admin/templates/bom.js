import {html} from '../lib/js/jinya-html.js';

export default function bom({descriptionDe, infoTextDeDe, infoTextEnDe, descriptionEn, infoTextDeEn, infoTextEnEn}) {
    return html`
        <form>
            <div class="cosmo-tab-control" data-control="bom">
                <div class="cosmo-tab-control__tabs">
                    <a data-action="german" class="cosmo-tab-control__tab-link cosmo-tab-control__tab-link--active">Deutsch</a>
                    <a data-action="english" class="cosmo-tab-control__tab-link">Englisch</a>
                </div>
                <div data-tab="german" class="cosmo-tab-control__content">
                    <div class="cosmo-input__group">
                        <label for="descriptionDe" class="cosmo-label cosmo-label--textarea">Beschreibung</label>
                        <textarea class="cosmo-textarea cosmo-textarea--full-width" id="descriptionDe"
                                  rows="30">${descriptionDe}</textarea>
                        <label for="infoTextDeDe" class="cosmo-label cosmo-label--textarea">Infotext Deutsch</label>
                        <textarea class="cosmo-textarea cosmo-textarea--full-width" id="infoTextDeDe"
                                  rows="30">${infoTextDeDe}</textarea>
                        <label for="infoTextEnDe" class="cosmo-label cosmo-label--textarea">Infotext Englisch</label>
                        <textarea class="cosmo-textarea cosmo-textarea--full-width" id="infoTextEnDe"
                                  rows="30">${infoTextEnDe}</textarea>
                    </div>
                </div>
                <div data-tab="english" class="cosmo-tab-control__content rc-hidden">
                    <div class="cosmo-input__group">
                        <label for="descriptionEn" class="cosmo-label cosmo-label--textarea">Beschreibung</label>
                        <textarea class="cosmo-textarea cosmo-textarea--full-width" id="descriptionEn"
                                  rows="30">${descriptionEn}</textarea>
                        <label for="infoTextDeEn" class="cosmo-label cosmo-label--textarea">Infotext Deutsch</label>
                        <textarea class="cosmo-textarea cosmo-textarea--full-width" id="infoTextDeEn"
                                  rows="30">${infoTextDeEn}</textarea>
                        <label for="infoTextEnEn" class="cosmo-label cosmo-label--textarea">Infotext Englisch</label>
                        <textarea class="cosmo-textarea cosmo-textarea--full-width" id="infoTextEnEn"
                                  rows="30">${infoTextEnEn}</textarea>
                    </div>
                </div>
            </div>
            <div class="cosmo-button__container">
                <button class="cosmo-button" type="submit">Speichern</button>
            </div>
        </form>`;
}