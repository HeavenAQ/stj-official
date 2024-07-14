import { useState, useEffect, useCallback, useMemo } from 'react'
import { useUserQuery } from './query'

// use default lat and lng for google map marker (+ convert them to number)
const lat = +import.meta.env.VITE_GOOGLE_MAP_DEFAULT_LAT
const lng = +import.meta.env.VITE_GOOGLE_MAP_DEFAULT_LNG

// google map hook
export const useGoogleMap = (defaultLat: number, defaultLng: number) => {
  const defaultPosition = useMemo(
    () => ({ lat: defaultLat, lng: defaultLng }),
    [defaultLat, defaultLng]
  )
  const [backupAddress, setBackupAddress] = useState<string>('')
  const [markerPosition, setMarkerPosition] = useState(defaultPosition)
  const [selectedPlace, setSelectedPlace] =
    useState<google.maps.places.PlaceResult | null>(null)

  useEffect(() => {
    const newPosition = {
      lat: selectedPlace?.geometry?.location?.lat() || defaultPosition.lat,
      lng: selectedPlace?.geometry?.location?.lng() || defaultPosition.lng
    }
    setMarkerPosition(newPosition)
  }, [defaultPosition, selectedPlace?.geometry?.location])

  const handlePlaceSelect = useCallback(
    (place: google.maps.places.PlaceResult | null) => {
      setSelectedPlace(place)
    },
    []
  )

  const handleBackupAddressChange = useCallback((address: string) => {
    setBackupAddress(address)
  }, [])

  return {
    backupAddress,
    setBackupAddress: handleBackupAddressChange,
    markerPosition,
    setMarkerPosition,
    selectedPlace,
    setSelectedPlace: handlePlaceSelect
  }
}

export const useUserDeliveryInfo = () => {
  // fetch user data
  const { data: user } = useUserQuery()

  // form control
  const [firstName, setFirstName] = useState<string>('')
  const [lastName, setLastName] = useState<string>('')
  const [email, setEmail] = useState<string>('')
  const [phone, setPhone] = useState<any>('')
  const [birthYear, setBirthYear] = useState<number>(0)
  const [sameAsUser, setSameAsUser] = useState<boolean>(false)

  // wrap useState funciton inside useCallBack
  const handleFirstNameChange = useCallback((firstName: string) => {
    setFirstName(firstName)
  }, [])

  const handleLastNameChange = useCallback((lastName: string) => {
    setLastName(lastName)
  }, [])

  const handleEmailChange = useCallback((email: string) => {
    setEmail(email)
  }, [])

  const handlePhoneChange = useCallback((phone: any) => {
    setPhone(phone)
  }, [])

  const handleBirthYearChange = useCallback((birthYear: number) => {
    setBirthYear(birthYear)
  }, [])

  const handleSameAsUserChange = useCallback((sameAsUser: boolean) => {
    setSameAsUser(sameAsUser)
  }, [])

  // google map
  const {
    backupAddress,
    setBackupAddress,
    setMarkerPosition,
    markerPosition,
    selectedPlace,
    setSelectedPlace
  } = useGoogleMap(lat, lng)

  // fill in user data if sameAsUser is checked
  useEffect(() => {
    if (sameAsUser && user) {
      setFirstName(user.data.first_name)
      setLastName(user.data.last_name)
      setEmail(user.data.email)
      setPhone(user.data.phone)
      setBirthYear(user.data.birth_year || 0)
      setBackupAddress(user.data.address)
      setMarkerPosition({
        lat: user.data.latitude || lat,
        lng: user.data.longitude || lng
      })
    } else if (!sameAsUser) {
      setFirstName('')
      setLastName('')
      setEmail('')
      setPhone('')
      setBirthYear(0)
      setBackupAddress('')
      setMarkerPosition({ lat, lng })
    }
  }, [sameAsUser, user, setBackupAddress, setMarkerPosition])

  return {
    firstName,
    setFirstName: handleFirstNameChange,
    lastName,
    setLastName: handleLastNameChange,
    email,
    setEmail: handleEmailChange,
    phone,
    setPhone: handlePhoneChange,
    birthYear,
    setBirthYear: handleBirthYearChange,
    sameAsUser,
    setSameAsUser: handleSameAsUserChange,
    backupAddress,
    setBackupAddress: setBackupAddress,
    markerPosition,
    setMarkerPosition,
    selectedPlace,
    setSelectedPlace
  }
}
