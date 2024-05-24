import React from 'react'

interface FormProps {
  title: string
  children: React.ReactNode
}

const FormLayout: React.FC<FormProps> = ({ title, children }) => {
  return (
    <div className="max-h-[700px] w-[90%] sm:w-[80%] md:w-[70%] max-w-[500px] h-[60%] absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 rounded-xl bg-zinc-600 flex flex-col px-10 shadow-xl shadow-gray-500 justify-center animate-fade">
      <h1 className="mb-3 text-3xl text-center text-white">{title}</h1>
      {children}
    </div>
  )
}

export default FormLayout
