import React from 'react'
import { useState } from 'react';
import SingleTour from './SingleTour';

const Tours = ({ dataTours }) => {
    const [tours, setTours] = useState(dataTours)
    const deleteTours = (id) => {
        const newTours = tours.filter((tours) => tours.id !== id)
        setTours(newTours)
    }
    if (tours.length < 1) {
        return <button className='delete-btn' onClick={() => setTours(dataTours)}>Not Interested</button>
    }
    return <div className="tours">
        {tours.map((tours) => {
            const { id, image, info, name, price } = tours;
            return <SingleTour key={id} image={image} info={info} name={name} price={price} deleteTours={deleteTours} />
        })}
    </div>
}

export default Tours