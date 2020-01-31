import React from "react";
import { Habit } from "./Habit";
import "./Habit.css";

export function Habits(props) {
  const habits = props.habits.map((pr, i) => <Habit habit={pr} key={i} />);
  return <div>{habits}</div>;
}
