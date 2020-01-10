import React from 'react'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'
import styled from 'styled-components'

const StyledLoader = styled.div`
  width: 100%;
  text-align: center;
`

function LoadingIndicator (props) {
  if (props.message.length === 0) {
    return (null)
  }

  return (
    <StyledLoader>{props.message}</StyledLoader>
  )
}

LoadingIndicator.propTypes = {
  message: PropTypes.string.isRequired
}

const stateToProps = function (state) {
  return {
    message: state.loadingIndicator.message
  }
}

export default connect(stateToProps)(LoadingIndicator)