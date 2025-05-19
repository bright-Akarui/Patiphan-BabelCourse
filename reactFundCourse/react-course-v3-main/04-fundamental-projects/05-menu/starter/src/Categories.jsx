import React from 'react'

function Categories({ categories }) {
    return (
        <div className='btn-container'>
            {categories.map((categories) => {
                return <div key={categories}>
                    <button className='btn'>{categories}</button>
                </div>
            })}
        </div>
    )
}

export default Categories