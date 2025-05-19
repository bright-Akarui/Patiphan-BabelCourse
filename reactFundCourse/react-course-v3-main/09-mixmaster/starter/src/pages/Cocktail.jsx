import axios from 'axios'
import React from 'react'
import { useLoaderData } from 'react-router-dom';
const SingleCocktail = 'https://www.thecocktaildb.com/api/json/v1/1/lookup.php?i='

export const Loading = async({params}) =>{
  const { id } = params;
  const resp = await axios.get(`${SingleCocktail}${id}`);
  return {drinks: resp.data.drinks}
}


function Cocktail() {
  const {drinks} = useLoaderData();
  const format = drinks[0]
  const {
    idDrink : id,
    strDrink : name,
    strDrinkThumb : image,
    strAlcoholic : info,
    strGlass : glass,
    strInstructions: instruction
  } = format
  const valid = Object.keys(format).filter((key)=>key.startsWith('strIngredient') && 
  format[key] !== null).map((key)=> format[key]);
  console.log(valid.length)
  return (
    <div>
      Cocktail
      <h2>{name}</h2>
      <img src={image} alt="etest" />
      <h2>{info}</h2>
      <h2>{glass}</h2>
      {valid.map((item,index)=>{
        return <span key={index}>{item}{index < valid.length - 1?',':''}</span>})}
    </div>
  )
}

export default Cocktail