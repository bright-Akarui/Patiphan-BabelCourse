import React from 'react'
import { useLoaderData } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import CocktailList from '../components/CocktailList';
import SearchForm from '../components/SearchForm';
const CocktailUrl = 'https://www.thecocktaildb.com/api/json/v1/1/search.php?s='

const selectCocktailQuery = (searchTerm) => {
  return{
    queryKey: ['search', searchTerm || 'all'],
    queryFn: async()=>{
      const res = await axios.get(`${CocktailUrl}${searchTerm}`);
      return res.data.drinks
    }
  }
}

export const Loading = async ({request}) =>{
  const url = new URL(request.url)
  const searchTerm = url.searchParams.get('search') || 'a';
  // const res = await axios.get(`${CocktailUrl}${searchTerm}`);
  return { searchTerm }
}

function Landing() {
  const {searchTerm } = useLoaderData();
  const {data:drinks} = useQuery(selectCocktailQuery(searchTerm))
  console.log(drinks);
  return <>
  <SearchForm searchTerm={searchTerm}/>
  <CocktailList drinks={drinks}/>
  </>
}

export default Landing