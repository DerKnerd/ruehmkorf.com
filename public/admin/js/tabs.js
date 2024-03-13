export function toggleTab(control, tab) {
    const tabControl = document.querySelector(`[data-control=${control}]`);
    tabControl.querySelectorAll('.is--active').forEach(link => link.classList.remove('is--active'));
    tabControl.querySelector(`[data-action=${tab}]`).classList.add('is--active');

    tabControl.querySelectorAll('.cosmo-tab__content').forEach(link => link.classList.add('rc-hidden'));
    tabControl.querySelector(`[data-tab=${tab}]`).classList.remove('rc-hidden');
}