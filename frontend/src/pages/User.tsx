import { useQuery } from '@tanstack/react-query'
import { useState, useEffect } from 'react'
import { redirect } from 'react-router-dom'
import { getUser } from '../api/user'
import PhoneInput from 'react-phone-number-input'
import Loading from '../components/Loading'
import { APIProvider, Map, Marker } from '@vis.gl/react-google-maps'
import { PlaceAutocomplete } from '../components/PlaceAutocomplete'
import MapHandler from '../components/MapHandler'
import { Gender } from '../types'

const googleMapAPIKey = process.env.REACT_APP_GOOGLE_MAP_API_KEY as string

const YearOptions = () => {
  const currentYear = new Date().getFullYear()
  const years = Array.from({ length: 82 }, (_, i) => currentYear - i - 18)
  return years.map(year => (
    <option key={year} value={year}>
      {year}
    </option>
  ))
}

const User = () => {
  // fetch user data
  const { data: user } = useQuery({
    queryKey: ['user'],
    queryFn: () => getUser()
  })

  if (user === undefined) {
    redirect('/login')
  }

  // form control
  const [firstName, setFirstName] = useState<string>(
    user?.data.first_name || ''
  )
  const [lastName, setLastName] = useState<string>(user?.data.last_name || '')
  const [email, setEmail] = useState<string>(user?.data.email || '')
  const [phone, setPhone] = useState<any>(user?.data.phone)
  const [gender, setGender] = useState<string>(user?.data.gender || '')
  const [birthYear, setBirthYear] = useState<string | number>(
    user?.data.birth_year || ''
  )
  const [lineID, setLineID] = useState<string>(user?.data.line_id || '')
  const [spinner, setSpinner] = useState<boolean>(false)

  // google map
  const defaultPosition = { lat: 25.033964, lng: 121.564472 }
  const [position, setPosition] = useState(defaultPosition)
  const [selectedPlace, setSelectedPlace] =
    useState<google.maps.places.PlaceResult | null>(null)

  useEffect(() => {
    const newPosition = {
      lat: selectedPlace?.geometry?.location?.lat() || defaultPosition.lat,
      lng: selectedPlace?.geometry?.location?.lng() || defaultPosition.lng
    }
    setPosition(newPosition)
  }, [selectedPlace])

  return (
    <section className="lg:mt-28 mt-32 bg-zinc-600 rounded-xl w-[90%] sm:w-[80%] md:w-[70%]  lg:w-[60%] max-w-[1000px]  mx-auto md:px-0 px-6 shadow-xl shadow-gray-500 mb-32 pb-9">
      <div className="pt-10 mx-auto w-full md:w-4/5 lg:w-2/3 max-w-[1000px]">
        <h1 className="my-8 text-3xl text-center text-white lg:mb-12 lg:text-3xl tracking-[20px]">
          會員資料
        </h1>
        {spinner && (
          <div className="flex justify-center items-center mx-auto mt-20 w-32 h-32">
            <Loading />
          </div>
        )}
        <form className={`${spinner ? 'hidden' : ''} mt-6 w-full rounded-xl`}>
          <div className="flex flex-wrap -mx-3 mb-6">
            <div className="px-3 mb-6 w-full md:mb-0 md:w-1/2">
              <input
                className="block py-3 px-4 mb-3 w-full leading-tight rounded border appearance-none"
                id="grid-first-name"
                type="text"
                placeholder="姓"
                onChange={ev => setLastName(ev.target.value)}
                value={lastName}
                required
              />
            </div>
            <div className="px-3 w-full md:mb-0 md:w-1/2">
              <input
                className="block py-3 px-4 mb-3 w-full leading-tight rounded border appearance-none"
                id="grid-last-name"
                type="text"
                placeholder="名"
                onChange={ev => setFirstName(ev.target.value)}
                value={firstName}
                required
              />
            </div>
          </div>
          <div className="flex flex-wrap -mx-3 mb-6">
            <div className="px-3 mb-6 w-full md:mb-0 md:w-1/2">
              <select
                className="block py-3 px-4 mb-3 w-full leading-tight rounded border appearance-none"
                id="gender"
                onChange={ev => setGender(ev.target.value)}
                value={gender}
                required
              >
                <option value="">性別</option>
                <option value={Gender.FEMALE}>女性</option>
                <option value={Gender.MALE}>男性</option>
                <option value={Gender.NOT_SPECIFIED}>其他</option>
                <option value={Gender.NOT_DISCLOSED}>不方便透露</option>
              </select>
            </div>
            <div className="px-3 w-full md:mb-0 md:w-1/2">
              <select
                className="block py-3 px-4 mb-3 w-full leading-tight rounded border appearance-none"
                id="birth-year"
                onChange={ev => setBirthYear(ev.target.value)}
                value={birthYear}
                required
              >
                <option value="">出生年份</option>
                {YearOptions()}
              </select>
            </div>
          </div>
          <div className="flex flex-wrap px-3 -mx-3 mb-6">
            <input
              className="block py-3 px-4 mb-3 w-full leading-tight rounded border appearance-none"
              id="line-id"
              type="text"
              placeholder="LINE ID"
              onChange={ev => setLineID(ev.target.value)}
              value={lineID}
            />
          </div>
          <div className="flex flex-wrap px-3 -mx-3 mb-6">
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
          <div className="flex flex-wrap px-3 -mx-3 mb-6">
            <PhoneInput
              className="block py-3 px-4 mb-3 w-full leading-tight bg-white rounded border appearance-none"
              placeholder="請輸入電話號碼"
              defaultCountry="TW"
              value={phone}
              onChange={setPhone}
            />
          </div>
          <APIProvider apiKey={googleMapAPIKey} language="zh-TW">
            <div className="flex flex-wrap px-3 -mx-3 mb-6">
              <PlaceAutocomplete onPlaceSelect={setSelectedPlace} />
            </div>
            <div className="flex flex-wrap px-3 -mx-3 mb-6 rounded-lg">
              <Map
                defaultCenter={position}
                defaultZoom={10}
                className="overflow-hidden w-full rounded-lg h-[400px]"
              >
                <Marker position={position} />
              </Map>
              <MapHandler place={selectedPlace} />
            </div>
          </APIProvider>
          <div className="flex flex-wrap -mx-3">
            <div className="flex justify-center items-end w-full h-auto text-white">
              <button
                type="submit"
                className="inline-flex overflow-hidden relative justify-start items-center py-3 px-6 w-40 font-medium rounded-xl transition-all bg-zinc-500 group"
              >
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
        </form>
      </div>
    </section>
  )
}

export default User
