import { RATES_GET_FAIL, RATES_FETCHING } from '../actions/rates'

const errorState = {
  message: ''
}

export default function (state = errorState, action) {
  switch (action.type) {
    case RATES_GET_FAIL:
    case RATES_FETCHING:
      return {
        message: action.data
      }
    default:
      break
  }

  return state
}