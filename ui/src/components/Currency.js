import React from 'react'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'

import { validationError, ratesWidgetUpdate, getRates } from '../actions/rates'
import styled from 'styled-components'

const CURRENCIES = window.env.Currencies

const StyledCurrencyContainer = styled.div`
  width: 100%;
`

const StyledInputContainer = styled.div`
  border: 1px solid #ccc;
  width: 120px;
  border-radius: 3px;
  margin: 0 auto 12px auto;
  overflow: hidden;
  padding: 5px 8px;
  background: #fafafa no-repeat 90% 50%;
`

const StyledDropdown = styled.select`
  padding: 0;
  width: 100%;
  border: none;
  box-shadow: none;
  background: transparent;
  background-image: none;
  -webkit-appearance: none;
  &:focus {
    outline: none;
  }
`

const StyledInput = styled.input`
  padding: 0;
  width: 100%;
  border: none;
  box-shadow: none;
  background: transparent;
  background-image: none;
  &:focus {
    outline: none;
  }
`

class Currency extends React.Component {
  constructor (props) {
    super(props)

    this.state = {
      fromCurrency: '',
      toCurrency: '',
      amount: ''
    }

    this.currencies = CURRENCIES

    this.onChange = this.onChange.bind(this)
  }

  exchange = () => {
    const {fromCurrency, toCurrency, amount} = this.state

    this.props.ratesWidgetUpdate()

    // let's assume that our exchange converter accepts only two digits precision
    let regex = /^[1-9]\d*(((,\d{3}){1})?(\.\d{0,2})?)$/
    let amountFormatValid = regex.test(amount)

    if (
      amount > 0 &&
      amountFormatValid &&
      fromCurrency !== toCurrency &&
      this.currencies.includes(fromCurrency) &&
      this.currencies.includes(toCurrency)
    ) {
      // amount should be given in cents
      let amountInCents = amount * 100
      this.props.getRates(fromCurrency, toCurrency, amountInCents)
      return
    }

    if (amount !== '' && amount <= 0) {
      this.props.validationError('Amount should be greater then zero.')
      return
    }

    if (amount > 0 && !amountFormatValid) {
      this.props.validationError('Please use two decimal parts in currency amount.')
      return
    }

    if (fromCurrency === toCurrency) {
      this.props.validationError('Please select not equal currencies')
      return
    }

    if (!this.currencies.includes(fromCurrency) && !this.currencies.includes(toCurrency)) {
      this.props.validationError('Your provided currency is not supported.')
    }
  }

  onChange (e) {
    this.setState({[e.target.name]: e.target.value}, this.exchange)
  }

  render () {
    const currencyOptions = this.currencies.map(currency => (
      <option key={currency}>{currency}</option>
    ))

    return (
      <StyledCurrencyContainer>
        <StyledInputContainer>
          <StyledDropdown
            name="fromCurrency"
            value={this.state.fromCurrency}
            onChange={this.onChange}
          >
            <option value="">from currency</option>
            {currencyOptions}
          </StyledDropdown>
        </StyledInputContainer>
        <StyledInputContainer>
          <StyledDropdown
            name="toCurrency"
            value={this.state.toCurrency}
            onChange={this.onChange}
          >
            <option value="">to currency</option>
            {currencyOptions}
          </StyledDropdown>
        </StyledInputContainer>
        <StyledInputContainer>
          <StyledInput
            name="amount"
            type="number"
            autoComplete="off"
            value={this.state.amount}
            onChange={this.onChange}
          />
        </StyledInputContainer>
      </StyledCurrencyContainer>
    )
  }
}

Currency.propTypes = {
  getRates: PropTypes.func.isRequired,
  ratesWidgetUpdate: PropTypes.func.isRequired,
  validationError: PropTypes.func.isRequired,
}

const stateToProps = function () {
  return {}
}

export default connect(stateToProps, {ratesWidgetUpdate, validationError, getRates})(
  Currency
)