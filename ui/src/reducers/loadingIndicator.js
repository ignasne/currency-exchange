import { RATES_GET_FAIL, RATES_SET, RATES_FETCHING } from '../actions/rates'

const indicatorState = {
  message: ''
}

export default function (state = indicatorState, action) {
  switch (action.type) {
    case RATES_GET_FAIL:
    case RATES_SET:
      return {
        message: ''
      }
    case RATES_FETCHING:
      return {
        message: 'Loading ...'
      }
    default:
      break
  }

  return state
}