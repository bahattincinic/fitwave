import { defineStore } from 'pinia';
import Cookies from 'js-cookie';

const cookieKey = 'strava_accessToken';

export const useStravaStore = defineStore('strava', {
  state: () => ({
    accessToken: '',
    user: {},
  }),
  actions: {
    logout() {
      this.user = {};
      this.accessToken = '';
      Cookies.remove(cookieKey);
    },
    login(accessToken, user) {
      this.setUser(accessToken, user);
      Cookies.set(cookieKey, accessToken);
    },
    setUser(accessToken, user) {
      this.accessToken = accessToken;
      this.user = user;
    },
    getAccessToken() {
      return Cookies.get(cookieKey);
    },
    isAuthenticated() {
      return this.accessToken !== '';
    },
  },
});
