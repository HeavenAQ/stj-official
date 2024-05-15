import axios, { AxiosResponse, HttpStatusCode } from 'axios'
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
  refresh_token: string
  access_expires_in: number
  refresh_expires_in: number
  session_id: string
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

// Get user data
export const getUser = refreshUserAccess(async () => {
  const headers = {
    Authorization: `Bearer ${sessionStorage.getItem('access_token')}`,
    'Content-Type': 'application/json'
  }

  return axios.get('http://localhost:8080/api/v1/users', {
    headers: headers
  }) as Promise<AxiosResponse<UserData>>
})

// update user data
export const updateUser = refreshUserAccess(async (data: UserData) => {
  const headers = {
    Authorization: `Bearer ${sessionStorage.getItem('access_token')}`,
    'Content-Type': 'application/json'
  }

  return axios.put('http://localhost:8080/api/v1/users', data, {
    headers: headers
  }) as Promise<AxiosResponse<UserData>>
})

// a curry function that handle cases where the user access token is expired
interface RefreshUserAccessResponse {
  access_token: string
}
export function refreshUserAccess(
  func: (...args: any[]) => Promise<AxiosResponse<UserData>>
): (...args: any[]) => Promise<AxiosResponse<UserData>> {
  // Return a function that will refresh the access token if it is invalid
  return async (...args: any[]) => {
    return func(...args).catch(async error => {
      // if the error is not an unauthorized error, reject the promise directly
      if (error.response.status !== HttpStatusCode.Unauthorized) {
        return Promise.reject(error)
      }

      // If the call fails, try to refresh the access token
      const data = {
        refresh_token: sessionStorage.getItem('refresh_token'),
        session_id: sessionStorage.getItem('session_id')
      }

      return axios
        .post('http://localhost:8080/api/v1/auth/refresh', data)
        .then((res: AxiosResponse<RefreshUserAccessResponse>) => {
          // Store the new access token in session storage and call the function again
          sessionStorage.setItem('access_token', res.data.access_token)
          return func()
        })
        .catch(refreshError => {
          return Promise.reject(refreshError)
        })
    })
  }
}
