import { API_BASE_URL } from './api';

export async function fetchActivities(page) {
  const response = await fetch(`${API_BASE_URL}/activities?page=${page}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error('Could not fetch activities');
  }

  return await response.json();
}

export async function getActivity(id) {
  const endpoint = `${API_BASE_URL}/activities/${id}`;

  const response = await fetch(endpoint, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error('Could not get activity');
  }

  return await response.json();
}

export async function getActivityGPX(id, accessToken) {
  const endpoint = `${API_BASE_URL}/activities/${id}/gpx`;

  const response = await fetch(endpoint, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${accessToken}`,
    },
  });

  if (!response.ok) {
    throw new Error('Could not get activity');
  }

  return await response.blob();
}

export async function getActivityLaps(id, accessToken) {
  const endpoint = `${API_BASE_URL}/activities/${id}/laps`;

  const response = await fetch(endpoint, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${accessToken}`,
    },
  });

  if (!response.ok) {
    throw new Error('Could not get activity');
  }

  return await response.json();
}
