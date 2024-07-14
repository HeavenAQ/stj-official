import { FC } from 'react'
import { YearOptions } from '../utils/options'
import PhoneInput from 'react-phone-number-input'
import 'react-phone-number-input/style.css'
import GoogleMapInput from './GoogleMapInput'

interface DeliveryInfoProps {
  firstName: string
  setFirstName: (firstName: string) => void
  lastName: string
  setLastName: (lastName: string) => void
  email: string
  setEmail: (email: string) => void
  phone: any
  setPhone: (phone: any) => void
  birthYear: number
  setBirthYear: (birthYear: number) => void
  sameAsUser: boolean
  setSameAsUser: (sameAsUser: boolean) => void
  backupAddress: string
  setBackupAddress: (address: string) => void
  markerPosition: { lat: number; lng: number }
  setMarkerPosition: (position: { lat: number; lng: number }) => void
  selectedPlace: google.maps.places.PlaceResult | null
  setSelectedPlace: (place: google.maps.places.PlaceResult | null) => void
}

const DeliveryInfo: FC<DeliveryInfoProps> = ({
  firstName,
  setFirstName,
  lastName,
  setLastName,
  email,
  setEmail,
  phone,
  setPhone,
  birthYear,
  setBirthYear,
  sameAsUser,
  setSameAsUser,
  backupAddress,
  setBackupAddress,
  markerPosition,
  selectedPlace,
  setSelectedPlace
}) => {
  return (
    <section className=" bg-zinc-600 rounded-xl mx-auto md:px-0 px-6 mb-14 pb-9 animate-fade w-5/6">
      <div className="pt-10 mx-auto w-full md:w-4/5 max-w-[1000px]">
        <h1 className="my-8 text-3xl text-center text-white lg:mb-12 lg:text-3xl tracking-[20px] font-bold">
          聯絡資訊
        </h1>

        <label
          htmlFor="toggle"
          className="inline-flex cursor-pointer text-slate-200 space-x-2 w-32"
        >
          <input
            type="checkbox"
            id="toggle"
            className="cursor-pointer"
            onClick={() => setSameAsUser(!sameAsUser)}
          />
          <p>同會員資訊</p>
        </label>
        <div className={`mt-6 w-full rounded-xl`}>
          <div className="flex flex-wrap -mx-3 mb-3">
            <div className="px-3 mb-6 space-y-2 w-full md:mb-0 md:w-1/2">
              <label
                className="font-medium text-slate-200"
                htmlFor="grid-last-name"
              >
                姓氏
              </label>
              <input
                className="block py-3 px-4 w-full leading-tight rounded border appearance-none"
                id="grid-last-name"
                placeholder="姓氏"
                type="text"
                onChange={ev => setLastName(ev.target.value)}
                value={lastName}
                required
              />
            </div>
            <div className="px-3 space-y-2 w-full md:mb-0 md:w-1/2">
              <label
                className="font-medium text-slate-200"
                htmlFor="grid-first-name"
              >
                名字
              </label>
              <input
                className="block py-3 px-4 mb-3 w-full leading-tight rounded border appearance-none"
                id="grid-first-name"
                type="text"
                onChange={ev => setFirstName(ev.target.value)}
                placeholder="名字"
                value={firstName}
                required
              />
            </div>
          </div>
          <div className="flex flex-wrap px-3 -mx-3 mb-3 space-y-2">
            <label className="font-medium text-slate-200" htmlFor="birth-year">
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
            <label className="font-medium text-slate-200" htmlFor="email">
              信箱
            </label>
            <input
              className="block py-3 px-4 mb-3 w-full leading-tight rounded border appearance-none"
              id="email"
              type="email"
              placeholder="example@mail.com"
              onChange={ev => setEmail(ev.target.value)}
              value={email}
              autoComplete="email"
              required
            />
          </div>
          <div className="flex flex-wrap px-3 -mx-3 mb-3 space-y-2">
            <label className="font-medium text-slate-200" htmlFor="phone">
              電話
            </label>
            <PhoneInput
              id="phone"
              className="py-3 px-4 mb-3 w-full leading-tight bg-white rounded border appearance-none"
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
        </div>
      </div>
    </section>
  )
}

export default DeliveryInfo
