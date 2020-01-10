import { combineReducers } from "redux";

import resultReducer from "./result";

export const reducers = combineReducers({
  result: resultReducer
});