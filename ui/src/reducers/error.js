import { RATES_GET_FAIL, RATES_FETCHING, RATES_VALIDATION_ERROR, RATES_WIDGET_UPDATE } from '../actions/rates'

const errorState = {
  message: ''
}

export default function (state = errorState, action) {
  switch (action.type) {
    case RATES_GET_FAIL:
    case RATES_FETCHING:
    case RATES_VALIDATION_ERROR:
    case RATES_WIDGET_UPDATE:
      return {
        message: action.data
      }
    default:
      break
  }

  return state
}