import React from 'react'
import { Link } from 'react-router-dom'
import styled from 'styled-components'

const CocktailCard = ({id,name,info,glass,image}) => {
  return (<Wrapper>
    <div>{name}</div>
    <Link to={`/cocktail/${id}`}className='btn'>Detail</Link>
    </Wrapper>
  )
}
export default CocktailCard





const Wrapper = styled.article`
  background: var(--white);
  box-shadow: var(--shadow-2);
  transition: var(--transition);
  display: grid;
  grid-template-rows: auto 1fr;
  border-radius: var(--borderRadius);
  :hover {
    box-shadow: var(--shadow-4);
  }
  img {
    height: 15rem;
    border-top-left-radius: var(--borderRadius);
    border-top-right-radius: var(--borderRadius);
  }
  .footer {
    padding: 1.5rem;
    h4,
    h5 {
      margin-bottom: 0.5rem;
    }
    h4 {
      font-weight: 700;
    }
    p {
      margin-bottom: 1rem;
      color: var(--grey-500);
    }
  }
`;