import { combineReducers } from 'redux'

import resultReducer from './result'
import errorReducer from './error'
import loadingIndicatorReducer from './loadingIndicator'

export const reducers = combineReducers({
  result: resultReducer,
  error: errorReducer,
  loadingIndicator: loadingIndicatorReducer
})