import {html} from '../lib/js/jinya-html.js';

export default function downloadDetails({
                                            nameDe,
                                            nameEn,
                                            slug,
                                            public: isPublic,
                                            date,
                                            selfDestruct,
                                            selfDestructDays,
                                            descriptionDe,
                                            descriptionEn,
                                        }) {
    return html`
        <h1 class="cosmo-title">${nameDe} - ${nameEn}</h1>
        <div class="cosmo-toolbar">
            <div class="cosmo-toolbar__group">
                <button class="cosmo-button" type="button" data-action="editDownload">Bearbeiten</button>
                <button class="cosmo-button" type="button" data-action="deleteDownload">Löschen</button>
            </div>
            <div class="cosmo-toolbar__group">
                <button class="cosmo-button" type="button" data-action="uploadFile">Datei hochladen</button>
            </div>
        </div>
        <dl class="cosmo-list is--key-value">
            <dt>Slug</dt>
            <dd>${slug}</dd>
            <dt>Name - De</dt>
            <dd>${nameDe}</dd>
            <dt>Name - En</dt>
            <dd>${nameEn}</dd>
            <dt>Öffentlich</dt>
            <dd>${isPublic ? 'Ja' : 'Nein'}</dd>
            <dt>Deutscher Link</dt>
            <dd>
                <a href="/de/download/${slug}" target="_blank">https://ruehmkorf.com/de/download/${slug}</a>
            </dd>
            <dt>Englischer Link</dt>
            <dd>
                <a href="/en/download/${slug}" target="_blank">https://ruehmkorf.com/en/download/${slug}</a>
            </dd>
            <dt>Datum</dt>
            <dd>${date}</dd>
            ${selfDestruct ? `
                <dt>Löschen nach X Tagen</dt>
                <dd>${selfDestructDays}</dd>` : ''}
        </dl>
        <div class="cosmo-tab" data-control="download">
            <div class="cosmo-tab__links">
                <a data-action="german"
                   class="cosmo-tab__link is--active">Deutsch</a>
                <a data-action="english" class="cosmo-tab__link">Englisch</a>
                <a data-action="preview" class="cosmo-tab__link">Vorschau Bild</a>
            </div>
            <div data-tab="german" class="cosmo-tab__content">
                <h4>Inhalt</h4>
                <textarea class="cosmo-textarea" id="descriptionDe"
                          rows="20">${descriptionDe}</textarea>
            </div>
            <div data-tab="english" class="cosmo-tab__content rc-hidden">
                <h4>Inhalt</h4>
                <textarea class="cosmo-textarea" id="descriptionEn"
                          rows="20">${descriptionEn}</textarea>
            </div>
            <div data-tab="preview" class="cosmo-tab__content rc-hidden">
                <img src="/download/preview/${slug}" alt="Hero">
            </div>
        </div>
        <div class="cosmo-button__container">
            <button class="cosmo-button" type="button" id="saveDownload">Speichern</button>
        </div>`;
}