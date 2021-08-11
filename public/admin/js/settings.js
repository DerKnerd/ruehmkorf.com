export async function navigate(sublink) {
    const data = await import((`./${sublink}.js`));
    await data.init();
}

export async function init() {
    document.querySelectorAll('[data-sublink]').forEach(link => link.addEventListener('click', async (e) => await navigate(e.target.getAttribute('data-sublink'))));

    await navigate('icons');
}
