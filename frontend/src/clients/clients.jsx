import React, { useState, useEffect } from "react"
import api from "../api"


const Clients = () => {
    const [clients, setClients] = useState([])

    useEffect(() => {
        console.log("Клиенты подключены")
        fetch('http://localhost:8080/clients?sort=a sc')
            .then((response) => response.json())
            .then((data) => {
               console.log("Клиенты: ", data)
                setClients(data)
            })
            .catch((error) => console.error("Ошбика запроса: ", error))
    }, []);

    const getColor = (client) => {
        switch (client.ration) {
            case "1200":
                return 'badge bg-warning'
            case "1500":
                return 'badge bg-primary'
            case "1800":
                return 'badge bg-success'
        }
    }

    return (
        <React.Fragment>
            <button className='btn btn-primary'>
                Фильтр
            </button>
            <table className='table table-borderless'>
                <thead>
                    <tr>
                        <th scope='col'>#</th>
                        <th scope='col'>Имя</th>
                        <th scope='col'>Рацион</th>
                    </tr>
                </thead>
                <tbody className='table-group-divider'>
                {clients.map((client) => (
                    <tr key={client.id}>
                        <th scope='col'>
                            {client.id}
                        </th>
                        <td>
                            {client.name}
                        </td>
                        <td className={getColor(client)}>
                            {client.ration}
                        </td>
                    </tr>)
                )}
                </tbody>
            </table>
        </React.Fragment>
    )
}


export default Clients
