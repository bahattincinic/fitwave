import { getDomain, makeRequest } from './api';

export const loginTypeEnum = {
  anonymous: 'anonymous',
  protected: 'protected',
};

export const loginTypes = [
  { name: 'Anonymous', code: loginTypeEnum.anonymous },
  { name: 'Protected', code: loginTypeEnum.protected },
];

export async function getStravaAuthorizationURL(accessToken) {
  const callbackURL = `${getDomain()}/app/strava-login`;

  return await makeRequest({
    endpoint: `/strava/authorization-url?callback_url=${callbackURL}`,
    json: true,
    error: 'Could not fetch user config',
    accessToken,
  });
}

export async function getStravaAccessToken(accessToken, code) {
  return await makeRequest({
    method: 'POST',
    endpoint: `/strava/token`,
    body: { code },
    error: 'Could not fetch access token',
    json: true,
    accessToken,
  });
}

export async function login(data) {
  return await makeRequest({
    method: 'POST',
    endpoint: '/auth/token',
    body: data,
    error: 'Could not login',
    json: true,
  });
}
