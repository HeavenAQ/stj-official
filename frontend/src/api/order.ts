import axios from 'axios'
import { CreateOrderRequest, CreateOrderResponse } from '../types/api/order'

export const createOrder = async (order: CreateOrderRequest) => {
  const headers = {
    'Content-TYpe': 'application/json',
    Authorization: `Bearer ${sessionStorage.getItem('access_token')}`
  }

  return axios.post('http://localhost:8080/api/v1/orders', order, {
    headers: headers
  }) as Promise<CreateOrderResponse>
}
