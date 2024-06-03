import { getApiBaseURL } from './api';

export async function fetchActivities(page) {
  const endpoint = `${getApiBaseURL()}/activities?page=${page}`;

  const response = await fetch(endpoint, {
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
  const endpoint = `${getApiBaseURL()}/activities/${id}`;

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
  const endpoint = `${getApiBaseURL()}/strava/activities/${id}/gpx`;

  const response = await fetch(endpoint, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      'X-Strava-Authorization': accessToken,
    },
  });

  if (!response.ok) {
    throw new Error('Could not get activity');
  }

  return await response.blob();
}

export async function getActivityLaps(id, accessToken) {
  const endpoint = `${getApiBaseURL()}/strava/activities/${id}/laps`;

  const response = await fetch(endpoint, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      'X-Strava-Authorization': accessToken,
    },
  });

  if (!response.ok) {
    throw new Error('Could not get activity');
  }

  return await response.json();
}
