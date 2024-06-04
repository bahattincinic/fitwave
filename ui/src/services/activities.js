import { makeRequest } from './api';

export async function fetchActivities(accessToken, page) {
  return await makeRequest({
    endpoint: `/activities?page=${page}`,
    accessToken,
    error: 'Could not fetch activities',
    json: true,
  });
}

export async function getActivity(accessToken, id) {
  return await makeRequest({
    endpoint: `/activities/${id}`,
    accessToken,
    error: 'Could not get activity',
    json: true,
  });
}

export async function getActivityGPX(id, stravaToken, accessToken) {
  return await makeRequest({
    endpoint: `/strava/activities/${id}/gpx`,
    stravaToken,
    accessToken,
    error: 'Could not get activity',
    blob: true,
  });
}

export async function getActivityLaps(id, stravaToken, accessToken) {
  return await makeRequest({
    endpoint: `/strava/activities/${id}/laps`,
    stravaToken,
    accessToken,
    error: 'Could not get activity laps',
    json: true,
  });
}
