import React, {
  createContext,
  FC,
  useContext,
  useEffect,
  useState
} from 'react'
import { toast } from 'react-hot-toast'

import { SakeInfo } from '../data/sakes'

export interface CartItem {
  product: SakeInfo
  quantity: number
}

interface CartContextType {
  cart: CartItem[]
  addToCart: (item: CartItem) => void
  updateCart: (itemID: string, quantity: number) => void
  removeFromCart: (itemID: string) => void
  emptyCart: () => void
}

export const CartContext = createContext<CartContextType | null>(null)

export const useCart = () => {
  const context = useContext(CartContext)
  if (!context) {
    throw new Error('useCart must be used withint CartProvider')
  }
  return context
}

interface CartProviderProps {
  children: React.ReactNode
}

export const CartProvider: FC<CartProviderProps> = ({ children }) => {
  const [cart, setCart] = useState<CartItem[]>(() => {
    const currentCartItems = localStorage.getItem('cart')
    return currentCartItems ? JSON.parse(currentCartItems) : []
  })

  useEffect(() => {
    localStorage.setItem('cart', JSON.stringify(cart))
  }, [cart])

  const updateCart = (itemID: string, quantity: number) => {
    const existingItem = cart.find(cart => cart.product.id === itemID)
    existingItem &&
      setCart(prevCart => {
        const newCart = prevCart.map(cart => {
          if (cart.product.id === itemID) {
            cart.quantity = quantity
          }
          return cart
        })
        return newCart
      })
  }

  const addToCart = (item: CartItem) => {
    const existingItem = cart.find(cart => cart.product.id === item.product.id)
    if (item.quantity <= 0) {
      toast.error('商品數量必須大於 0 件')
      return
    } else if (existingItem && existingItem.quantity === item.quantity) {
      toast.error('商品已在詢價單中')
      return
    } else if (existingItem) {
      updateCart(item.product.id, item.quantity)
      toast.success('商品數量已更新')
      return
    }
    setCart(prevCart => [...prevCart, item])
    toast.success('商品已加入詢價單')
  }

  const removeFromCart = (itemID: string) => {
    setCart(prevCart => prevCart.filter(cart => cart.product.id !== itemID))
  }

  const emptyCart = () => {
    setCart([])
  }

  return (
    <CartContext.Provider
      value={{
        cart,
        addToCart,
        updateCart,
        removeFromCart,
        emptyCart
      }}
    >
      {children}
    </CartContext.Provider>
  )
}
