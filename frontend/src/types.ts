export interface Order {
  name: string
  quantity: number
  note: string
}

export enum Gender {
  MALE = 'male',
  FEMALE = 'female',
  NOT_SPECIFIED = 'not_specified',
  NOT_DISCLOSED = 'not_disclosed'
}
