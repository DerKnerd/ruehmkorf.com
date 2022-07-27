const templateCache = {};

export async function compileTemplate(name, element, data = {}) {
    let html = 'Fehler beim laden';
    if (templateCache[name]) {
        html = templateCache[name](data);
    } else {
        const tmpl = (await import(`/public/admin/templates/${name}`)).default;
        templateCache[name] = tmpl;
        html = tmpl(data);
    }

    element.innerHTML = html;
}