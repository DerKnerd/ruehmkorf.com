import {html} from '../lib/js/jinya-html.js';

export default function newsDetails({
                                        titleDe,
                                        titleEn,
                                        slug,
                                        concatTags,
                                        public: isPublic,
                                        gistDe,
                                        gistEn,
                                        contentDe,
                                        contentEn,
                                    }) {
    return html`
        <h1 class="cosmo-title">${titleDe} - ${titleEn}</h1>
        <div class="cosmo-toolbar">
            <div class="cosmo-toolbar__group">
                <button class="cosmo-button" type="button" data-action="editNews">Bearbeiten</button>
                <button class="cosmo-button" type="button" data-action="deleteNews">Löschen</button>
            </div>
        </div>
        <dl class="cosmo-key-value-list">
            <dt class="cosmo-key-value-list__key">Slug</dt>
            <dd class="cosmo-key-value-list__value">${slug}</dd>
            <dt class="cosmo-key-value-list__key">Titel - De</dt>
            <dd class="cosmo-key-value-list__value">${titleDe}</dd>
            <dt class="cosmo-key-value-list__key">Titel - En</dt>
            <dd class="cosmo-key-value-list__value">${titleEn}</dd>
            <dt class="cosmo-key-value-list__key">Tags</dt>
            <dd class="cosmo-key-value-list__value">${concatTags}</dd>
            <dt class="cosmo-key-value-list__key">Öffentlich</dt>
            <dd class="cosmo-key-value-list__value">${isPublic ? 'Ja' : 'Nein'}</dd>
            <dt class="cosmo-key-value-list__key">Deutscher Link</dt>
            <dd class="cosmo-key-value-list__value">
                <a href="/de/news/${slug}" target="_blank">https://ruehmkorf.com/de/news/${slug}</a>
            </dd>
            <dt class="cosmo-key-value-list__key">Englischer Link</dt>
            <dd class="cosmo-key-value-list__value">
                <a href="/en/news/${slug}" target="_blank">https://ruehmkorf.com/en/news/${slug}</a>
            </dd>
        </dl>
        <div class="cosmo-tab-control" data-control="news">
            <div class="cosmo-tab-control__tabs">
                <a data-action="german"
                   class="cosmo-tab-control__tab-link cosmo-tab-control__tab-link--active">Deutsch</a>
                <a data-action="english" class="cosmo-tab-control__tab-link">Englisch</a>
                <a data-action="hero" class="cosmo-tab-control__tab-link">Hero Bild</a>
            </div>
            <div data-tab="german" class="cosmo-tab-control__content">
                <h4>Anreißer</h4>
                <textarea class="cosmo-textarea cosmo-textarea--full-width" id="gistDe" rows="10">${gistDe}</textarea>
                <h4>Inhalt</h4>
                <textarea class="cosmo-textarea cosmo-textarea--full-width" id="contentDe"
                          rows="20">${contentDe}</textarea>
            </div>
            <div data-tab="english" class="cosmo-tab-control__content rc-hidden">
                <h4>Anreißer</h4>
                <textarea class="cosmo-textarea cosmo-textarea--full-width" id="gistEn" rows="10">${gistEn}</textarea>
                <h4>Inhalt</h4>
                <textarea class="cosmo-textarea cosmo-textarea--full-width" id="contentEn"
                          rows="20">${contentEn}</textarea>
            </div>
            <div data-tab="hero" class="cosmo-tab-control__content rc-hidden">
                <img src="/news/hero/${slug}" alt="Hero">
            </div>
        </div>
        <div class="cosmo-button__container">
            <button class="cosmo-button" type="button" id="saveNews">Speichern</button>
        </div>`;
}