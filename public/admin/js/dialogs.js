import {alertTemplate, confirmTemplate} from "./templates.js";

export async function confirm(title, message, confirm = 'Ok', decline = 'Abbrechen') {
    return await new Promise((resolve) => {
        const content = confirmTemplate({title, message, confirm, decline});
        const container = document.createElement('div');
        container.innerHTML = content;
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
    return await new Promise((resolve) => {
        const content = alertTemplate({title, message, acknowledge});
        const container = document.createElement('div');
        container.innerHTML = content;
        document.body.appendChild(container);

        container.querySelector('#acknowledge').addEventListener('click', () => {
            document.body.removeChild(container);
            resolve(true);
        });
    });
}
