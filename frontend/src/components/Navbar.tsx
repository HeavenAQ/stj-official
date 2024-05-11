import {
  HiOutlineUserCircle,
  HiOutlineShoppingCart,
  HiBars3BottomLeft,
} from "react-icons/hi2";

export default function Navbar() {
  return (
    <div className="flex fixed top-0 left-1/2 z-50 justify-evenly items-center px-3 w-full h-20 bg-white -translate-x-1/2 max-w-[1000px]">
      <div className="inline-flex justify-evenly items-center space-x-2">
        <HiBars3BottomLeft className="text-3xl text-zinc-900" />
        <HiOutlineUserCircle className="text-3xl cursor-pointer text-zinc-900" />
      </div>
      <img
        className="pr-8 mx-auto h-full"
        src="/images/misc/logo.svg"
        alt="logo"
      />

      <HiOutlineShoppingCart className="text-3xl font-thin text-zinc-900" />
    </div>
  );
}
