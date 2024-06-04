import { makeRequest } from './api';

export const taskStatusEnum = {
  pending: 'pending',
  success: 'success',
  error: 'error',
  archived: 'archived',
};

export async function getStravaUser(accessToken, stravaToken) {
  return await makeRequest({
    endpoint: '/strava/me',
    json: true,
    error: 'Could not fetch user',
    stravaToken,
    accessToken,
  });
}

export async function getTaskDetail(accessToken, id) {
  return await makeRequest({
    endpoint: `/user/task/${id}`,
    json: true,
    error: 'Could not fetch task detail',
    accessToken,
  });
}

export async function triggerSync(accessToken, stravaToken) {
  return await makeRequest({
    method: 'POST',
    endpoint: '/strava/sync',
    json: true,
    error: 'Could not trigger a sync task',
    accessToken,
    stravaToken,
  });
}

export async function runQuery(accessToken, query) {
  return await makeRequest({
    endpoint: '/user/query',
    method: 'POST',
    body: query,
    json: true,
    error: 'Could not run query',
    accessToken,
  });
}

export async function waitAsyncTask(accessToken, task) {
  const delay = async (ms) => {
    return new Promise((resolve) => setTimeout(resolve, ms));
  };

  let taskStatus = task.status;
  while (![taskStatusEnum.success, taskStatusEnum.error].includes(taskStatus)) {
    await delay(1000);
    task = await getTaskDetail(accessToken, task.id);
    taskStatus = task.status;
  }

  if (taskStatus === taskStatusEnum.error) {
    throw new Error('Async Task Failed');
  }

  return task;
}
