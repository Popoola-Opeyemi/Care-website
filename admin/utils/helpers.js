export function clone(val) {
  return JSON.parse(JSON.stringify(val))
}

export function Notification(self, message, type, duration = 2000) {
  self.$buefy.snackbar.open({
    duration: duration,
    message: message,
    type: type,
    position: 'is-top-right'
  })
}

export function Notify(self, message, type, duration = 2000) {
  self.$buefy.notification.open({
    message: message,
    type: type,
    duration: duration
  })
}

export function Dialog(self, message, type, onConfirm) {
  self.$buefy.dialog.alert({
    message: message,
    type: type,
    position: 'is-top',
    onConfirm: onConfirm
  })
}

export function Slugify(string) {
  const a =
    'àáâäæãåāăąçćčđďèéêëēėęěğǵḧîïíīįìłḿñńǹňôöòóœøōõṕŕřßśšşșťțûüùúūǘůűųẃẍÿýžźż·/_,:;'
  const b =
    'aaaaaaaaaacccddeeeeeeeegghiiiiiilmnnnnooooooooprrsssssttuuuuuuuuuwxyyzzz------'
  const p = new RegExp(a.split('').join('|'), 'g')

  return string
    .toString()
    .toLowerCase()
    .replace(/\s+/g, '-') // Replace spaces with -
    .replace(p, c => b.charAt(a.indexOf(c))) // Replace special characters
    .replace(/&/g, '-and-') // Replace & with 'and'
    .replace(/[^\w\-]+/g, '') // Remove all non-word characters
    .replace(/\-\-+/g, '-') // Replace multiple - with single -
    .replace(/^-+/, '') // Trim - from start of text
    .replace(/-+$/, '') // Trim - from end of text
}

export function unSlugify(str) {
  var newStr = str
    .split('-')
    .filter(function(item) {
      item = item ? item.replace(/-/g, '') : item
      return item
    })
    .join(' ')
  return newStr
}

export function dateFormater(date) {
  let year = date.getFullYear().toString(),
    month = (date.getMonth() + 1).toString(),
    day = date.getDate().toString()
  if (month.length === 1) {
    month = `0${month}`
  }
  if (day.length === 1) {
    day = `0${day}`
  }
  date = `${year}-${month}-${day}T00:00:00.000Z`
  return date
}

export function timeFormater(date) {
  let year = date.getFullYear().toString(),
    month = (date.getMonth() + 1).toString(),
    day = date.getDate().toString()
  let hrs = date.getHours().toString(),
    min = date.getMinutes().toString()
  if (year.length === 1) {
    year = `000${year}`
  }
  if (month.length === 1) {
    month = `0${month}`
  }
  if (day.length === 1) {
    day = `0${day}`
  }
  if (hrs.length === 1) {
    hrs = `0${hrs}`
  }
  if (min.length === 1) {
    min = `0${min}`
  }
  date = `${year}-${month}-${day}T${hrs}:${min}:00.000`
  let time = new Date(date)
  return time
}

export function timeParser(date) {
  let time = new Date(date).toTimeString().split(' ')[0]
  let hr = time.split(':')[0],
    min = time.split(':')[1]
  return `${hr}:${min}`
}

export function dateParser(date) {
  date = new Date(date).toDateString().split(' ')
  let [dayOfWeek, month, day, year] = date
  return `${day}, ${month} ${year}`
}

export function copy(val) {
  if (typeof val === 'object') {
    return JSON.parse(JSON.stringify(val))
  }
  return null
}
