import axios, { AxiosResponse } from 'axios'
import { Gender } from '../types'

interface CreateUserRequest {
  email: string
  password: string
  language: 'chn'
  gender: Gender
}

export async function registerUser(email: string, password: string) {
  const data: CreateUserRequest = {
    email,
    password: btoa(password), // password is encoded in base64
    language: 'chn',
    gender: Gender.NOT_DISCLOSED
  }

  const headers = {
    'Content-Type': 'application/json'
  }

  return axios.post('http://localhost:8080/api/v1/auth/register', data, {
    headers: headers
  })
}

export type UserData = {
  email: string
  phone: string
  first_name: string
  last_name: string
  language: string
  address: string
  line_id: string
  gender: Gender
  birth_year: number
  latitude: number
  longitude: number
}

interface LoginUserResponse {
  access_token: string
  user: UserData
}

export async function loginUser(email: string, password: string) {
  const data = {
    email,
    password: btoa(password)
  }

  const headers = {
    'Content-Type': 'application/json'
  }

  return axios.post('http://localhost:8080/api/v1/auth/login', data, {
    headers: headers
  }) as Promise<AxiosResponse<LoginUserResponse>>
}

export async function getUser() {
  const headers = {
    Authorization: `Bearer ${sessionStorage.getItem('access_token')}`,
    'Content-Type': 'application/json'
  }

  return axios.get('http://localhost:8080/api/v1/users', {
    headers: headers
  }) as Promise<AxiosResponse<UserData>>
}

export async function updateUser(data: UserData) {
  const headers = {
    Authorization: `Bearer ${sessionStorage.getItem('access_token')}`,
    'Content-Type': 'application/json'
  }

  return axios.put('http://localhost:8080/api/v1/users', data, {
    headers: headers
  }) as Promise<AxiosResponse<UserData>>
}
