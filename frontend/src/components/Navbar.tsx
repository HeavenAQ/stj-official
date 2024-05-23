import {
  HiOutlineUserCircle,
  HiOutlineShoppingCart,
  HiBars3BottomLeft
} from 'react-icons/hi2'
import { useUserQuery } from '../utils/query'

export default function Navbar() {
  const { data: user, isError, error } = useUserQuery()
  if (isError) {
    console.log(error)
  }

  return (
    <div className="flex fixed top-0 left-1/2 z-50 justify-evenly items-center px-3 w-full h-20 bg-white -translate-x-1/2">
      <div className="flex relative justify-between items-center w-full h-full max-w-[1000px]">
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
            src="/images/misc/logo.svg"
            alt="logo"
          />
        </a>
        <a href="/order">
          <HiOutlineShoppingCart className="text-3xl font-thin text-zinc-900" />
        </a>
      </div>
    </div>
  )
}
