import React from 'react';
import { useState, useEffect } from 'react';

const Fetch = () => {
    const [characters, setCharacters] = useState([]);
    useEffect(() => {
        fetch('http://localhost:8080/characters')
            .then((res) => {
                return res.json()
            })
            .then((data) => {
                console.log(data.characters);
                setCharacters(data.characters);
            });
    }, []);
    console.log("Characters is:", characters);
    if (characters.length == 0) {
        return null;
    }

    return <table>
        <thead>
            <tr>
                <th>Name</th>
                <th>Class</th>
                <th>Level</th>
            </tr>
        </thead>
        <tbody>
            {characters.map(character =>
                <tr key={character.id}>
                    <td>{character.name}</td>
                    <td>{character.class}</td>
                    <td>{character.level}</td>
                </tr>)}
        </tbody>
    </table>
}

const CharacterList = () => {
    return (<div>
        <h1>Characters</h1>
        <Fetch />
    </div>)
}

export default CharacterList;