import {html} from '../lib/js/jinya-html.js';

export default function newsList({news}) {
    return html`
        <div class="cosmo-list">
            <nav class="cosmo-list__items">
                ${news.map(({slug}) => `<a class="cosmo-list__item" data-action="changeNews" data-news-slug="${slug}">${slug}</a>`)}
                <button class="cosmo-button cosmo-button--full-width" data-action="addNews">Neue Nachricht</button>
            </nav>
            <div class="cosmo-list__content" id="newsContent">
            </div>
        </div>`;
}