import React from 'react'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'

import { getRates } from '../actions/rates'

const CURRENCIES = window.env.Currencies

class CurrencySection extends React.Component {
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
      <div>
        <select
          name="fromCurrency"
          value={this.state.fromCurrency}
          onChange={this.onChange}
        >
          <option value="">from</option>
          {currencyOptions}
        </select>
        <select
          name="toCurrency"
          value={this.state.toCurrency}
          onChange={this.onChange}
        >
          <option value="">to</option>
          {currencyOptions}
        </select>
        <input
          name="amount"
          type="number"
          autoComplete="off"
          value={this.state.amount}
          onChange={this.onChange}
        />
      </div>
    )
  }
}

CurrencySection.propTypes = {
  getRates: PropTypes.func.isRequired,
}

// @TODO errors handling
const mapStateToProps = function (state) {
  return {}
}

export default connect(mapStateToProps, {getRates: getRates})(
  CurrencySection
)