import {compileTemplate} from "./template.js";

export async function confirm(title, message, confirm = 'Ok', decline = 'Abbrechen') {
    return await new Promise(async (resolve) => {
        const container = document.createElement('div');
        await compileTemplate('confirm.hbs', container, {title, message, confirm, decline});
        document.body.appendChild(container);

        container.querySelector('#confirm').addEventListener('click', () => {
            document.body.removeChild(container);
            resolve(true);
        });
        container.querySelector('#decline').addEventListener('click', () => {
            document.body.removeChild(container);
            resolve(false);
        });
    });
}

export async function alert(title, message, acknowledge = 'Ok') {
    return await new Promise(async (resolve) => {
        const container = document.createElement('div');
        await compileTemplate('alert.hbs', container, {title, message, acknowledge});
        document.body.appendChild(container);

        container.querySelector('#acknowledge').addEventListener('click', () => {
            document.body.removeChild(container);
            resolve(true);
        });
    });
}
