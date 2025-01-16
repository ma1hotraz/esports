export function matchUsed(a, b) {
  const relevantFields = [
    "player",
    "stat_type",
    "timestamp",
    "team",
    "opponent",
    "sport",
  ];
  return relevantFields.every((field) =>
    compareStringsIgnoreCase(a[field], b[field])
  );
}

function compareStringsIgnoreCase(str1, str2) {
  return str1.toLowerCase() === str2.toLowerCase();
}

export function isTeamUsed(a, b) {
  const relevantFields = ["team", "opponent", "sport"];
  return relevantFields.every((field) =>
    compareStringsIgnoreCase(a[field], b[field])
  );
}

export function formatTime(isoString) {
  const date = new Date(isoString);
  const hours = date.getHours();
  const minutes = date.getMinutes();
  const ampm = hours >= 12 ? "PM" : "AM";
  const formattedHours = hours % 12 || 12;
  const formattedMinutes = minutes < 10 ? `0${minutes}` : minutes;
  return `${formattedHours}:${formattedMinutes} ${ampm}`;
}

export function formatDate(isoString) {
  const date = new Date(isoString);
  const month = date.getMonth() + 1;
  const day = date.getDate();
  const formattedMonth = month < 10 ? `0${month}` : month;
  const formattedDay = day < 10 ? `0${day}` : day;
  return `${formattedMonth}/${formattedDay}`;
}
