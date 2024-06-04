import { makeRequest } from './api';

export async function getUserConfig(accessToken) {
  return await makeRequest({
    endpoint: '/config',
    accessToken,
    error: 'Could not fetch user config',
    json: true,
  });
}

export async function checkSetupCompleted() {
  return await makeRequest({
    endpoint: '/config/setup',
    json: true,
    error: 'Could not fetch setup status',
  });
}

export async function completeSetup(config) {
  return await makeRequest({
    method: 'POST',
    endpoint: '/config/setup',
    body: config,
    error: 'Could not save setup',
    json: true,
  });
}

export async function saveUserConfig(accessToken, config) {
  return await makeRequest({
    method: 'PUT',
    endpoint: '/config',
    body: config,
    accessToken,
    json: true,
    error: 'Could not save user config',
  });
}
