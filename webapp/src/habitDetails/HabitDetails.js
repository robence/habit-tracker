import React from "react";
import "./HabitDetails.css";

export function HabitDetails(props) {
  const cl = `skill card-panel lighten-5`;
  return (
    <div className={cl}>
      <span>{props.startDate}</span>
    </div>
  );
}
