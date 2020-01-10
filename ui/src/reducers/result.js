import { SET_RATES } from '../actions/rates'

const initialState = {
  exchangeRate: 0,
  amount: 0
}

export default function (state = initialState, action) {
  if (action.type === SET_RATES) {
    return {
      ...state,
      exchangeRate: action.data['exchange_rate'],
      amount: action.data['amount'],
      currencyTo: action.data['currency_code']
    }
  }

  return state
}