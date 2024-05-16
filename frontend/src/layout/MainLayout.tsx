import React from 'react'
import { useLocation } from 'react-router-dom'
import Footer from '../components/Footer'
import Navbar from '../components/Navbar'
import { MdKeyboardArrowRight } from 'react-icons/md'
import { pathnameMap } from '../data/path'

interface MainLayoutProps {
  children: React.ReactNode
}

const MainLayout: React.FC<MainLayoutProps> = ({ children }) => {
  const location = useLocation()
  const locationPath = location.pathname.split('/')
  return (
    <>
      <Navbar />
      <div className="inline-flex mt-20 ml-[50%] -translate-x-1/2 w-[90%] sm:w-[80%] md:w-[75%] lg:w-[60%] h-16 max-w-[1000px] items-center space-x-1">
        {location.pathname !== '/' &&
          locationPath.map((path, i) => {
            const pathKey = path as keyof typeof pathnameMap
            return (
              <div
                key={i}
                className="inline-flex justify-center items-center space-x-1 h-10 text-gray-500 text-md"
              >
                <a className="text-sky-600" href={`/${path}`}>
                  {pathnameMap[pathKey]}
                </a>
                {i !== locationPath.length - 1 && (
                  <MdKeyboardArrowRight className="text-lg text-gray-500" />
                )}
              </div>
            )
          })}
      </div>
      {children}
      <Footer />
    </>
  )
}
export default MainLayout
