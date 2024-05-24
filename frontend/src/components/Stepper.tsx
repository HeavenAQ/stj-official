import { FaCheck } from 'react-icons/fa'
import { MdContactMail } from 'react-icons/md'
import { BsBagCheckFill } from 'react-icons/bs'
import { FC } from 'react'
import { IconType } from 'react-icons'

interface StepperProps {
  activeIdx: 0 | 1 | 2
}

interface StepProps {
  index: number
  activeIdx: number
  Icon: IconType
  isLast: boolean
}

const getClassNames = (
  index: number,
  activeIdx: number,
  baseClasses: string,
  activeClasses: string,
  inactiveClasses: string
) => {
  return `${baseClasses} ${activeIdx >= index ? activeClasses : inactiveClasses}`
}

const Step: FC<StepProps> = ({ index, activeIdx, Icon, isLast }) => (
  <li
    className={`${getClassNames(
      index,
      activeIdx,
      'flex w-full items-center',
      'text-blue-600 dark:text-blue-500',
      'text-gray-500 dark:text-gray-400'
    )} ${
      !isLast
        ? getClassNames(
            index,
            activeIdx,
            "after:content-[''] after:w-full after:h-1 after:border-b after:border-4 after:inline-block",
            'after:border-blue-100 dark:after:border-blue-800',
            'after:border-gray-100 dark:after:border-gray-700'
          )
        : ''
    }`}
  >
    <span
      className={getClassNames(
        index,
        activeIdx,
        'flex items-center justify-center w-10 h-10 rounded-full lg:h-12 lg:w-12 shrink-0',
        'bg-blue-100 dark:bg-blue-800',
        'bg-gray-100 dark:bg-gray-700'
      )}
    >
      <Icon
        className={getClassNames(
          index,
          activeIdx,
          'w-3.5 h-3.5 lg:w-4 lg:h-4',
          'text-blue-600 dark:text-blue-300',
          'text-gray-500 dark:text-gray-100'
        )}
      />
    </span>
  </li>
)

const Stepper: FC<StepperProps> = ({ activeIdx }) => {
  return (
    <ol className="inline-flex items-center w-11/12 mb-10 mt-5 justify-center">
      <Step index={0} activeIdx={activeIdx} Icon={FaCheck} isLast={false} />
      <Step
        index={1}
        activeIdx={activeIdx}
        Icon={MdContactMail}
        isLast={false}
      />
      <li
        className={getClassNames(
          2,
          activeIdx,
          'flex items-center',
          'text-blue-600 dark:text-blue-500',
          'text-gray-500 dark:text-gray-400'
        )}
      >
        <span
          className={getClassNames(
            2,
            activeIdx,
            'flex items-center justify-center w-10 h-10 rounded-full lg:h-12 lg:w-12 shrink-0',
            'bg-blue-100 dark:bg-blue-800',
            'bg-gray-100 dark:bg-gray-700'
          )}
        >
          <BsBagCheckFill
            className={getClassNames(
              2,
              activeIdx,
              'w-4 h-4 lg:w-5 lg:h-5',
              'text-blue-600 dark:text-blue-300',
              'text-gray-500 dark:text-gray-100'
            )}
          />
        </span>
      </li>
    </ol>
  )
}

export default Stepper
