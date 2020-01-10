import { applyMiddleware, compose, createStore } from "redux";
import thunk from "redux-thunk";

import { reducers } from "./reducers";

export default function getStore(initialState = {}) {
  const doCompose =
    process.env.NODE_ENV === "development" &&
    typeof window === "object" &&
    window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__
      ? window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__
      : compose;

  return createStore(
    reducers,
    initialState,
    doCompose(applyMiddleware(thunk))
  );
}