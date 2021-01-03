import React, {useEffect} from 'react';
import {withRouter} from 'react-router-dom';
import {ACCESS_TOKEN_NAME} from '../../constants/apiConstants';
import axios from 'axios'
import NewEndpointForm from "../NewEndpoint/NewEndpoint";

function Home(props) {

    return (
        <div className="mt-2">
            <NewEndpointForm showError={props.showError}/>
        </div>
    )
}

export default withRouter(Home);
