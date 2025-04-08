import React, {useState, use} from "react";
import "./clientStyles.css"
import "../../src/styles.css"

const AddClient = () => {
    const [client, setClient] = useState({});

    return (
        <React.Fragment>
            <h2 className="w-25 mx-auto">Добавление клиента</h2>
            <div className="container col-md-1 mx-auto">
                <div className="border border-4">
                    Test
                </div>
            </div>
        </React.Fragment>
    )
}

export default AddClient;