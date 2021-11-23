/**
 * @param alarm
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
