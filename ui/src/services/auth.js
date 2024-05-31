import { getDomain, getApiBaseURL } from './api';

export async function getAuthorizationURL() {
  const callbackURL = `${getDomain()}/app/login`;
  const endpoint = `${getApiBaseURL()}/auth/authorization-url?callback_url=${callbackURL}`;

  const response = await fetch(endpoint, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error('Could not fetch user config');
  }

  return await response.json();
}

export async function getAccessToken(code) {
  const endpoint = `${getApiBaseURL()}/auth/token`;

  const response = await fetch(endpoint, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ code }),
  });

  if (!response.ok) {
    throw new Error('Could not fetch access token');
  }

  return await response.json();
}
