import React from 'react'
import CocktailCard from './CocktailCard'
import styled from 'styled-components'
const CocktailList = ({drinks}) => {
    const formatDrinks = drinks.map((item)=>{
        const {idDrink,strDrink,strDrinkThumb, strAlcoholic, strGlass} = item
        return {
          id: idDrink,
          name: strDrink,
          image: strDrinkThumb,
          info: strAlcoholic,
          glass:strGlass
        }
      })
    return <Wrapper>
    {formatDrinks.map((item)=>{
        return <CocktailCard key={item.id} {...item}/>
  })}
  </Wrapper>
}

export default CocktailList
const Wrapper = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 2rem;
`;