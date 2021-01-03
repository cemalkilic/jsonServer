import React, { useState } from 'react'
import axios from 'axios'
import './styled.scss'
import {ACCESS_TOKEN_NAME} from "../../constants/apiConstants";

const buildState = () => ({
  endpoint: "test/my/endpoint" + Math.floor(Math.random() * 169) + 13,
  content: `{
    "_id": "5fef60d16b19c93c8a736c8b",
    "isActive": true,
    "balance": "$1,340.01",
    "picture": "http://placehold.it/32x32"
}`,
  statusCode: 200,
});

const NewEndpointForm = (props) => {
  const [formData, setFormData] = useState(buildState())
  const [createdEndpoint, setCreatedEndpoint] = useState({})

  const refreshFormData = () => {
    setFormData({
      ...buildState()
    })
  }

  const updateCreatedEndpoint = endpoint => {
    setCreatedEndpoint({ endpoint: endpoint} )
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
    const payload = {
      "content": formData.content,
      "endpoint": formData.endpoint,
      "statusCode": formData.statusCode,
    }
    const token = localStorage.getItem(ACCESS_TOKEN_NAME);
    axios.post(
      '/addEndpoint',
        payload,
        token ? {headers: {'Authorization': 'Bearer ' + token}} : {}
    )
      .then(res => {
        updateCreatedEndpoint(res.data.endpoint)
        props.showError(null)
      })
      .catch(error => {
        updateCreatedEndpoint(null)
        const errorMessage = error.response.data.error || 'API Error!'
        props.showError(errorMessage)
      })

    refreshFormData()
  }

  return (
    <>
      <form className="newEndpointForm" onSubmit={handleSubmit}>
        <input
          type="text"
          name="endpoint"
          required={true}
          placeholder="Endpoint"
          onChange={updateInput}
          value={formData.endpoint || ''}
        />
        <input
            type="number"
            name="statusCode"
            required={true}
            placeholder="Status Code"
            onChange={updateInput}
            value={formData.statusCode || ''}
        />
        <textarea
            type="text"
            name="content"
            required={true}
            placeholder="JSON Content"
            onChange={updateInput}
            value={formData.content || ''}
        />

        {formData.error && <span className="error">{formData.error}</span>}

        <button type="submit">Create Endpoint</button>

        <div className={"alert-width"}>
        <div
            style={{display: createdEndpoint.endpoint ? 'block' : 'none'}}
            className={"alert alert-success"}
        >
          <a href={createdEndpoint.endpoint} target={"blank"} className="alert-link">Click</a> and see it in action!
          <textarea
              className={"createdEndpoint"}
              readOnly={true}
              value={createdEndpoint.endpoint || 'Created endpoint will be here!'}
          />
        </div>
        </div>
      </form>
    </>
  )
}

export default NewEndpointForm
