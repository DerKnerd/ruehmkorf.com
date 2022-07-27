import {html} from '../lib/js/jinya-html.js';

export default function icons() {
    return html`
        <form>
            <div class="cosmo-input__group">
                <label for="logo" class="cosmo-label">Logo</label>
                <input class="cosmo-input" type="file" id="logo">
                <label for="touchIcon" class="cosmo-label">Touchicon</label>
                <input class="cosmo-input" type="file" id="touchIcon">
                <label for="favicon" class="cosmo-label">Favicon</label>
                <input class="cosmo-input" type="file" id="favicon">
            </div>
            <div class="cosmo-button__container">
                <button class="cosmo-button" type="submit">Speichern</button>
            </div>
        </form>
        <h4>Logo</h4>
        <img src="/logo.png" alt="Logo">
        <h4>Touchicon</h4>
        <img src="/touchicon.png" alt="Touchicon">
        <h4>Favicon</h4>
        <img src="/favicon.ico" alt="Favicon">
    `;
}