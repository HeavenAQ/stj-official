import React, { FC, useEffect, useState } from 'react'
import { YearOptions } from '../utils/options'
import PhoneInput from 'react-phone-number-input'
import 'react-phone-number-input/style.css'
import { useUserQuery } from '../utils/query'
import { PlaceAutocomplete } from './PlaceAutocomplete'
import { Map, Marker } from '@vis.gl/react-google-maps'
import MapHandler from './MapHandler'
import useGoogleMap from '../utils/hooks'
import GoogleMapInput from './GoogleMapInput'

const DeliveryInfo: FC = () => {
  // fetch user data
  const { data: user } = useUserQuery()

  // form control
  const [firstName, setFirstName] = useState<string>('')
  const [lastName, setLastName] = useState<string>('')
  const [email, setEmail] = useState<string>('')
  const [phone, setPhone] = useState<any>('')
  const [birthYear, setBirthYear] = useState<number>(0)

  // google map
  const {
    backupAddress,
    setBackupAddress,
    markerPosition,
    selectedPlace,
    setSelectedPlace
  } = useGoogleMap(25.033964, 121.564472)

  return (
    <section className=" bg-zinc-600 rounded-xl mx-auto md:px-0 px-6 mb-32 pb-9 animate-fade w-5/6">
      <div className="pt-10 mx-auto w-full md:w-4/5 max-w-[1000px]">
        <h1 className="my-8 text-3xl text-center text-white lg:mb-12 lg:text-3xl tracking-[20px]">
          聯絡資訊
        </h1>
        <div className={`mt-6 w-full rounded-xl`}>
          <div className="flex flex-wrap -mx-3 mb-3">
            <div className="px-3 mb-6 space-y-2 w-full md:mb-0 md:w-1/2">
              <label className="font-medium text-slate-200">姓氏</label>
              <input
                className="block py-3 px-4 w-full leading-tight rounded border appearance-none"
                id="grid-last-name"
                placeholder="姓氏"
                type="text"
                onChange={ev => setFirstName(ev.target.value)}
                value={firstName}
                required
              />
            </div>
            <div className="px-3 space-y-2 w-full md:mb-0 md:w-1/2">
              <label className="font-medium text-slate-200">名字</label>
              <input
                className="block py-3 px-4 mb-3 w-full leading-tight rounded border appearance-none"
                id="grid-first-name"
                type="text"
                placeholder="名字"
                onChange={ev => setLastName(ev.target.value)}
                value={lastName}
                required
              />
            </div>
          </div>
          <div className="flex flex-wrap px-3 -mx-3 mb-3 space-y-2">
            <label className="font-medium text-slate-200">
              出生年份 <span className="text-red-400">(未滿18歲禁止飲酒)</span>
            </label>
            <select
              className="block py-3 px-4 w-full leading-tight rounded border appearance-none"
              id="birth-year"
              onChange={ev => setBirthYear(Number(ev.target.value))}
              value={birthYear}
              required
            >
              <option value={0}></option>
              <YearOptions />
            </select>
          </div>
          <div className="flex flex-wrap px-3 -mx-3 mb-3 space-y-2">
            <label className="font-medium text-slate-200">信箱</label>
            <input
              className="block py-3 px-4 mb-3 w-full leading-tight rounded border appearance-none"
              id="email"
              type="email"
              placeholder="example@mail.com"
              onChange={ev => setEmail(ev.target.value)}
              value={email}
              required
            />
          </div>
          <div className="flex flex-wrap px-3 -mx-3 mb-3 space-y-2">
            <label className="font-medium text-slate-200">電話</label>
            <PhoneInput
              className="block py-3 px-4 mb-3 w-full leading-tight bg-white rounded border appearance-none"
              defaultCountry="TW"
              value={phone}
              onChange={ev => setPhone(ev)}
            />
          </div>
          <GoogleMapInput
            backupAddress={backupAddress}
            setBackupAddress={setBackupAddress}
            markerPosition={markerPosition}
            selectedPlace={selectedPlace}
            setSelectedPlace={setSelectedPlace}
          />
          <div className="flex flex-wrap -mx-3">
            <div className="flex justify-center items-end w-full h-auto text-white">
              <button className="inline-flex overflow-hidden relative justify-start items-center py-3 px-6 w-40 font-medium rounded-xl transition-all bg-zinc-500 group">
                <span className="inline-block absolute top-0 right-0 w-4 h-4 rounded transition-all duration-500 ease-in-out group-hover:-mt-4 group-hover:-mr-4 bg-zinc-700">
                  <span className="absolute top-0 right-0 w-5 h-5 bg-white rotate-45 translate-x-1/2 -translate-y-1/2"></span>
                </span>
                <span className="absolute bottom-0 left-0 w-full h-full rounded-xl transition-all duration-500 ease-in-out delay-200 -translate-x-full translate-y-full group-hover:mb-12 group-hover:translate-x-0 bg-zinc-400"></span>
                <span className="relative w-full text-center text-white transition-colors duration-200 ease-in-out group-hover:text-white">
                  送出
                </span>
              </button>
            </div>
          </div>
        </div>
      </div>
    </section>
  )
}

export default DeliveryInfo
