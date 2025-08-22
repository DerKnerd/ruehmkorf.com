import { put } from '../lib/jinya-http.js';

Alpine.data('accountData', () => ({
  changePasswordDialog: {
    open: false,
    oldPassword: '',
    newPassword: '',
    error: false,
    openDialog() {
      this.open = true;
      this.oldPassword = '';
      this.newPassword = '';
      this.error = false;
    },
  },
  async changePassword() {
    try {
      await put('/api/authentication/password', {
        oldPassword: this.changePasswordDialog.oldPassword,
        newPassword: this.changePasswordDialog.newPassword,
      });
      this.changePasswordDialog.open = false;
      await Alpine.store('authentication').logout();
    } catch (e) {
      this.changePasswordDialog.error = true;
    }
  },
}));
