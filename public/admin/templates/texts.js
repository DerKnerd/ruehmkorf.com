import {html} from '../lib/js/jinya-html.js';

export default function texts({cookiesDe, dataProtectionDe, cookiesEn, dataProtectionEn}) {
    return html`
        <form>
            <div class="cosmo-tab-control" data-control="texts">
                <div class="cosmo-tab-control__tabs">
                    <a data-action="german" class="cosmo-tab-control__tab-link cosmo-tab-control__tab-link--active">Deutsch</a>
                    <a data-action="english" class="cosmo-tab-control__tab-link">Englisch</a>
                </div>
                <div data-tab="german" class="cosmo-tab-control__content">
                    <div class="cosmo-input__group">
                        <label for="cookiesDe" class="cosmo-label cosmo-label--textarea">Cookies</label>
                        <textarea class="cosmo-textarea cosmo-textarea--full-width" id="cookiesDe"
                                  rows="30">${cookiesDe}</textarea>
                        <label for="dataProtectionDe"
                               class="cosmo-label cosmo-label--textarea">Datenschutzerklärung</label>
                        <textarea class="cosmo-textarea cosmo-textarea--full-width" id="dataProtectionDe"
                                  rows="30">${dataProtectionDe}</textarea>
                    </div>
                </div>
                <div data-tab="english" class="cosmo-tab-control__content rc-hidden">
                    <div class="cosmo-input__group">
                        <label for="cookiesEn" class="cosmo-label cosmo-label--textarea">Cookies</label>
                        <textarea class="cosmo-textarea cosmo-textarea--full-width" id="cookiesEn"
                                  rows="30">${cookiesEn}</textarea>
                        <label for="dataProtectionEn"
                               class="cosmo-label cosmo-label--textarea">Datenschutzerklärung</label>
                        <textarea class="cosmo-textarea cosmo-textarea--full-width" id="dataProtectionEn"
                                  rows="30">${dataProtectionEn}</textarea>
                    </div>
                </div>
            </div>
            <div class="cosmo-button__container">
                <button class="cosmo-button" type="submit">Speichern</button>
            </div>
        </form>`;
}