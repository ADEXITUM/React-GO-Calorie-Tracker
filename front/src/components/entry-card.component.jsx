import React, {useState, useEffect} from 'react';
import 'bootstrap/dist/css/bootstrap.css';

import { Button, Col, Row, Card } from 'react-bootstrap';

const EntryCard = ({data, deleteEntry, setEditEntry}) => {
    return(
        <Card>
            <Row>
                <Col>Dish:{data !== undefined && data.dish}</Col>
                <Col>Protein:{data !== undefined && data.macroes.protein}</Col>
                <Col>Fat:{data !== undefined && data.macroes.fat}</Col>
                <Col>Carbs:{data !== undefined && data.macroes.carbs}</Col>
                <Col>Calories:{data !== undefined && data.calories}</Col>
                <Col><Button onClick={()=> deleteEntry(data._id)}>Delete</Button></Col>
                <Col><Button onClick={()=> editEntry()}>Edit</Button></Col>
            </Row>
        </Card>
    )

    function editEntry() {
        setEditEntry(
            {
                "change":true,
                "id":data._id
            }
        )
    }
}



export default EntryCard;
