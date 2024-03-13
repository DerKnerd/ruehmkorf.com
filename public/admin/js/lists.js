export function unmarkListLinks() {
    document.querySelectorAll('.is--active').forEach((item) => {
        item.classList.remove('is--active');
    });
}