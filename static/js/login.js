import { post } from '../lib/jinya-http.js';

import QrcodeStyling from '../lib/qrcode-styling.js';

function base32Encode(bytes) {
  const base32Chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ234567';
  let result = '';
  let buffer = 0;
  let bitsLeft = 0;
  for (const byte of bytes) {
    buffer = (buffer << 8) | byte;
    bitsLeft += 8;

    while (bitsLeft >= 5) {
      result += base32Chars[(buffer >>> (bitsLeft - 5)) & 31];
      bitsLeft -= 5;
    }
  }

  if (bitsLeft > 0) {
    result += base32Chars[(buffer << (5 - bitsLeft)) & 31];
  }

  return result;
}

Alpine.data('loginData', () => ({
  username: '',
  password: '',
  twoFactorCode: '',
  twoFactorNeeded: false,
  firstAuthStepDone: false,
  setupTwoFactor: false,
  loginError: false,
  twoFactorError: false,
  twoFactorSecret: '',
  twoFactorEnableFailed: false,
  twoFactorQrCode: '',
  async login() {
    try {
      const res = await post('/api/authentication/login', {
        username: this.username,
        password: this.password,
      });
      this.firstAuthStepDone = true;
      if (res.twoFactorSetupNeeded) {
        this.setupTwoFactor = true;
        this.twoFactorNeeded = false;
        const bytes = new Uint8Array(20);
        crypto.getRandomValues(bytes);
        this.twoFactorSecret = base32Encode(bytes);

        const params = new URLSearchParams({
          secret: this.twoFactorSecret,
          issuer: 'Reemt Rühmkorfs Website',
          algorithm: 'SHA1',
          digits: '6',
          period: '30',
        });
        const otpUrl = `otpauth://totp/${encodeURIComponent(`${this.username} – ruehmkorf.com`)}?${params.toString()}`;
        const qrCode = new QrcodeStyling({
          type: 'svg',
          shape: 'square',
          width: 256,
          height: 256,
          data: otpUrl,
          qrOptions: { errorCorrectionLevel: 'Q' },
          dotsOptions: { type: 'extra-rounded', color: 'var(--white)' },
          backgroundOptions: {
            round: 0.1,
            color: 'var(--primary-color)',
          },
          cornersSquareOptions: { type: 'extra-rounded', color: 'var(--white)' },
        });
        await qrCode._svgDrawingPromise;
        this.twoFactorQrCode = qrCode._svg.outerHTML;
      } else {
        this.twoFactorNeeded = true;
      }
      this.loginError = false;
    } catch (e) {
      this.loginError = true;
    }
  },
  async twoFactorLogin() {
    try {
      const res = await post('/api/authentication/login', {
        username: this.username,
        password: this.password,
        twoFactorCode: this.twoFactorCode,
      });
      if (res.token) {
        this.$router.navigate('/profiles');
      }
    } catch (e) {
      this.twoFactorError = true;
    }
  },
  async enableTwoFactor() {
    try {
      const res = await post('/api/authentication/2fa', {
        secret: this.twoFactorSecret,
        code: this.twoFactorCode,
        username: this.username,
        password: this.password,
      });
      this.firstAuthStepDone = false;
      this.setupTwoFactor = false;
      this.twoFactorNeeded = false;
      this.twoFactorCode = '';
      this.twoFactorSecret = '';
    } catch (e) {
      this.twoFactorEnableFailed = true;
    }
  },
}));
