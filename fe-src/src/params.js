const queryParameters = new URLSearchParams(window.location.search)
export const id = queryParameters.get("id")
export const type = queryParameters.get("type")
