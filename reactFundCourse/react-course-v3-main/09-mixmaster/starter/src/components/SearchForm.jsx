import Wrapper from '../assets/wrappers/SearchForm';
import { Form, useNavigation } from 'react-router-dom';

const SearchForm = ({ searchTerm }) => {
  return (
    <Wrapper>
      <Form className='form'>
        <input
          type='search'
          name='search'
          className='form-input'
          defaultValue={searchTerm}
        />
        <button type='submit' className='btn'>
          search
        </button>
      </Form>
    </Wrapper>
  );
};
export default SearchForm;
