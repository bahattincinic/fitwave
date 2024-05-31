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

export { getDomain, getApiBaseURL };
