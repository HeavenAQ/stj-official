import axios from "axios";

interface CreateUserRequest {
  email: string;
  password: string;
  language: "chn";
}

export async function createUser(email: string, password: string) {
  const data: CreateUserRequest = {
    email,
    password,
    language: "chn",
  };

  const headers = {
    "Content-Type": "application/json",
  };

  return axios.post("http://localhost:8080/api/v1/users", data, {
    headers: headers,
  });
}
