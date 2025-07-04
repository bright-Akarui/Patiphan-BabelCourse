import React from 'react'
import { Outlet, useNavigation } from 'react-router-dom'
import Navbar from '../components/Navbar'

function HomeLayout() {
  const navigation = useNavigation();
  const isPageLoading = navigation.state === 'loading'
  return (
    <div>
      <Navbar />
      <section className='page'>
        {isPageLoading ? 'Loading...':
        <Outlet />}
        
      </section>
    </div>
    
  )
}

export default HomeLayout