import React from 'react'

export default function Footer() {
  return (
    <div className="flex fixed bottom-0 left-0 z-50 justify-evenly items-center px-4 w-full h-24 text-center text-white bg-zinc-800">
      <img
        alt="warning-banner"
        src="/images/misc/warning-banner.webp"
        className="md:h-10 lg:h-14 xl:h-16"
      />
    </div>
  )
}
