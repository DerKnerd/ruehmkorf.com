import {html} from '../lib/js/jinya-html.js';

export default function texts({cookiesDe, dataProtectionDe, cookiesEn, dataProtectionEn}) {
    return html`
        <form>
            <div class="cosmo-tab" data-control="texts">
                <div class="cosmo-tab__links">
                    <a data-action="german" class="cosmo-tab__link is--active">Deutsch</a>
                    <a data-action="english" class="cosmo-tab__link">Englisch</a>
                </div>
                <div data-tab="german" class="cosmo-tab__content">
                    <div class="cosmo-input__group">
                        <label for="cookiesDe" class="cosmo-label is--textarea">Cookies</label>
                        <textarea class="cosmo-textarea" id="cookiesDe"
                                  rows="30">${cookiesDe}</textarea>
                        <label for="dataProtectionDe"
                               class="cosmo-label is--textarea">Datenschutzerklärung</label>
                        <textarea class="cosmo-textarea" id="dataProtectionDe"
                                  rows="30">${dataProtectionDe}</textarea>
                    </div>
                </div>
                <div data-tab="english" class="cosmo-tab__content rc-hidden">
                    <div class="cosmo-input__group">
                        <label for="cookiesEn" class="cosmo-label is--textarea">Cookies</label>
                        <textarea class="cosmo-textarea" id="cookiesEn"
                                  rows="30">${cookiesEn}</textarea>
                        <label for="dataProtectionEn"
                               class="cosmo-label is--textarea">Datenschutzerklärung</label>
                        <textarea class="cosmo-textarea" id="dataProtectionEn"
                                  rows="30">${dataProtectionEn}</textarea>
                    </div>
                </div>
            </div>
            <div class="cosmo-button__container">
                <button class="cosmo-button" type="submit">Speichern</button>
            </div>
        </form>`;
}