import { useState, useEffect, useCallback, useMemo } from 'react'

const useGoogleMap = (defaultLat: number, defaultLng: number) => {
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

export default useGoogleMap
