export const getSession = key => {
  const json = window.sessionStorage[key]
  if (json == 'undefined' || json == undefined) {
    return false
  }
  return json ? JSON.parse(json) : false
}

export const saveSession = (key, value) => {
  window.sessionStorage[key] = JSON.stringify(value)
}

export const removeSession = key => {
  delete window.sessionStorage[key]
}
