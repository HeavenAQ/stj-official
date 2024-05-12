import axios from "axios";

interface CreateUserRequest {
  email: string;
  password: string;
  language: "chn";
}

export async function registerUser(email: string, password: string) {
  const data: CreateUserRequest = {
    email,
    password: btoa(password), // password is encoded in base64
    language: "chn",
  };

  const headers = {
    "Content-Type": "application/json",
  };

  return axios.post("http://localhost:8080/api/v1/auth/register", data, {
    headers: headers,
  });
}

export async function loginUser(email: string, password: string) {
  const data = {
    email,
    password: btoa(password),
  };

  const headers = {
    "Content-Type": "application/json",
  };

  return axios.post("http://localhost:8080/api/v1/auth/login", data, {
    headers: headers,
  });
}
