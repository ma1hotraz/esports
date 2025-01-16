/**
 * Generate a random string of the specified length.
 *
 * @param {number} length - The length of the random string.
 * @returns {string} The generated random string.
 */
export function generateRandomString(length) {
  const characters =
    "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
  let result = "";
  for (let i = 0; i < length; i++) {
    result += characters.charAt(Math.floor(Math.random() * characters.length));
  }
  return result;
}

/**
 * Get the projection URL.
 *
 * @param {string[]} projectStringList - The list of project strings.
 * @returns {string} The projection URL.
 */
export function getProjectionUrl(projectStringList) {
  return "https://app.prizepicks.com/?projections=" + projectStringList.join(",");
}


/**
 * @typedef {Object} FieldValues
 * @property {number} PP - Value of the PP field
 * @property {number} UD - Value of the UD field
 * @property {number} SP - Value of the SP field
 * @property {boolean} hidePP - Whether to hide the PP field
 * @property {boolean} hideUD - Whether to hide the UD field
 * @property {boolean} hideSP - Whether to hide the SP field
 */

/**
 * @typedef {Object} Result
 * @property {boolean} prize_picks - Result for the PP field
 * @property {boolean} underdog - Result for the UD field
 * @property {boolean} sleeper - Result for the SP field
 */

/**
 * Compare the values of PP, UD, and SP and return which is the minimum, considering hidden or zero fields.
 * 
 * @param {FieldValues} fields - The field values and their hide flags
 * @returns {Result} The result indicating which field(s) is the minimum
 */
export function compareFields({ PP, UD, SP, hidePP, hideUD, hideSP }) {
  // Collect valid values, ignoring those with hide flags or zeros
  const validFields = [
    { value: PP, hidden: hidePP, field: 'PP' },
    { value: UD, hidden: hideUD, field: 'UD' },
    { value: SP, hidden: hideSP, field: 'SP' }
  ].filter(f => !f.hidden && f.value !== 0); // Filter out hidden or zero values

  // If there are no valid fields or only one valid field, return all false
  if (validFields.length < 2) {
    return { PP: false, UD: false, SP: false };
  }

  // Find the minimum value from the valid fields
  const minValue = Math.min(...validFields.map(f => f.value));

  // Create a result object where only the minimum value is set to true
  const result = {
    prize_picks: validFields.some(f => f.field === 'PP' && f.value === minValue),
    underdog: validFields.some(f => f.field === 'UD' && f.value === minValue),
    sleeper: validFields.some(f => f.field === 'SP' && f.value === minValue),
  };

  return result;
}

