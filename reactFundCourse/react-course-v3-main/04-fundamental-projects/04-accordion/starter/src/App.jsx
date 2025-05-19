import { useState } from "react";
import questions from "./data";
import Questions from "./Questions";
const App = () => {

  return <div>
    <h2>Accordion Starter</h2>
    {questions.map((questions, index) => {
      const { id, title, info } = questions
      return <Questions key={id} title={title} info={info} index={index} />
    })}
  </div >;
};
export default App;
