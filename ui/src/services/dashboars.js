import { makeRequest } from './api';

export async function fetchDashboards(accessToken) {
  return await makeRequest({
    endpoint: '/dashboards',
    json: true,
    error: 'Could not fetch dashboards',
    accessToken,
  });
}

export async function getDashboard(accessToken, id) {
  return await makeRequest({
    endpoint: `/dashboards/${id}`,
    json: true,
    error: 'Could not get dashboard',
    accessToken,
  });
}

export async function createDashboard(accessToken, data) {
  return await makeRequest({
    endpoint: '/dashboards',
    method: 'POST',
    body: data,
    json: true,
    error: 'Could not create a dashboard',
    accessToken,
  });
}

export async function updateDashboard(accessToken, id, data) {
  return await makeRequest({
    endpoint: `/dashboards/${id}`,
    method: 'PUT',
    body: data,
    json: true,
    error: 'Could not update a dashboard',
    accessToken,
  });
}

export async function deleteDashboard(accessToken, id) {
  return await makeRequest({
    endpoint: `/dashboards/${id}`,
    method: 'DELETE',
    error: 'Could not delete a dashboard',
    accessToken,
  });
}

export async function runDashboard(accessToken, id) {
  return await makeRequest({
    endpoint: `/dashboards/${id}/run`,
    method: 'POST',
    json: true,
    error: 'Could not run dashboard',
    accessToken,
  });
}
