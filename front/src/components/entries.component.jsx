import React, {useState, useEffect} from 'react';
import axios from 'axios';
import {Button, Form, Container, Modal, Col} from 'react-bootstrap';
import Swal from 'sweetalert2';
import withReactContent from 'sweetalert2-react-content'
import '../index.css'

import EntryCard from './entry-card.component';

const MySwal = withReactContent(Swal)

const Entries = () => {
    const [newEntry, setNewEntry] = useState({"dish":"", "macroes":{"protein":0, "fat": 0, "carbs":0}, "calories":0})
    const [newCalculateMassIndex, setNewCalculateMassIndex] = useState(false)
    const [newCalculateCalories, setNewCalculateCalories] = useState(false)
    const [newCalculate, setNewCalculate] = useState({"age":0, "gender":"", "height":0, "weight":0})

    const [entries, setEntries] = useState([])
    const [refreshData, setRefreshData] = useState(false)
    const [editEntry, setEditEntry] = useState({"change":false, "id":0})
    const [addNewEntry, setAddNewEntry] = useState(false)    
    const [showMisc, setMisc] = useState(false)

    const [showDeleteAll, setShowDeleteAll] = useState(false)
    const toggleButton = () => {
        setShowDeleteAll(!showDeleteAll);
      };


    const totalCalories = !entries ? null : entries.map((entry) => entry.calories).reduce((acc, value) => acc + + value, 0)

    useEffect(() => {
        getAllEntries();
    }, [])

    if(refreshData){
        setRefreshData(false);
        getAllEntries();
    } 

    return (
        <div>
            <Container>
                <Button onClick={() => setAddNewEntry(true)}>Track today's calories</Button>
                <Button onClick={() => setMisc(true)}>Misc</Button>
            </Container>
            <Container class="calorieCounter">                
                {showDeleteAll && <div>Your today's calories: {totalCalories}</div>}
            </Container>
            <Container>
                {entries != null && entries.map((entry, i) =>(
                    <EntryCard data={entry} deleteEntry={deleteEntry} setEditEntry={setEditEntry}/>
                ))}
            </Container>
            <Container>
                {showDeleteAll && <Button onClick={() => deleteAll()}>DELETE ALL</Button>}
            </Container>
            <Modal show={showMisc} onHide={() => setMisc(false)} centred>
                <Modal.Header closeButton>                    
                    <Modal.Title>What do you want to do?</Modal.Title>
                </Modal.Header>                
                <Modal.Body>                    
                    <Button onClick={() => setNewCalculateCalories(true)}>Calculate my Calorie Requirement</Button>
                    <Button onClick={() => setNewCalculateMassIndex(true)}>Calculate my Mass Index</Button>
                </Modal.Body>
            </Modal>
            <Modal show={newCalculateCalories} onHide={() => setNewCalculateCalories(false)} centred>
                <Modal.Header closeButton>                    
                    <Modal.Title>Enter the following data to calculate your Calorie Requirement</Modal.Title>
                </Modal.Header>                
                <Modal.Body>
                    <Form.Group>
                        <Form.Label>Age</Form.Label>
                        <Form.Control type="number" onChange={(event) => {newCalculate.age = event.target.value}}></Form.Control>
                        <Form>
                        <Form.Label>Gender</Form.Label>
                            {['radio'].map((type) => (
                                <div key={`inline-${type}`} className="mb-3">
                                    <Form.Check inline label="Male" value="Male" name="group1" type={type} onChange={(event) => {newCalculate.gender = event.target.value}}/>
                                    <Form.Check inline label="Female" value="Female" name="group1" type={type} onChange={(event) => {newCalculate.gender = event.target.value}}/>
                                </div>
                            ))}
                        </Form>
                        <Form.Label>Height</Form.Label>
                        <Form.Control type="number" onChange={(event) => {newCalculate.height = event.target.value}}></Form.Control>
                        <Form.Label>Weight</Form.Label>
                        <Form.Control type="number" onChange={(event) => {newCalculate.weight = event.target.value}}></Form.Control>
                    </Form.Group>
                    <Button onClick={() => calculateCalories()}> Calculate!</Button>
                    <Button onClick={() => setNewCalculateCalories(false)}>Cancel</Button>
                </Modal.Body>
            </Modal>
            <Modal show={newCalculateMassIndex} onHide={() => setNewCalculateMassIndex(false)} centred>
                <Modal.Header closeButton>                    
                    <Modal.Title>Enter the following data to calculate your Mass Index</Modal.Title>
                </Modal.Header>                
                <Modal.Body>
                    <Form.Group>
                        <Form.Label>Height</Form.Label>
                        <Form.Control type="number" onChange={(event) => {newCalculate.height = event.target.value}}></Form.Control>
                        <Form.Label>Weight</Form.Label>
                        <Form.Control type="number" onChange={(event) => {newCalculate.weight = event.target.value}}></Form.Control>
                    </Form.Group>
                    <Button onClick={() => calculateMassIndex()}>Calculate!</Button>
                    <Button onClick={() => setNewCalculateMassIndex(false)}>Cancel</Button>
                </Modal.Body>
            </Modal>
            <Modal show={addNewEntry} onHide={() => setAddNewEntry(false)} centred>
                <Modal.Header closeButton>                    
                    <Modal.Title>Add Meal</Modal.Title>
                </Modal.Header>                
                <Modal.Body>
                    <Form.Group>
                        <Form.Label>Dish</Form.Label>
                        <Form.Control onChange={(event) => {newEntry.dish = event.target.value}}></Form.Control>
                        <Form.Label>Protein</Form.Label>
                        <Form.Control type="number" onChange={(event) => {newEntry.macroes.protein = event.target.value}}></Form.Control>
                        <Form.Label>Fat</Form.Label>
                        <Form.Control type="number" onChange={(event) => {newEntry.macroes.fat = event.target.value}}></Form.Control>
                        <Form.Label>Carbs</Form.Label>
                        <Form.Control type="number" onChange={(event) => {newEntry.macroes.carbs = event.target.value}}></Form.Control>
                        <Form.Label>Calories</Form.Label>
                        <Form.Control type="number" onChange={(event) => {newEntry.calories = event.target.value}}></Form.Control>
                    </Form.Group>
                    <Button onClick={() => addEntry()}>Add</Button>
                    <Button onClick={() => setAddNewEntry(false)}>Cancel</Button>
                </Modal.Body>
            </Modal>
            <Modal show={editEntry.change} onHide={() => setEditEntry({"change":false, "id":0})} centred>
                <Modal.Header closeButton>                    
                    <Modal.Title>Change</Modal.Title>
                </Modal.Header>                
                <Modal.Body>
                    <Form.Group>
                        <Form.Label>Dish</Form.Label>
                        <Form.Control onChange={(event) => {newEntry.dish = event.target.value}}></Form.Control>
                        <Form.Label>Protein</Form.Label>
                        <Form.Control type="number" onChange={(event) => {newEntry.macroes.protein = event.target.value}}></Form.Control>
                        <Form.Label>Fat</Form.Label>
                        <Form.Control type="number" onChange={(event) => {newEntry.macroes.fat = event.target.value}}></Form.Control>
                        <Form.Label>Carbs</Form.Label>
                        <Form.Control type="number" onChange={(event) => {newEntry.macroes.carbs = event.target.value}}></Form.Control>
                        <Form.Label>Calories</Form.Label>
                        <Form.Control type="number" onChange={(event) => {newEntry.calories = event.target.value}}></Form.Control>
                    </Form.Group>
                    <Button onClick={() => editSingleEntry()}>Edit</Button>
                    <Button onClick={() => setEditEntry({"change":false, "id":0})}>Cancel</Button>
                </Modal.Body>
            </Modal>
        </div>
    );
    
    function addEntry() {
        setAddNewEntry(false)

        var url = "http://localhost:8080/entry/create/"
        axios.post(url, {
            "dish":newEntry.dish,
            "macroes": {
                "protein": parseInt(newEntry.macroes.protein),
                "fat": parseInt(newEntry.macroes.fat),
                "carbs": parseInt(newEntry.macroes.carbs),
            },
            "calories":parseInt(newEntry.calories),
        }).then(response =>{
            if(response.status == 200){
                setRefreshData(true)
                setShowDeleteAll(true)
            }
        })
    }

    function editSingleEntry() {
        editEntry.change = false;
        var url = "http://localhost:8080/entry/update/" + editEntry.id
        axios.put(url, {
            "dish":newEntry.dish,
            "macroes": {
                "protein": parseInt(newEntry.macroes.protein),
                "fat": parseInt(newEntry.macroes.fat),
                "carbs": parseInt(newEntry.macroes.carbs),
            },
            "calories":parseInt(newEntry.calories),
        })
        .then(response=>{
            if(response.status==200){
                setRefreshData(true)
            }
        })
    }

    function deleteEntry(id){
        var url = "http://localhost:8080/entry/delete/" + id
        axios.delete(url, {

        }).then(response =>{
            if(response.status == 200) {
                setRefreshData(true)
            }
        })
    }

    function getAllEntries() {
        var url = "http://localhost:8080/entries"
        axios.get(url, {
            responseType:"json",
        }).then(response => {
            if(response.status == 200) {
                setEntries(response.data)
            }
        })
    }

    function calculateCalories() {
        var url = "http://localhost:8080/calculate/calories/"
        axios.post(url, {
            "age":parseInt(newCalculate.age),
            "gender":newCalculate.gender,
            "height":parseInt(newCalculate.height),
            "weight":parseInt(newCalculate.weight),
        }, 
        {'Content-Type': 'application/json'})
        .then(response => {
            if(response.status == 200) {
                console.log(response.data)
                MySwal.fire("Your Calorie Requirement is: " + response.data)
                setRefreshData(true)
            }
        })
    }

    function calculateMassIndex() {
        var url = "http://localhost:8080/calculate/massindex/"
        axios.post(url, {
            "height":parseInt(newCalculate.height),
            "weight":parseInt(newCalculate.weight),
        }).then(response => {
            if(response.status == 200) {   
                console.log(response.data)             
                MySwal.fire("Your Body Mass Index is: " + response.data)
                setRefreshData(true)
            }
        }) 
    }

    function deleteAll() {
        var url = "http://localhost:8080/entry/delete/all"
        axios.delete(url, {
    
        }).then(response => {
            if(response.status == 200){
                MySwal.fire('All entries are deleted!',
                            '',
                            'success')
                setRefreshData(true)
                setShowDeleteAll(false)
                toggleButton()
            }
        })
    }
}



export default Entries; 