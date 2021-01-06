/* eslint-disable no-empty */
const TOKEN_KEY = "easyEtcdTokens";
const ENDPOINTS_KEY = "endpoints";
const DEFAULT_FORMAT_KEY = "default_format";

export function setTokens(key, value) {
  const tokens = getTokens();

  tokens[key] = value;
  let tokenStr = "";

  try {
    tokenStr = JSON.stringify(tokens);
  } catch (error) {}

  return sessionStorage.setItem(TOKEN_KEY, tokenStr);
}

export function getTokens() {
  const tokenStr = sessionStorage.getItem(TOKEN_KEY);

  let tokens = {};
  if (!tokenStr) return tokens;

  try {
    tokens = JSON.parse(tokenStr);
  } catch (error) {}

  return tokens;
}

export function setEndpoint(endpoint) {
  const url = endpoint.url;
  const list = getEndpoint();

  const idx = list.findIndex(el => el.url === url);
  if (idx === -1) {
    list.push(endpoint);
  } else {
    list.splice(idx, 1, endpoint);
  }

  let str = "";
  try {
    str = JSON.stringify(list);
  } catch (error) {}

  return localStorage.setItem(ENDPOINTS_KEY, str);
}

export function removeEndpoint(url) {
  const list = getEndpoint();

  const idx = list.findIndex(el => el.url === url);
  if (idx === -1) return;

  list.splice(idx, 1);
  let str = "";
  try {
    str = JSON.stringify(list);
  } catch (error) {}

  return localStorage.setItem(ENDPOINTS_KEY, str);
}

export function getEndpoint() {
  const str = localStorage.getItem(ENDPOINTS_KEY);

  let list = [];
  if (!str) return list;

  try {
    list = JSON.parse(str);
  } catch (error) {}

  return list;
}

export function setDefaultFormat(endpoint, format) {
  return localStorage.setItem(`${DEFAULT_FORMAT_KEY}_${endpoint}`, format);
}

export function getDefaultFormat(endpoint) {
  return localStorage.getItem(`${DEFAULT_FORMAT_KEY}_${endpoint}`);
}
