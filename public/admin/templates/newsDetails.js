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
        <dl class="cosmo-list is--key-value">
            <dt>Slug</dt>
            <dd>${slug}</dd>
            <dt>Titel - De</dt>
            <dd>${titleDe}</dd>
            <dt>Titel - En</dt>
            <dd>${titleEn}</dd>
            <dt>Tags</dt>
            <dd>${concatTags}</dd>
            <dt>Öffentlich</dt>
            <dd>${isPublic ? 'Ja' : 'Nein'}</dd>
            <dt>Deutscher Link</dt>
            <dd>
                <a href="/de/news/${slug}" target="_blank">https://ruehmkorf.com/de/news/${slug}</a>
            </dd>
            <dt>Englischer Link</dt>
            <dd>
                <a href="/en/news/${slug}" target="_blank">https://ruehmkorf.com/en/news/${slug}</a>
            </dd>
        </dl>
        <div class="cosmo-tab" data-control="news">
            <div class="cosmo-tab__links">
                <a data-action="german"
                   class="cosmo-tab__link is--active">Deutsch</a>
                <a data-action="english" class="cosmo-tab__link">Englisch</a>
                <a data-action="hero" class="cosmo-tab__link">Hero Bild</a>
            </div>
            <div data-tab="german" class="cosmo-tab__content">
                <h4>Anreißer</h4>
                <textarea class="cosmo-textarea" id="gistDe" rows="10">${gistDe}</textarea>
                <h4>Inhalt</h4>
                <textarea class="cosmo-textarea" id="contentDe" rows="20">${contentDe}</textarea>
            </div>
            <div data-tab="english" class="cosmo-tab__content rc-hidden">
                <h4>Anreißer</h4>
                <textarea class="cosmo-textarea" id="gistEn" rows="10">${gistEn}</textarea>
                <h4>Inhalt</h4>
                <textarea class="cosmo-textarea" id="contentEn" rows="20">${contentEn}</textarea>
            </div>
            <div data-tab="hero" class="cosmo-tab__content rc-hidden">
                <img src="/news/hero/${slug}" alt="Hero">
            </div>
        </div>
        <div class="cosmo-button__container">
            <button class="cosmo-button" type="button" id="saveNews">Speichern</button>
        </div>`;
}