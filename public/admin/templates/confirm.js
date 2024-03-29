import {html} from '../lib/js/jinya-html.js';

export default function confirm({title, message, decline, confirm}) {
    return html`
        <div class="cosmo-modal__container">
            <div class="cosmo-modal">
                <h1 class="cosmo-modal__title">${title}</h1>
                <div class="cosmo-modal__content">
                    <p>${message}</p>
                </div>
                <div class="cosmo-modal__button-bar">
                    <button id="decline" class="cosmo-button">${decline}</button>
                    <button id="confirm" class="cosmo-button">${confirm}</button>
                </div>
            </div>
        </div>`;
}