import React, { useState } from 'react'
import Stepper from '../components/Stepper'
import DeliveryInfoStep from '../components/DeliveryInfoStep'
import { useUserDeliveryInfo } from '../utils/hooks'
import { toast } from 'react-hot-toast'
import CheckOrderStep from '../components/CheckOrderStep'
import OrderSummaryStep from '../components/OrderSummaryStep'

const activeContentStyle = (activeTab: number) => {
  switch (activeTab) {
    case 0:
      return 'translate-x-[calc(100%/3)]'
    case 1:
      return 'translate-x-0'
    case 2:
      return '-translate-x-[calc(100%/3)]'
    default:
      return ''
  }
}

interface StepButtonProps {
  text: string
  onClick: () => void
}
const StepButton: React.FC<StepButtonProps> = ({ text, onClick }) => {
  return (
    <button
      className="w-28 py-3 text-center tracking-[5px] bg-zinc-400 rounded-lg text-white duration-300 font-semibold hover:bg-zinc-300"
      onClick={onClick}
    >
      {text}
    </button>
  )
}

const submitOrder = async () => {
  // submit order to backend
}

const OrderPage: React.FC = () => {
  const [activeTab, setActiveTab] = useState(0)
  const [address, setAddress] = useState<string>('')
  const {
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
    setMarkerPosition,
    selectedPlace,
    setSelectedPlace
  } = useUserDeliveryInfo()

  const noEmptyFields = () => {
    return (
      firstName &&
      lastName &&
      email &&
      phone &&
      birthYear &&
      (selectedPlace || backupAddress)
    )
  }

  return (
    <div className="w-full h-full flex flex-col items-center justify-center mb-28 overflow-x-hidden animate-fade-down">
      <Stepper activeIdx={activeTab as 0 | 1 | 2} />

      <div
        className={`py-4 w-[300%] grid grid-cols-3 transition-transform duration-300 ease-in-out ${activeContentStyle(activeTab)}`}
      >
        <div className="w-full text-center">
          <CheckOrderStep />
          <StepButton
            text="下一步"
            onClick={() => setActiveTab(prevTab => Math.min(prevTab + 1, 2))}
          />
        </div>
        <div>
          <DeliveryInfoStep
            firstName={firstName}
            setFirstName={setFirstName}
            lastName={lastName}
            setLastName={setLastName}
            email={email}
            setEmail={setEmail}
            phone={phone}
            setPhone={setPhone}
            birthYear={birthYear}
            setBirthYear={setBirthYear}
            sameAsUser={sameAsUser}
            setSameAsUser={setSameAsUser}
            backupAddress={backupAddress}
            setBackupAddress={setBackupAddress}
            markerPosition={markerPosition}
            setMarkerPosition={setMarkerPosition}
            selectedPlace={selectedPlace}
            setSelectedPlace={setSelectedPlace}
          />
          <div className="inline-flex w-full justify-evenly">
            <StepButton
              text="上一步"
              onClick={() => {
                setActiveTab(prevTab => Math.max(prevTab - 1, 0))
                window.scrollTo(0, 0)
              }}
            />
            <StepButton
              text="下一步"
              onClick={() => {
                setAddress(selectedPlace?.formatted_address || backupAddress)
                noEmptyFields()
                  ? setActiveTab(prevTab => Math.min(prevTab + 1, 2))
                  : toast.error('請填寫所有欄位')
                window.scrollTo(0, 0)
              }}
            />
          </div>
        </div>
        <div>
          <OrderSummaryStep
            firstName={firstName}
            lastName={lastName}
            email={email}
            phone={phone}
            birthYear={birthYear}
            address={address}
          />
          <div className="inline-flex w-full justify-evenly">
            <StepButton
              text="上一步"
              onClick={() => {
                setActiveTab(prevTab => Math.max(prevTab - 1, 0))
                window.scrollTo(0, 0)
              }}
            />
            <StepButton
              text="送出訂單"
              onClick={() => {
                setAddress(selectedPlace?.formatted_address || backupAddress)
                noEmptyFields()
                  ? setActiveTab(prevTab => Math.min(prevTab + 1, 2))
                  : toast.error('請填寫所有欄位')
              }}
            />
          </div>
        </div>
      </div>
    </div>
  )
}
export default OrderPage
