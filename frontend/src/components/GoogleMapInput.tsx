import { Map, Marker } from '@vis.gl/react-google-maps'
import { FC } from 'react'
import MapHandler from './MapHandler'
import { PlaceAutocomplete } from './PlaceAutocomplete'

interface GoogleMapInputProps {
  backupAddress: string
  setBackupAddress: (address: string) => void
  markerPosition: { lat: number; lng: number }
  selectedPlace: google.maps.places.PlaceResult | null
  setSelectedPlace: (place: google.maps.places.PlaceResult | null) => void
}

const GoogleMapInput: FC<GoogleMapInputProps> = ({
  backupAddress,
  setBackupAddress,
  markerPosition,
  selectedPlace,
  setSelectedPlace
}) => {
  return (
    <>
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
    </>
  )
}

export default GoogleMapInput
