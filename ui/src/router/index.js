import { createRouter, createWebHistory } from 'vue-router';
import StravaLoginPage from '@/pages/StravaLoginPage';
import HomePage from '@/pages/HomePage';
import SettingsPage from '@/pages/SettingsPage';
import DashboardDetail from '@/pages/DashboardDetail';
import AthletesPage from '@/pages/AthletesPage';
import ActivitiesPage from '@/pages/ActivitiesPage';
import GearsPage from '@/pages/GearsPage';
import SetupPage from '@/pages/SetupPage';
import LoginPage from '@/pages/LoginPage';
import { useStravaStore } from '@/store/strava';
import { useUserStore } from '@/store/user';
import { getStravaUser } from '@/services/user';
import { checkSetupCompleted, getUserConfig } from '@/services/config';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: HomePage,
  },
  {
    path: '/app/dashboard/:id',
    name: 'DashboardDetail',
    component: DashboardDetail,
  },
  {
    path: '/app/strava-login',
    name: 'StravaLogin',
    component: StravaLoginPage,
  },
  {
    path: '/app/settings',
    name: 'Settings',
    component: SettingsPage,
  },
  {
    path: '/app/activities',
    name: 'Activities',
    component: ActivitiesPage,
  },
  {
    path: '/app/athletes',
    name: 'Athletes',
    component: AthletesPage,
  },
  {
    path: '/app/gears',
    name: 'Gears',
    component: GearsPage,
  },
  {
    path: '/app/setup',
    name: 'Setup',
    component: SetupPage,
  },
  {
    path: '/app/login',
    name: 'Login',
    component: LoginPage,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach(async (to, from, next) => {
  const userStore = useUserStore();
  const stravaStore = useStravaStore();
  const stravaToken = stravaStore.getAccessToken();
  const accessToken = userStore.getAccessToken();

  // Check if the setup is completed
  if (!userStore.setupCompleted) {
    const resp = await checkSetupCompleted();
    userStore.setSetupCompleted(resp.login_type, resp.completed);

    // If setup is not completed and the target page is not 'Setup', redirect to 'Setup' page
    if (!resp.completed && to.name !== 'Setup') {
      return next({ name: 'Setup' });
    }
  }

  // Authentication Check
  const loginNeeded = userStore.loginNeeded();
  if (userStore.setupCompleted && !loginNeeded) {
    const resp = await getUserConfig('');
    userStore.setConfig(resp);
  } else if (!userStore.isAuthenticated() && accessToken) {
    try {
      const resp = await getUserConfig(accessToken);
      userStore.login(resp, accessToken);
    } catch {
      userStore.logout();
    }
  }

  // If login is needed and the user is not authenticated, redirect to 'Login' page
  if (loginNeeded && !userStore.isAuthenticated() && to.name !== 'Login') {
    return next({ name: 'Login' });
  }

  // Strava Token management
  if (!stravaStore.isAuthenticated() && stravaToken) {
    try {
      const resp = await getStravaUser(userStore.accessToken, stravaToken);
      stravaStore.setUser(stravaToken, resp.athlete);
    } catch {
      stravaStore.logout();
    }
  }

  // If setup is completed and the target page is 'Setup', redirect to 'Home' page
  if (userStore.setupCompleted && to.name === 'Setup') {
    return next({ name: 'Home' });
  }

  // If login is not needed and the target page is 'Login', redirect to 'Home' page
  if (!loginNeeded && to.name === 'Login') {
    return next({ name: 'Home' });
  }

  // If the user is authenticated and the target page is 'Login', redirect to 'Home' page
  if (userStore.isAuthenticated() && to.name === 'Login') {
    return next({ name: 'Home' });
  }

  next();
});

export default router;
