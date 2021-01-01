import React, { useState } from 'react'
import axios from 'axios'
import './styled.scss'

const buildState = () => ({
  endpoint: "test/my/endpoint",
  content: `{
    "_id": "5fef60d16b19c93c8a736c8b",
    "isActive": true,
    "balance": "$1,340.01",
    "picture": "http://placehold.it/32x32"
}`,
  statusCode: 200,
});

const NewEndpointForm = () => {
  const [formData, setFormData] = useState(buildState())

  const updateCreatedEndpoint = endpoint => {
    setFormData({

      resultEndpoint: endpoint
    })
  }

  const updateInput = e => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.type === 'number' ? parseInt(e.target.value) : e.target.value,
    })
  }

  const validateJSONInput = () => {
    try {
      JSON.parse(formData.content);
      return true;
    } catch (e) {
      return false;
    }
  }

  const handleSubmit = event => {
    event.preventDefault()

    const valid = validateJSONInput()
    if (!valid) {
      setFormData({
        ...formData,
        error: "JSON is invalid!"
      })
    } else {
      createEndpoint()
    }
  }

  const createEndpoint = () => {
    axios.post(
      '/addEndpoint',
      formData
    )
      .then(res => {
        updateCreatedEndpoint(res.data.endpoint)
      })
      .catch(error => {
        console.log(error.response)
        setFormData({
          ...formData,
          error: error.response.data.error
        })
      })
  }

  return (
    <>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          name="endpoint"
          placeholder="Endpoint"
          onChange={updateInput}
          value={formData.endpoint || ''}
        />
        <input
            type="number"
            name="statusCode"
            placeholder="Status Code"
            onChange={updateInput}
            value={formData.statusCode || ''}
        />
        <textarea
            type="text"
            name="content"
            placeholder="JSON Content"
            onChange={updateInput}
            value={formData.content || ''}
        />

        {formData.error && <span className="error">{formData.error}</span>}

        <button type="submit">Submit</button>

        <textarea
            readOnly={true}
            value={formData.resultEndpoint || 'Created endpoint will be here!'}
        />
      </form>
    </>
  )
}

export default NewEndpointForm
