function getDomain() {
  const protocol = window.location.protocol;
  const hostname = window.location.hostname;
  const port = window.location.port;
  return `${protocol}//${hostname}${port ? `:${port}` : ''}`;
}

function getApiBaseURL() {
  if (process.env.VUE_APP_API_BASE_URL) {
    return process.env.VUE_APP_API_BASE_URL;
  }
  return `${getDomain()}/api`;
}

class APIError extends Error {
  constructor(message, response) {
    super(message);
    this.response = response;
  }
}

async function makeRequest(payload) {
  const endpoint = `${getApiBaseURL()}${payload.endpoint}`;

  const response = await fetch(endpoint, {
    method: payload.method || 'GET',
    headers: {
      'Content-Type': 'application/json',
      ...(payload.stravaToken
        ? {
            'X-Strava-Authorization': payload.stravaToken,
          }
        : {}),
      ...(payload.accessToken
        ? {
            Authorization: `Bearer ${payload.accessToken}`,
          }
        : {}),
    },
    ...(payload.body ? { body: JSON.stringify(payload.body) } : {}),
  });

  if (!response.ok) {
    const errorMsg = payload.error || 'API Request Failed';
    console.error(errorMsg, response);
    throw new APIError(errorMsg, response);
  }

  if (payload.json) {
    return await response.json();
  }

  if (payload.blob) {
    return await response.blob();
  }

  return response;
}

export { getDomain, getApiBaseURL, makeRequest, APIError };
