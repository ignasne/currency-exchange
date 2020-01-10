import { RATES_SET } from '../actions/rates'

const initialState = {
  exchangeRate: 0,
  amount: 0,
  currencyTo: ""
}

export default function (state = initialState, action) {
  if (action.type === RATES_SET) {
    return {
      ...state,
      exchangeRate: action.data['exchange_rate'],
      amount: action.data['amount'],
      currencyTo: action.data['currency_code']
    }
  }

  return state
}