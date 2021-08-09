const newsListTemplateContent = document.getElementById('newsListTemplate').innerHTML;
export const newsListTemplate = Handlebars.compile(newsListTemplateContent);

const newsDetailsTemplateContent = document.getElementById('newsDetailTemplate').innerHTML;
export const newsDetailsTemplate = Handlebars.compile(newsDetailsTemplateContent);

const confirmTemplateContent = document.getElementById('confirmTemplate').innerHTML;
export const confirmTemplate = Handlebars.compile(confirmTemplateContent);

const alertTemplateContent = document.getElementById('alertTemplate').innerHTML;
export const alertTemplate = Handlebars.compile(alertTemplateContent);

const editNewsTemplateContent = document.getElementById('editNewsTemplate').innerHTML;
export const editNewsTemplate = Handlebars.compile(editNewsTemplateContent);

const addNewsTemplateContent = document.getElementById('addNewsTemplate').innerHTML;
export const addNewsTemplate = Handlebars.compile(addNewsTemplateContent);
