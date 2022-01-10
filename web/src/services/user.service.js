import axios from "axios";
import authHeader from "./auth-header";

import { API_URL_TEST } from "../constants";

const getPublicContent = () => {
  return axios.get(`${API_URL_TEST}/all`);
};

const getUserBoard = () => {
  return axios.get(`${API_URL_TEST}/user`, { headers: authHeader() });
};

const getModeratorBoard = () => {
  return axios.get(`${API_URL_TEST}/moderator`, { headers: authHeader() });
};

const getAdminBoard = () => {
  return axios.get(`${API_URL_TEST}/admin`, { headers: authHeader() });
};

export default {
  getPublicContent,
  getUserBoard,
  getModeratorBoard,
  getAdminBoard,
};