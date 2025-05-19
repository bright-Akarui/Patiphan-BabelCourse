import React from 'react'
import { useState } from 'react';
const SingleTour = ({ id, image, info, name, price, deleteTours }) => {

    const [readMore, setReadMore] = useState(false)
    return (
        <div key={id}>
            <div className="single-tour">
                <img src={image} alt="tour1" className="img" />
                <span className='tour-price'>{price}</span>
                <div className="tour-info">
                    <h5>{name}</h5>
                    <p>{readMore ? info : `${info.substring(0, 200)}...`}
                        <button className='info-btn' onClick={() => setReadMore(!readMore)}>{readMore ? 'Show less' : 'Read more'}</button>
                    </p>
                    <button className='delete-btn' onClick={() => deleteTours(id)}>Not Interested</button>
                </div>
            </div>
        </div>
    )
}

export default SingleTour