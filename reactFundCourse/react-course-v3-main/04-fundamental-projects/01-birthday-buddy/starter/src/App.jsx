import { useState } from 'react';
import data from './data'
import People from './People';

const App = () => {
  const[people,setPeople] = useState(data);
  // console.log(people);
  return (
  <div>
    <People people={people}/>
    <button className='btn' onClick={()=>setPeople([])}>Clear person</button>
  </div>);
};
export default App;
