import React from 'react'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'

import { getRates } from '../actions/rates'
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

    if (amount <= 0) {
      console.log('@error amount less equal then zero')
      return
    }

    if (fromCurrency === toCurrency) {
      console.log('@error please select different currencies')
      return
    }

    if (!this.currencies.includes(fromCurrency) || !this.currencies.includes(toCurrency)) {
      console.log('@error Currency is not available')
      return
    }

    this.props.getRates(fromCurrency, toCurrency, amount)
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
}

const stateToProps = function () {
  return {}
}

export default connect(stateToProps, {getRates: getRates})(
  Currency
)