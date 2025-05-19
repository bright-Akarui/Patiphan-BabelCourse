import { FaAngleLeft, FaAngleRight } from "react-icons/fa";
import reviews from "../../final/src/data";
import { useState } from "react";

const App = () => {
  const [index, setIndex] = useState(0)
  const { id, name, job, image, text } = reviews[index]
  const nextPeople = () => {
    setIndex((nextPeople) => {
      const newIndex = (nextPeople + 1) % reviews.length
      return newIndex
    })
  }
  const prePeople = () => {
    setIndex((preIndex) => {
      const newIndex = (preIndex - 1) % reviews.length
      console.log(newIndex);
      return newIndex
    })
  }
  const randomPeople = () => {
    let randomNumber = Math.floor(Math.random() * reviews.length);
    if (randomNumber === index) {
      randomNumber = index + 1;
    }
    randomNumber = randomNumber % 4
    console.log(index)
    console.log(randomNumber)
    setIndex(randomNumber)
  }
  return <div>
    <h2> Reviews Starter</h2>
    <h4>{name}</h4>
    <img src={image} alt="images" style={{ width: 200 }} />
    <h5>{job}</h5>
    <h4>{text}</h4>
    <button onClick={prePeople}>
      <FaAngleLeft />
    </button>
    <button onClick={nextPeople}>
      <FaAngleRight />
    </button>
    <button onClick={randomPeople}>
      random
    </button>

  </div>

};
export default App;
