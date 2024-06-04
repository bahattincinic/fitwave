import { makeRequest } from './api';

export async function fetchAthletes(accessToken, page) {
  return await makeRequest({
    endpoint: `/athletes?page=${page}`,
    json: true,
    error: 'Could not fetch athletes',
    accessToken,
  });
}

export async function getAthlete(accessToken, id) {
  return await makeRequest({
    endpoint: `/athletes/${id}`,
    json: true,
    error: 'Could not get athlete',
    accessToken,
  });
}
