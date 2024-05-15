import { useState, useEffect, useMemo } from 'react'
import { redirect } from 'react-router-dom'
import { updateUser, UserData } from '../api/user'
import PhoneInput from 'react-phone-number-input'
import Loading from '../components/Loading'
import { Map, Marker } from '@vis.gl/react-google-maps'
import { PlaceAutocomplete } from '../components/PlaceAutocomplete'
import MapHandler from '../components/MapHandler'
import { Gender } from '../types'
import { useUserQuery } from '../utils/query'
import { QueryClient, useMutation, useQueryClient } from '@tanstack/react-query'
import { toast } from 'react-hot-toast'
import { AxiosError, HttpStatusCode } from 'axios'

const YearOptions = () => {
  const currentYear = new Date().getFullYear()
  const years = Array.from({ length: 82 }, (_, i) => currentYear - i - 18)
  return years.map(year => (
    <option key={year} value={year}>
      {year}
    </option>
  ))
}

const useUpdateUserMutation = (queryClient: QueryClient) => {
  return useMutation({
    mutationFn: async (data: UserData) => {
      return updateUser(data)
    },
    onError: error => {
      const err = error as AxiosError<{ error: string }>
      if (err.response?.status === HttpStatusCode.Conflict) {
        if (err.response?.data.error === 'email already exists') {
          toast.error('此電子信箱已被註冊')
        } else if (err.response?.data.error === 'phone already exists') {
          toast.error('此電話號碼已被註冊')
        } else if (err.response?.data.error === 'line id already exists') {
          toast.error('此LINE ID已被註冊')
        }
      } else if (err.response?.status === HttpStatusCode.Unauthorized) {
        toast.error('請先登入')
        redirect('/login')
      } else {
        toast.error('更新失敗，請再試一次')
      }
    },
    onSuccess: () => {
      toast.success('更新成功')
      queryClient.invalidateQueries({ queryKey: ['user'] })
    }
  })
}

