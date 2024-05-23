import { API_BASE_URL } from './api';

export async function getUserConfig() {
  const response = await fetch(`${API_BASE_URL}/user/config`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
  });

  if (!response.ok) {
    throw new Error('Could not fetch user config');
  }

  return await response.json();
}

export async function saveUserConfig(config) {
  const response = await fetch(`${API_BASE_URL}/user/config`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(config)
  });

  if (!response.ok) {
    throw new Error('Could not save user config');
  }

  return await response.json();
}
