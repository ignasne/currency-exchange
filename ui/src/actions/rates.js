import axios from 'axios'

const API_URL = window.env.api.url.replace(/\/$/, '')

export const SET_RATES = 'SET_RATES'

export const getRates = (fromCurrency, toCurrency, amount) => dispatch => {
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
        type: SET_RATES,
        data: resp.data.data
      })
    )
    .catch(err => console.log(err))
}