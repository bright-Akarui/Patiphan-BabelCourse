import React from 'react'

const SingleMenu = ({ id, category, desc, img, price, title }) => {
    return (
        <article className='menu-item'>
            <img src={img} alt={title} className='img' />
            <div className='item-info'>
                <header>
                    <h5>{title}</h5>
                    <span className='item-price'>${price}</span>
                </header>
                <p className='item-text'>{desc}</p>
            </div>
        </article>
    )
}

export default SingleMenu