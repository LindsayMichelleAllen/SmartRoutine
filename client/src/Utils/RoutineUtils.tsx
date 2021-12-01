/**
 * This function gets the alarm time (date object) as a consistently human-readable object. Note
 * that this is the human-readable value which is different from the stored value.
 *
 * @param alarm The date value associated with the alarm.
 * @returns The human-readable string representing the {@link alarm}.
 */
export function GetAlarmText(alarm: Date) {
  // 1-index -> 0-index -> 1-index
  const alarmHours = ((alarm.getHours() - 1) % 12) + 1;
  const alarmMinutes = alarm.getMinutes();

  const alarmHourSyntax = alarmHours >= 10 ? `${alarmHours}` : `0${alarmHours}`;
  const alarmMinutesSyntax = alarmMinutes >= 10 ? `${alarmMinutes}` : `0${alarmMinutes}`;
  const meridian = alarm.getHours() <= 12 ? 'AM' : 'PM';

  return `${alarmHourSyntax}:${alarmMinutesSyntax} ${meridian}`;
}
