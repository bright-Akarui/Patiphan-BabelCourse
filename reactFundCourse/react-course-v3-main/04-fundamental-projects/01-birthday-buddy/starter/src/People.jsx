import React from 'react'
const People = ({people}) => {
    return (
        <div>
          <h2>{people.length} Birthdays Today</h2>
          {people.map((people)=>{
            const {id,name,age,image} = people;
            return <div key={id}>
              <div>{name}</div>
              <img style={{width:100}} src={image} alt="test" />
              <div>age:{age}</div>
            </div>
          })}
        </div>);
}

export default People