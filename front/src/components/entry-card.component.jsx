import React from 'react';
import 'bootstrap/dist/css/bootstrap.css';
import '../index.css';

import { Button, Col, Row, Card } from 'react-bootstrap';

const EntryCard = ({data, deleteEntry, setEditEntry}) => {
    return(
        <Card class="entryCard">
            <Row class="row">
                <Col>Dish:{data !== undefined && data.dish}</Col>
                <Col>Protein:{data !== undefined && data.macroes.protein}</Col>
                <Col>Fat:{data !== undefined && data.macroes.fat}</Col>
                <Col>Carbs:{data !== undefined && data.macroes.carbs}</Col>
                <Col>Calories:{data !== undefined && data.calories}</Col>
                <Col><Button variant="outline-danger" onClick={()=> deleteEntry(data._id)}>Delete</Button></Col>
                <Col><Button variant="outline-secondary" onClick={()=> editEntry()}>Edit</Button></Col>
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
