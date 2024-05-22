import { useLocation } from 'react-router-dom'
import React from 'react'
import Footer from '../components/Footer'
import Navbar from '../components/Navbar'
import { MdKeyboardArrowRight } from 'react-icons/md'
import { pathnameMap } from '../data/path'

interface MainLayoutProps {
  children: React.ReactNode
}

const showRightArrow = (i: number, locationPath: string[]) => {
  const isNextNumber = !isNaN(Number(locationPath[i + 1]))
  return i !== locationPath.length - 1 && !isNextNumber
}

const MainLayout: React.FC<MainLayoutProps> = ({ children }) => {
  const location = useLocation()
  const locationPath = location.pathname.split('/')
  return (
    <>
      <Navbar />
      <div className="w-[90%] sm:w-[80%] md:w-[75%] lg:w-[70%] max-w-[1500px] mx-auto">
        <div
          className={`inline-flex mt-20  w-[90%] sm:w-[80%] md:w-[75%] lg:w-[60%] h-16 max-w-[1000px] items-center ${location.pathname === '/' ? 'hidden' : ''}`}
        >
          {location.pathname !== '/' &&
            locationPath.map((path, i) => {
              const pathKey = path as keyof typeof pathnameMap
              return (
                <div
                  key={i}
                  className="inline-flex justify-center items-center h-10 text-gray-500 text-md"
                >
                  <a
                    className="px-3 rounded-lg duration-300 hover:text-gray-700 hover:bg-gray-200 text-sky-600"
                    href={`/${path}`}
                  >
                    {pathnameMap[pathKey]}
                  </a>
                  {showRightArrow(i, locationPath) && (
                    <MdKeyboardArrowRight className="text-lg text-gray-500" />
                  )}
                </div>
              )
            })}
        </div>
        {children}
      </div>
      <Footer />
    </>
  )
}
export default MainLayout
