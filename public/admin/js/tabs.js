export function toggleTab(control, tab) {
    const tabControl = document.querySelector(`[data-control=${control}]`);
    tabControl.querySelectorAll('.cosmo-tab-control__tab-link--active').forEach(link => link.classList.remove('cosmo-tab-control__tab-link--active'));
    tabControl.querySelector(`[data-action=${tab}]`).classList.add('cosmo-tab-control__tab-link--active');

    tabControl.querySelectorAll('.cosmo-tab-control__content').forEach(link => link.classList.add('rc-hidden'));
    tabControl.querySelector(`[data-tab=${tab}]`).classList.remove('rc-hidden');
}