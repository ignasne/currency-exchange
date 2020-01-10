import axios from 'axios'

const API_URL = window.env.api.url.replace(/\/$/, '')

export const RATES_SET = 'RATES_SET'
export const RATES_GET_FAIL = 'RATES_GET_FAIL'
export const RATES_FETCHING = 'RATES_FETCHING'

export const getRates = (fromCurrency, toCurrency, amount) => dispatch => {
  // TODO manage fetching state and truncate lasting request to api if new one was initiated
  dispatch({
    type: RATES_FETCHING,
    data: ''
  })

  axios
    .get(`${API_URL}/api/quote`, {
      params: {
        from_currency_code: fromCurrency,
        to_currency_code: toCurrency,
        amount: amount
      }
    })
    .then(resp =>
      dispatch({
        type: RATES_SET,
        data: resp.data.data
      })
    )
    .catch(err => {
      // TODO process error messages from server accordingly show message to UI
      dispatch({
        type: RATES_GET_FAIL,
        data: 'Could not fetch rates from exchange.'
      })
    })
}