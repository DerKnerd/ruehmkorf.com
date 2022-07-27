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
        <dl class="cosmo-key-value-list">
            <dt class="cosmo-key-value-list__key">Slug</dt>
            <dd class="cosmo-key-value-list__value">${slug}</dd>
            <dt class="cosmo-key-value-list__key">Name - De</dt>
            <dd class="cosmo-key-value-list__value">${nameDe}</dd>
            <dt class="cosmo-key-value-list__key">Name - En</dt>
            <dd class="cosmo-key-value-list__value">${nameEn}</dd>
            <dt class="cosmo-key-value-list__key">Öffentlich</dt>
            <dd class="cosmo-key-value-list__value">${isPublic ? 'Ja' : 'Nein'}</dd>
            <dt class="cosmo-key-value-list__key">Deutscher Link</dt>
            <dd class="cosmo-key-value-list__value">
                <a href="/de/download/${slug}" target="_blank">https://ruehmkorf.com/de/download/${slug}</a>
            </dd>
            <dt class="cosmo-key-value-list__key">Englischer Link</dt>
            <dd class="cosmo-key-value-list__value">
                <a href="/en/download/${slug}" target="_blank">https://ruehmkorf.com/en/download/${slug}</a>
            </dd>
            <dt class="cosmo-key-value-list__key">Datum</dt>
            <dd class="cosmo-key-value-list__value">${date}</dd>
            ${selfDestruct ? `
                <dt class="cosmo-key-value-list__key">Löschen nach X Tagen</dt>
                <dd class="cosmo-key-value-list__value">${selfDestructDays}</dd>` : ''}
        </dl>
        <div class="cosmo-tab-control" data-control="download">
            <div class="cosmo-tab-control__tabs">
                <a data-action="german"
                   class="cosmo-tab-control__tab-link cosmo-tab-control__tab-link--active">Deutsch</a>
                <a data-action="english" class="cosmo-tab-control__tab-link">Englisch</a>
                <a data-action="preview" class="cosmo-tab-control__tab-link">Vorschau Bild</a>
            </div>
            <div data-tab="german" class="cosmo-tab-control__content">
                <h4>Inhalt</h4>
                <textarea class="cosmo-textarea cosmo-textarea--full-width" id="descriptionDe"
                          rows="20">${descriptionDe}</textarea>
            </div>
            <div data-tab="english" class="cosmo-tab-control__content rc-hidden">
                <h4>Inhalt</h4>
                <textarea class="cosmo-textarea cosmo-textarea--full-width" id="descriptionEn"
                          rows="20">${descriptionEn}</textarea>
            </div>
            <div data-tab="preview" class="cosmo-tab-control__content rc-hidden">
                <img src="/download/preview/${slug}" alt="Hero">
            </div>
        </div>
        <div class="cosmo-button__container">
            <button class="cosmo-button" type="button" id="saveDownload">Speichern</button>
        </div>`;
}