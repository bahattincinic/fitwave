import { defineStore } from 'pinia';
import Cookies from 'js-cookie';
import { loginTypeEnum } from '@/services/auth';

const cookieKey = 'app_accessToken';

export const useUserStore = defineStore('user', {
  state: () => ({
    accessToken: '',
    config: {},
    setupCompleted: false,
    loginType: '',
  }),
  actions: {
    logout() {
      this.accessToken = '';
      this.config = {};
      Cookies.remove(cookieKey);
    },
    login(config, accessToken) {
      this.accessToken = accessToken;
      this.config = config;
      Cookies.set(cookieKey, accessToken);
    },
    getAccessToken() {
      return Cookies.get(cookieKey);
    },
    setSetupCompleted(loginType, status) {
      this.setupCompleted = status;
      this.loginType = loginType;
    },
    setConfig(config) {
      this.config = config;
    },
    isAuthenticated() {
      return this.accessToken !== '';
    },
    loginNeeded() {
      return this.loginType === loginTypeEnum.protected;
    },
  },
});
