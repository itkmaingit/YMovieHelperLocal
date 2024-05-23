import Axios from "axios";

export const clientAxios = Axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_ENDPOINT_ON_CLIENT,
  withCredentials: true,
});

export const serverAxios = Axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_ENDPOINT_ON_SERVER,
  withCredentials: true,
});
