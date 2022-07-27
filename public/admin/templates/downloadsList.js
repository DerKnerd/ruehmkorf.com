import {html} from '../lib/js/jinya-html.js';

export default function downloadsList({downloads}) {
    return html`
        <div class="cosmo-list">
            <nav class="cosmo-list__items">
                ${downloads.map(({slug}) => `<a class="cosmo-list__item" data-action="changeDownload" data-download-slug="${slug}">${slug}</a>`)}
                <button class="cosmo-button cosmo-button--full-width" data-action="addDownload">Neuer Download</button>
            </nav>
            <div class="cosmo-list__content" id="downloadContent">
            </div>
        </div>`;
}