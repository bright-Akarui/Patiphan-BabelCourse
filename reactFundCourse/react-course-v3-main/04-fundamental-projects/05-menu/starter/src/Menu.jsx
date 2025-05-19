import React from 'react'
import SingleMenu from "./SingleMenu";
const Menu = ({ menu }) => {
    console.log(menu)
    return (<div className='section-center'>
        {menu.map((menu) => {
            return <SingleMenu key={menu.id} {...menu} />
        })}
    </div>
    )
}

export default Menu