import { format, isToday, isThisYear, parseISO } from "date-fns";

export function formatDate(dateString: string | Date): string {
  const date =
    typeof dateString === "string" ? parseISO(dateString) : dateString;

  if (isToday(date)) {
    return format(date, "h:mm a"); // "2:02 AM"
  } else if (isThisYear(date)) {
    return format(date, "MMM d"); // "Oct 22"
  } else {
    return format(date, "MMM d, yyyy"); // "Oct 22, 2025"
  }
}
