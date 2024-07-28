import axios, { AxiosResponse, HttpStatusCode } from 'axios'
import { UserDataResponse } from '../types/api/user'

// a curry function that handle cases where the user access token is expired
interface RefreshUserAccessResponse {
  access_token: string
}
export function refreshUserAccess(
  func: (...args: any[]) => Promise<AxiosResponse<UserDataResponse>>
): (...args: any[]) => Promise<AxiosResponse<UserDataResponse>> {
  // Return a function that will refresh the access token if it is invalid
  return async (...args: any[]) => {
    return func(...args).catch(async error => {
      // if the error is not an unauthorized error, reject the promise directly
      if (error.response.status !== HttpStatusCode.Unauthorized) {
        return Promise.reject(error)
      }

      // If the call fails, try refreshing the access token
      const data = {
        refresh_token: sessionStorage.getItem('refresh_token'),
        session_id: sessionStorage.getItem('session_id')
      }

      return axios
        .post('http://localhost:8080/api/v1/auth/refresh', data)
        .then((res: AxiosResponse<RefreshUserAccessResponse>) => {
          // Store the new access token in session storage and call the function again
          sessionStorage.setItem('access_token', res.data.access_token)
          return func(...args)
        })
        .catch(refreshError => {
          return Promise.reject(refreshError)
        })
    })
  }
}
