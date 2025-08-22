import { get, httpDelete, post, put } from '../lib/jinya-http.js';
import confirm from '../lib/ui/confirm.js';
import alert from '../lib/ui/alert.js';

import '../lib/ui/toolbar-editor.js';

Alpine.data('profilesData', () => ({
  loading: true,
  profiles: [],
  selectedProfile: null,
  addProfile: {
    open: false,
    linkTarget: '',
    linkLabel: '',
    description: '',
    error: false,
    openDialog() {
      this.linkTarget = '';
      this.linkLabel = '';
      this.description = '';
      this.error = false;
      this.open = true;
    },
  },
  editProfile: {
    open: false,
    linkTarget: '',
    linkLabel: '',
    description: '',
    error: false,
    openDialog(profile) {
      this.linkTarget = profile.linkTarget;
      this.linkLabel = profile.linkLabel;
      this.description = profile.description;
      this.error = false;
      this.open = true;
    },
  },
  async init() {
    this.loading = true;
    this.profiles = await get('/api/profile');
    if (this.profiles.length > 0) {
      this.selectedProfile = this.profiles[0];
    }
    this.loading = false;
  },
  async selectProfile(profile) {
    this.selectedProfile = profile;
  },
  async createProfile() {
    try {
      const newProfile = await post('/api/profile', {
        linkLabel: this.addProfile.linkLabel,
        linkTarget: this.addProfile.linkTarget,
        description: this.addProfile.description,
      });
      this.profiles.push(newProfile);
      await this.selectProfile(newProfile);
      this.addProfile.open = false;
    } catch (e) {
      this.addProfile.error = true;
    }
  },
  async saveProfile() {
    try {
      await put(`/api/profile/${this.selectedProfile.id}`, {
        linkLabel: this.editProfile.linkLabel,
        linkTarget: this.editProfile.linkTarget,
        description: this.editProfile.description,
      });
      this.selectedProfile.linkLabel = this.editProfile.linkLabel;
      this.selectedProfile.linkTarget = this.editProfile.linkTarget;
      this.selectedProfile.description = this.editProfile.description;
      this.editProfile.open = false;
    } catch (e) {
      this.editProfile.error = true;
    }
  },
  async deleteProfile(profile) {
    if (
      await confirm({
        title: 'Delete profile',
        message: `Are you sure you want to delete the profile ${profile.linkLabel}?`,
        declineLabel: "Don't delete",
        approveLabel: 'Delete profile',
        negative: true,
      })
    ) {
      try {
        await httpDelete(`/api/profile/${this.selectedProfile.id}`);
        this.profiles = this.profiles.filter((p) => p.id !== profile.id);
      } catch (e) {
        alert({
          title: 'Failed to delete profile',
          message: `Failed to delete the profile ${profile.linkLabel}. Please contact support.`,
          negative: true,
          closeLabel: 'Okay',
        });
      }
    }
  },
}));
