import React from 'react'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'

function ResultSection (props) {
  if (props.exchangeRate > 0) {
    return (
      <div>
        <span>Currency rate:</span>
        {props.exchangeRate}
        <span> | </span>
        <span>Total amount:</span>
        {props.amount}
      </div>
    )
  }

  return <span>{null}</span>
}

ResultSection.propTypes = {
  exchangeRate: PropTypes.number.isRequired,
  amount: PropTypes.number.isRequired
}

const stateToProps = function (state) {
  return {
    exchangeRate: state.result.exchangeRate,
    amount: state.result.amount
  }
}

export default connect(stateToProps)(ResultSection)