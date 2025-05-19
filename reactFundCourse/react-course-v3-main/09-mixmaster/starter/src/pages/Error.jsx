import React from 'react'
import { useRouteError } from 'react-router-dom'

function Error() {
  const Error = useRouteError()
  return (
    <div>{Error.message}</div>
  )
}

export default Error