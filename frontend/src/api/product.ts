import axios, { AxiosResponse } from 'axios'
import { ProductWithInfo } from '../types/api/product'

export const listProducts = (limit: number, offset: number) => {
  const headers = {
    'Content-Type': 'application/json'
  }

  return axios.get('http://localhost:8080/api/v1/products', {
    headers: headers,
    params: {
      limit: limit,
      offset: offset,
      language: 'chn'
    }
  }) as Promise<AxiosResponse<ProductWithInfo[]>>
}

export const getProduct = (id: string) => {
  const headers = {
    'Content-Type': 'application/json'
  }

  return axios.get(`http://localhost:8080/api/v1/products/${id}`, {
    headers: headers,
    params: { language: 'chn' }
  }) as Promise<AxiosResponse<ProductWithInfo>>
}
