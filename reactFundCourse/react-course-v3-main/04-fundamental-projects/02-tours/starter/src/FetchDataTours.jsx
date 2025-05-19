import { useState, useEffect, useCallback } from "react";
import axios from "axios";

const FetchDataTours = (url) => {
  const [data, setData] = useState([]);

  const fetchData = useCallback(async () => {
    try {
      const response = await axios.get(url);
      setData(response.data);
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  }, [url]);

  useEffect(() => {
    fetchData();
  }, [fetchData]);

  return { data };
};

export default FetchDataTours;
