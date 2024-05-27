import { API_BASE_URL } from './api';

export async function getUserConfig() {
  const response = await fetch(`${API_BASE_URL}/user/config`, {
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

export async function getUserMe(accessToken) {
  const response = await fetch(`${API_BASE_URL}/user/me`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${accessToken}`,
    },
  });

  if (!response.ok) {
    throw new Error('Could not fetch user');
  }

  return await response.json();
}

export async function getTaskDetail(id) {
  const endpoint = `${API_BASE_URL}/user/task/${id}`;

  const response = await fetch(endpoint, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error('Could not fetch task detail');
  }

  return await response.json();
}

export async function triggerSync(accessToken) {
  const response = await fetch(`${API_BASE_URL}/user/sync`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${accessToken}`,
    },
  });

  if (!response.ok) {
    throw new Error('Could not fetch task detail');
  }

  return await response.json();
}

export async function saveUserConfig(config) {
  const response = await fetch(`${API_BASE_URL}/user/config`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(config),
  });

  if (!response.ok) {
    throw new Error('Could not save user config');
  }

  return await response.json();
}

export async function runQuery(query) {
  const response = await fetch(`${API_BASE_URL}/user/query`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(query),
  });

  if (!response.ok) {
    throw new Error('Could not run query');
  }

  return await response.json();
}
