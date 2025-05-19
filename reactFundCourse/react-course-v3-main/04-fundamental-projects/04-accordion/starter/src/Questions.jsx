import React from 'react'
import { useState } from 'react';
import { CiCirclePlus } from "react-icons/ci";

const Questions = ({ id, title, info, index }) => {
    const [toggle, setToggle] = useState(false);
    return (
        <div>
            <h4>{index + 1} {title} <button onClick={() => { setToggle(!toggle) }}><CiCirclePlus /></button></h4>
            {/* {toggle ? info : ''} */}
            {toggle && info}
            {/* false || true == true
            true || true == true1
            false && true == false
            true && true == true2
             */}
        </div>
    )
}

export default Questions