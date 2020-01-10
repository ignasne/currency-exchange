import React from 'react'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'
import styled from 'styled-components'

const StyledResultContainer = styled.div`
  width: 100%;
  text-align: center;
`
const StyledResult = styled.div`
  padding: 15px 0 0 0;
`
const StyledResultTitle = styled.div`
  font-size: 12px;
  padding-bottom: 12px;
  color: #777777;
`

function Result (props) {
  if (props.exchangeRate > 0) {
    // amount was retrieved in cents, lets update
    let amountFloat = props.amount / 100
    return (
      <StyledResultContainer>
        <StyledResult>
          <StyledResultTitle>Total amount</StyledResultTitle>
          {amountFloat} {props.currencyTo}
        </StyledResult>
        <StyledResult>
          <StyledResultTitle>Currency rate</StyledResultTitle>
          {props.exchangeRate}
        </StyledResult>
      </StyledResultContainer>
    )
  }

  return <span>{null}</span>
}

Result.propTypes = {
  exchangeRate: PropTypes.number.isRequired,
  amount: PropTypes.number.isRequired,
  currencyTo: PropTypes.string.isRequired
}

const stateToProps = function (state) {
  return {
    exchangeRate: state.result.exchangeRate,
    amount: state.result.amount,
    currencyTo: state.result.currencyTo
  }
}

export default connect(stateToProps)(Result)