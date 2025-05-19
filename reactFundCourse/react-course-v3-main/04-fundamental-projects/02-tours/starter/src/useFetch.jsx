import { useState, useEffect, useCallback } from "react";
import axios from "axios";

const useFetch = (url) => {
    const [isLoading, setIsLoading] = useState(true)
    const [isError, setIsError] = useState(false)
    const [data, setData] = useState([])

    const fetchData = useCallback(async () => {
        try {
            setIsLoading(true)
            const response = await axios.get(url)
            setData(response.data)
            setIsLoading(false)
            console.log(data)
        } catch (error) {
            setIsLoading(true)
            setIsLoading(false)
            setIsError(false)
            setIsError(true)
            console.error("Error fetching data:", error)
        }
    }, [url]);

    useEffect(() => {
        fetchData()
    }, [fetchData])

    return { isError, isLoading, data, fetchData }
};

export default useFetch