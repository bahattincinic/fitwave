import { makeRequest } from './api';

export async function fetchGears(accessToken, page) {
  return await makeRequest({
    endpoint: `/gears?page=${page}`,
    json: true,
    error: 'Could not fetch gears',
    accessToken,
  });
}

export async function getGear(accessToken, id) {
  return await makeRequest({
    endpoint: `/gears/${id}`,
    json: true,
    error: 'Could not get gear',
    accessToken,
  });
}
