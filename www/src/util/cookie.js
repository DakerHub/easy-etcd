import Cookies from "js-cookie";
const tokenKey = "token";

export function getToken() {
  return Cookies.get(tokenKey);
}
