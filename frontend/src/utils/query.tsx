import { useQuery } from '@tanstack/react-query'
import { getUser } from '../api/user'
import { getProduct, listProducts } from '../api/product'

export const useUserQuery = () => {
  return useQuery({
    queryKey: ['user'],
    queryFn: getUser
  })
}

export const useProductsQuery = (limit: number, offset: number) => {
  return useQuery({
    queryKey: ['sakes', limit, offset],
    queryFn: () => listProducts(limit, offset)
  })
}

export const useProductQuery = (productID: string) => {
  return useQuery({
    queryKey: ['sake', productID],
    queryFn: () => getProduct(productID)
  })
}
