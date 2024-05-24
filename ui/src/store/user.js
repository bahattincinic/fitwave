import { defineStore } from 'pinia';
import Cookies from "js-cookie";

export const useUserStore = defineStore('user', {
  state: () => ({
    accessToken: '',
    user: {}
  }),
  actions: {
    logout() {
      this.user = {};
      this.accessToken = '';
      Cookies.remove('accessToken');
    },
    login(accessToken, user) {
      this.setUser(accessToken, user);
      Cookies.set('accessToken', accessToken);
    },
    setUser(accessToken, user) {
      this.accessToken = accessToken;
      this.user = user;
    }
  }
});
