import {unmarkListLinks} from "./lists.js";
import {toggleTab} from "./tabs.js";
import {alert, confirm} from "./dialogs.js";
import {unmarkSubMenuLinks} from "./navigation.js";
import {compileTemplate} from "./template.js";

async function selectNews(slug) {
    if (!slug) {
        return;
    }

    unmarkListLinks();

    const element = document.querySelector(`[data-news-slug="${slug}"]`);
    element.classList.add('cosmo-list__item--active');

    const news = await (await fetch(`/admin/news?slug=${slug}`)).json();

    news.concatTags = news.tags ? news.tags.map(tag => tag.tag).join(', ') : 'Keine';
    await compileTemplate('newsDetails.js', document.getElementById('newsContent'), news);
    document.querySelector('[data-action=german]').addEventListener('click', () => toggleTab('news', 'german'));
    document.querySelector('[data-action=english]').addEventListener('click', () => toggleTab('news', 'english'));
    document.querySelector('[data-action=hero]').addEventListener('click', () => toggleTab('news', 'hero'));

    document.querySelector('[data-action=deleteNews]').addEventListener('click', async () => {
        const result = await confirm('Nachricht löschen', `Soll die Nachricht ${news.slug} wirklich gelöscht werden?`, 'Nachricht löschen', 'Nachricht behalten');
        if (result) {
            await fetch(`/admin/news?slug=${slug}`, {method: 'DELETE'});
            await init();
        }
    });

    document.getElementById('saveNews').addEventListener('click', async () => {
        const contentDe = document.getElementById('contentDe').value;
        const contentEn = document.getElementById('contentEn').value;
        const gistDe = document.getElementById('gistDe').value;
        const gistEn = document.getElementById('gistEn').value;

        const response = await fetch(`/admin/news?slug=${slug}`, {
            method: 'PUT',
            body: JSON.stringify({
                ...news,
                contentDe,
                contentEn,
                gistDe,
                gistEn
            }),
            headers: {
                'Content-Type': 'application/json',
            },
        });
        if (response.status !== 204) {
            await alert('Speichern fehlgeschlagen', 'Beim Speichern ist ein unbekannter Fehler aufgetreten.');
        }
    });
    document.querySelector('[data-action=editNews]').addEventListener('click', () => showEditModal({
        ...news,
        concatTags: news.tags ? news.tags.map(tag => tag.tag).join(',') : '',
    }));
}

async function showEditModal(news) {
    const container = document.createElement('div');
    await compileTemplate('newsEdit.js', container, news);
    document.body.appendChild(container);

    container.querySelector('form').addEventListener('submit', async (e) => {
        e.preventDefault();
        const heroImage = document.getElementById('heroImage');
        const titleDe = document.getElementById('titleDe').value;
        const titleEn = document.getElementById('titleEn').value;
        const tags = document.getElementById('tags').value.split(',');
        const publicChecked = document.getElementById('public').checked;
        const result = await fetch(`/admin/news?slug=${news.slug}`, {
            body: JSON.stringify({
                ...news,
                titleDe,
                titleEn,
                public: publicChecked,
                tags,
            }),
            method: 'PUT',
        });
        if (result.status !== 204) {
            await alert('Speichern fehlgeschlagen', 'Beim Speichern ist ein unbekannter Fehler aufgetreten');
        } else {
            if (heroImage.files.length > 0) {
                const fileUploadResult = await fetch(`/admin/news/hero?slug=${news.slug}`, {
                    method: 'POST',
                    body: heroImage.files.item(0)
                });
                if (fileUploadResult.status !== 204) {
                    await alert('Speichern fehlgeschlagen', 'Beim Speichern des Hero Bildes ist ein unbekannter Fehler aufgetreten');
                    return;
                }
            }
            document.body.removeChild(container);
            await selectNews(news.slug);
        }
    });
    container.querySelector('[data-action=cancelEdit]').addEventListener('click', () => document.body.removeChild(container));
}

async function showAddModal() {
    const container = document.createElement('div');
    await compileTemplate('newsAdd.js', container);
    document.body.appendChild(container);

    container.querySelector('form').addEventListener('submit', async (e) => {
        e.preventDefault();
        const heroImage = document.getElementById('heroImage');
        const titleDe = document.getElementById('titleDe').value;
        const titleEn = document.getElementById('titleEn').value;
        const slug = document.getElementById('slug').value;
        const tags = document.getElementById('tags').value.split(',');
        const publicChecked = document.getElementById('public').checked;
        const result = await fetch(`/admin/news`, {
            body: JSON.stringify({
                slug,
                titleDe,
                titleEn,
                public: publicChecked,
                tags,
            }),
            method: 'POST',
        });
        if (result.status === 409) {
            await alert('Speichern fehlgeschlagen', `Eine Nachricht mit dem Slug ${slug} existiert bereits`);
        } else if (result.status !== 201) {
            await alert('Speichern fehlgeschlagen', 'Beim Speichern ist ein unbekannter Fehler aufgetreten');
        } else {
            if (heroImage.files.length > 0) {
                const fileUploadResult = await fetch(`/admin/news/hero?slug=${slug}`, {
                    method: 'POST',
                    body: heroImage.files.item(0)
                });
                if (fileUploadResult.status !== 204) {
                    await alert('Speichern fehlgeschlagen', 'Beim Speichern des Hero Bildes ist ein unbekannter Fehler aufgetreten');
                    return
                }
            }
            document.body.removeChild(container);
            await init();
            await selectNews(slug);
        }
    });
    container.querySelector('[data-action=cancelAdd]').addEventListener('click', () => document.body.removeChild(container));
}

export async function init() {
    unmarkSubMenuLinks();
    document.querySelector('[data-sublink=news]').classList.add('cosmo-menu-bar__sub-item--active');
    const news = await (await fetch('/admin/news')).json();
    await compileTemplate('newsList.js', document.getElementById('rcContent'), {news});

    await selectNews(news[0]?.slug);
    document.querySelectorAll('[data-action=changeNews]').forEach(link => link.addEventListener('click', async (e) => {
        await selectNews(e.target.getAttribute('data-news-slug'));
    }));

    document.querySelector('[data-action=addNews]').addEventListener('click', showAddModal);
}