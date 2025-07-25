import { Form, redirect, useNavigation } from 'react-router-dom';
import axios from 'axios';
import { toast } from 'react-toastify';
import { useState } from 'react';

const newsletterUrl = 'https://www.course-api.com/cocktails-newsletter';

export const action = async ({ request }) =>{
try {
  const formData = await request.formData();
  const data = Object.fromEntries(formData)
  const resp = await axios.post(newsletterUrl,data)
  
  console.log(resp);
  toast.success('เย้')
  return redirect('/');
} catch (error) {
  toast.error(error.response.data.msg)
  console.log(error.response.data.msg);
  return error.response.data.msg 
}

}
const Newsletter = () => { 
  const navigation = useNavigation();
  const isSubmitting = navigation.state === 'submitting';
  return (
    <Form className='form' method='POST'>
      {isSubmitting?'Loading...':''}
      <h4 style={{ textAlign: 'center', marginBottom: '2rem' }}>
        our newsletter
      </h4>
      {/* name */}
      <div className='form-row'>
        <label htmlFor='name' className='form-label'>
          name
        </label>
        <input
          type='text'
          className='form-input'
          name='name'
          id='name'
          defaultValue='test@test.com'
          required
        />
      </div>
      {/* lastName */}
      <div className='form-row'>
        <label htmlFor='lastName' className='form-label'>
          last name
        </label>
        <input
          type='text'
          className='form-input'
          name='lastName'
          id='lastName'
          defaultValue='test@test.com'
          required
        />
      </div>
      {/* email */}
      <div className='form-row'>
        <label htmlFor='email' className='form-label'>
          email
        </label>
        <input
          type='text'
          className='form-input'
          name='email'
          id='email'
          defaultValue='test@test.com'
          required
        />
      </div>
      <button
        type='submit'
        className='btn btn-block'
        style={{ marginTop: '0.5rem' }}
      >submit
      </button>
    </Form>
  );
};

export default Newsletter