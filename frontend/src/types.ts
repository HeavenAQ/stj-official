export interface Order {
  name: string
  quantity: number
  note: string
}

export enum Gender {
  MALE = 'male',
  FEMALE = 'female',
  NOT_SPECIFIED = 'not-specified',
  NOT_DISCLOSED = 'not-disclosed'
}
