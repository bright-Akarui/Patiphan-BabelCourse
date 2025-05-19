
import Tours from "./Tours";
import useFetch from "./useFetch";

const url = 'https://www.course-api.com/react-tours-project';
const App = () => {
  const { isError, isLoading, data: dataTours, fetchData } = useFetch(url);
  // console.log(dataTours);
  return <div>
    <div className="title">
      <h2>Our Tour</h2>
      <div className="title-underline"></div>
    </div>

    {isLoading ? <div className="loading"></div>
      : isError ? <div><h2>Error...</h2></div>
        : <Tours dataTours={dataTours} />}
    <button onClick={() => fetchData()}>Refresh</button>
  </div>;
};
export default App;
