import axios from "axios";

import { API_URL_AUTH } from "../constants";

const register = (username, email, password) => {
    return axios.post(`${API_URL_AUTH}/register`, {
        username,
        email,
        password,
    });
};

const login = (username, password) => {
  return axios
    .post(`${API_URL_AUTH}/login`, {
      username,
      password,
    })
    .then((response) => {
      if (response.data.token) {
        localStorage.setItem("user", JSON.stringify(response.data));
      }

      return response.data;
    });
};

const logout = () => {
  localStorage.removeItem("user");
};

export default {
  register,
  login,
  logout,
};
