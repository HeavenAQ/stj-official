import axios, { AxiosResponse } from 'axios'
import {
  CreateUserRequest,
  LoginUserResponse,
  UserDataResponse
} from '../types/api/user'
import { refreshUserAccess } from './auth'
import { Gender } from '../types/enums'

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

export async function loginUser(email: string, password: string) {
  const data = {
    email,
    password: btoa(password) // backend login requires password to be encoded in base64
  }

  const headers = {
    'Content-Type': 'application/json'
  }

  return axios.post('http://localhost:8080/api/v1/auth/login', data, {
    headers: headers
  }) as Promise<AxiosResponse<LoginUserResponse>>
}

// Get user data
export const getUser = refreshUserAccess(async () => {
  const headers = {
    Authorization: `Bearer ${sessionStorage.getItem('access_token')}`,
    'Content-Type': 'application/json'
  }

  return axios.get('http://localhost:8080/api/v1/users', {
    headers: headers
  }) as Promise<AxiosResponse<UserDataResponse>>
})

// update user data
export const updateUser = refreshUserAccess(async (data: UserDataResponse) => {
  const headers = {
    Authorization: `Bearer ${sessionStorage.getItem('access_token')}`,
    'Content-Type': 'application/json'
  }

  return axios.put('http://localhost:8080/api/v1/users', data, {
    headers: headers
  }) as Promise<AxiosResponse<UserDataResponse>>
})
