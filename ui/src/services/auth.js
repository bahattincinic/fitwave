import { API_BASE_URL } from './api';

function getDomain() {
  const protocol = window.location.protocol;
  const hostname = window.location.hostname;
  const port = window.location.port;
  return `${protocol}//${hostname}${port ? `:${port}` : ''}`;
}

export async function getAuthorizationURL() {
  const callbackURL = `${getDomain()}/app/login`;
  const endpoint = `${API_BASE_URL}/auth/authorization-url?callback_url=${callbackURL}`;

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
  const endpoint = `${API_BASE_URL}/auth/token`;
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
