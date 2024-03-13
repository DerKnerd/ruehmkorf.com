import {html} from '../lib/js/jinya-html.js';

export default function downloadsList({downloads}) {
    return html`
        <div class="cosmo-side-list">
            <nav class="cosmo-side-list__items">
                ${downloads.map(({slug}) => `<a class="cosmo-side-list__item" data-action="changeDownload" data-download-slug="${slug}">${slug}</a>`)}
                <button class="cosmo-button is--full-width" data-action="addDownload">Neuer Download</button>
            </nav>
            <div class="cosmo-side-list__content" id="downloadContent">
            </div>
        </div>`;
}