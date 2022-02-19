const templateCache = {};

export async function compileTemplate(name, element, data = {}) {
    let html = 'Fehler beim laden';
    if (templateCache[name]) {
        html = templateCache[name](data);
    } else {
        const hbs = await fetch(`/public/admin/templates/${name}`);
        if (hbs.status === 200) {
            const template = Handlebars.compile(await hbs.text());
            templateCache[name] = template;
            html = template(data);
        }
    }

    element.innerHTML = html;
}