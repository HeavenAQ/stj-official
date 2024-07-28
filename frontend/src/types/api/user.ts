import { Gender } from '../enums'

export type UserDataResponse = {
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

export interface LoginUserResponse {
  access_token: string
  refresh_token: string
  access_expires_in: number
  refresh_expires_in: number
  session_id: string
  user: UserDataResponse
}

export interface CreateUserRequest {
  email: string
  password: string
  language: 'chn'
  gender: Gender
}
