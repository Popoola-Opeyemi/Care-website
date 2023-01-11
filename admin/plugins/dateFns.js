import Vue from "vue";
import dateFns from "date-fns/";
import parse from "date-fns/parse";
import format from "date-fns/format";
import formatIso from "date-fns/parseISO";

Vue.filter("formatDate", input => format(input, "MMMM D, YYYY"));
Vue.filter("formatIsoDate", input => format(formatIso(input), "MMMM d, yyyy"));
Vue.filter("formatIsoDateYear", input => format(formatIso(input), "yyyy"));
Vue.filter("output-date1", input => format(parse(input), "D MMM"));
Vue.filter("output-time", input => format(parse(input), "HH:mm a"));
