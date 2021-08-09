export function unmarkListLinks() {
    document.querySelectorAll('.cosmo-list__item--active').forEach((item) => {
        item.classList.remove('cosmo-list__item--active');
    });
}