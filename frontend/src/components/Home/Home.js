import React, {useEffect} from 'react';
import {withRouter} from 'react-router-dom';
import {ACCESS_TOKEN_NAME} from '../../constants/apiConstants';
import axios from 'axios'
import NewEndpointForm from "../NewEndpoint/NewEndpoint";

function Home(props) {
    useEffect(() => {
        axios.get('/user/me', {headers: {'Authorization': 'Bearer ' + localStorage.getItem(ACCESS_TOKEN_NAME)}})
            .then(function (response) {
                if (response.status !== 200) {
                    redirectToLogin()
                }
            })
            .catch(function (error) {
                redirectToLogin()
            });
    })

    function redirectToLogin() {
        props.history.push('/login');
    }

    return (
        <div className="mt-2">
            <NewEndpointForm/>
        </div>
    )
}

export default withRouter(Home);
