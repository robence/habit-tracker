import React from "react";
import { HabitDetails } from "../habitDetails/HabitDetails";

export function Habit(props) {
  const pr = props.habit;
  return (
    <div className="collection-item col s12 m8 offset-m2 l6 offset-l3">
      <div className="row valign-wrapper">
        <div className="col s10">
          <span className="name">{pr.name}</span>
          <br />
          <HabitDetails details={pr.details} />
        </div>
      </div>
    </div>
  );
}
