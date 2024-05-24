import { useEffect, useState } from 'react'
import {
  HiOutlineUserCircle,
  HiOutlineShoppingCart,
  HiBars3BottomLeft
} from 'react-icons/hi2'
import { useCart } from '../contexts/CartContext'
import { useUserQuery } from '../utils/query'

export default function Navbar() {
  const cartContext = useCart()
  const [itemTotal, setItemTotal] = useState(cartContext?.cart.length || 0)
  const { data: user, isError, error } = useUserQuery()
  if (isError) {
    console.log(error)
  }

  useEffect(() => {
    setItemTotal(cartContext?.cart.length || 0)
    console.log(itemTotal)
  }, [cartContext?.cart.length, itemTotal])

  return (
    <div className="flex fixed top-0 left-1/2 z-50 justify-evenly items-center px-3 w-full h-20 bg-white -translate-x-1/2">
      <div className="flex relative justify-between items-center h-full max-w-[1500px] w-[90%] sm:w-[80%] md:w-[75%] lg:w-[70%]">
        <div className="inline-flex justify-evenly items-center space-x-5">
          <HiBars3BottomLeft className="text-3xl text-zinc-900" />
          {user === undefined ? (
            <a href="/login">
              <HiOutlineUserCircle className="text-3xl cursor-pointer text-zinc-900" />
            </a>
          ) : (
            <a
              href="/user"
              className="inline-flex items-center space-x-3 w-auto"
            >
              <HiOutlineUserCircle className="text-3xl cursor-pointer text-zinc-900" />
              <p className="hidden md:inline-block">{user.data.email}</p>
            </a>
          )}
        </div>
        <a href="/">
          <img
            className="absolute top-1/2 left-1/2 h-full -translate-x-1/2 -translate-y-1/2"
            alt="logo"
            src="/images/misc/logo.svg"
          />
        </a>
        <a href="/order">
          <div className="relative">
            <span
              className={`absolute top-0 -right-1 inline-flex items-center justify-center w-4 h-4 text-xs font-bold text-white bg-red-500 rounded-full ${itemTotal === 0 ? 'hidden' : ''}`}
            >
              {itemTotal}
            </span>
            <HiOutlineShoppingCart className="text-3xl font-thin text-zinc-900" />
          </div>
        </a>
      </div>
    </div>
  )
}