const UserInfo = () => {
  // fetch user data
  const { data: user } = useUserQuery()
  if (user === undefined) {
    redirect('/login')
  }

  // form control
  const [firstName, setFirstName] = useState<string>('')
  const [lastName, setLastName] = useState<string>('')
  const [email, setEmail] = useState<string>('')
  const [phone, setPhone] = useState<any>('')
  const [gender, setGender] = useState<Gender | string>('')
  const [birthYear, setBirthYear] = useState<number>(0)
  const [lineID, setLineID] = useState<string>('')
  const [spinner, setSpinner] = useState<boolean>(false)
  const [backupAddress, setBackupAddress] = useState<string>('') // in case user does not select address

  // google map
  let defaultPosition = { lat: 25.033964, lng: 121.564472 }
  const [markerPosition, setMarkerPosition] = useState(defaultPosition)
  const [selectedPlace, setSelectedPlace] =
    useState<google.maps.places.PlaceResult | null>(null)

  // initialize form data if user data is fetched
  useEffect(() => {
    if (user) {
      setFirstName(user?.data.first_name || '')
      setLastName(user?.data.last_name || '')
      setEmail(user?.data.email || '')
      setPhone(user?.data.phone || '')
      setGender(user?.data.gender || '')
      setBirthYear(user?.data.birth_year || 0)
      setLineID(user?.data.line_id || '')
      setBackupAddress(user?.data.address || '')
      setMarkerPosition({
        lat: user?.data.latitude || defaultPosition.lat,
        lng: user?.data.longitude || defaultPosition.lng
      })
    }
  }, [user])

  // if selectedPlace is changed, update marker position
  useMemo(() => {
    const newPosition = {
      lat: selectedPlace?.geometry?.location?.lat() || defaultPosition.lat,
      lng: selectedPlace?.geometry?.location?.lng() || defaultPosition.lng
    }
    setMarkerPosition(newPosition)
  }, [selectedPlace])

  // submit form
  // mutate user data
  const queryClient = useQueryClient()
  const mutation = useUpdateUserMutation(queryClient)
  const handleSubmit = async () => {
    // no user data, return
    if (!user) return

    // start spinner
    setSpinner(true)

    // update user data
    const address = selectedPlace?.formatted_address || backupAddress
    const data: UserData = {
      first_name: firstName,
      last_name: lastName,
      gender: gender as Gender,
      birth_year: birthYear as number,
      line_id: lineID,
      email: email,
      phone: phone as string,
      language: user?.data.language,
      address: address,
      latitude: selectedPlace?.geometry?.location?.lat() || markerPosition.lat,
      longitude: selectedPlace?.geometry?.location?.lng() || markerPosition.lng
    }
    mutation.mutate(data)
    setSpinner(false)
  }

  return (
    <section className="lg:mt-28 mt-32 bg-zinc-600 rounded-xl w-[90%] sm:w-[80%] md:w-[75%] lg:w-[60%] max-w-[1000px]  mx-auto md:px-0 px-6 shadow-xl shadow-gray-500 mb-32 pb-9">
      <div className="pt-10 mx-auto w-full md:w-4/5 max-w-[1000px]">
        <h1 className="my-8 text-3xl text-center text-white lg:mb-12 lg:text-3xl tracking-[20px]">
          會員資料
        </h1>
        {spinner && (
          <div className="flex justify-center items-center mx-auto mt-20 w-32 h-32">
            <Loading />
          </div>
        )}
        <div className={`${spinner ? 'hidden' : ''} mt-6 w-full rounded-xl`}>
          <div className="flex flex-wrap -mx-3 mb-3">
            <div className="px-3 mb-6 space-y-2 w-full md:mb-0 md:w-1/2">
              <label className="font-medium text-slate-200">姓氏</label>
              <input
                className="block py-3 px-4 w-full leading-tight rounded border appearance-none"
                id="grid-last-name"
                type="text"
                placeholder="姓氏"
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
          <div className="flex flex-wrap -mx-3 mb-6">
            <div className="px-3 mb-6 space-y-2 w-full md:mb-0 md:w-1/2">
              <label className="font-medium text-slate-200">性別</label>
              <select
                className="block py-3 px-4 w-full leading-tight rounded border appearance-none"
                id="gender"
                onChange={ev => setGender(ev.target.value)}
                value={gender}
                required
              >
                <option value={Gender.FEMALE}>女性</option>
                <option value={Gender.MALE}>男性</option>
                <option value={Gender.NOT_SPECIFIED}>其他</option>
                <option value={Gender.NOT_DISCLOSED}>不方便透露</option>
              </select>
            </div>
            <div className="px-3 space-y-2 w-full md:mb-0 md:w-1/2">
              <label className="font-medium text-slate-200">出生年份</label>
              <select
                className="block py-3 px-4 w-full leading-tight rounded border appearance-none"
                id="birth-year"
                onChange={ev => setBirthYear(Number(ev.target.value))}
                value={birthYear}
                required
              >
                <option value={0}></option>
                {YearOptions()}
              </select>
            </div>
          </div>
          <div className="flex flex-wrap px-3 -mx-3 mb-3 space-y-2">
            <label className="font-medium text-slate-200">LINE ID</label>
            <input
              className="block py-3 px-4 mb-3 w-full leading-tight rounded border appearance-none"
              id="line-id"
              type="text"
              placeholder="LINE ID"
              onChange={ev => setLineID(ev.target.value)}
              value={lineID}
            />
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
              placeholder="請輸入電話號碼"
              defaultCountry="TW"
              value={phone}
              onChange={ev => setPhone(ev)}
            />
          </div>
          <div className="flex flex-wrap px-3 -mx-3 mb-8 space-y-2">
            <label className="font-medium text-slate-200">地址（寄貨用）</label>
            <PlaceAutocomplete
              onPlaceSelect={setSelectedPlace}
              setBackupAddress={setBackupAddress}
              backupAddress={backupAddress}
            />
          </div>
          <div className="flex flex-wrap px-3 -mx-3 mb-6 rounded-lg">
            <Map
              defaultCenter={markerPosition}
              defaultZoom={15}
              className="overflow-hidden w-full rounded-lg h-[400px]"
            >
              <Marker position={markerPosition} />
            </Map>
            <MapHandler place={selectedPlace} />
          </div>
          <div className="flex flex-wrap -mx-3">
            <div className="flex justify-center items-end w-full h-auto text-white">
              <button
                className="inline-flex overflow-hidden relative justify-start items-center py-3 px-6 w-40 font-medium rounded-xl transition-all bg-zinc-500 group"
                onClick={handleSubmit}
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
        </div>
      </div>
    </section>
  )
}

export default UserInfo
