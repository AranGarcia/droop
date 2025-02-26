import React, { useEffect, useState } from "react";
import { useParams } from "react-router";

import { characterType } from "../../types/character.types";

const CharacterDetail = () => {
    let {id} = useParams();
    const [character, setCharacter] = useState<characterType>({id: "",  class: "", level: 0, name: ""});
    useEffect(()=> {
        fetch(`http://localhost:8080/character/${id}`)
            .then((res) => {
                return res.json()
            })
            .then((data) => {
                setCharacter(data);
            });
    }, []);
    return <p>Hi, I am {character.name}, a {character.class} level {character.level}</p>
}

export default CharacterDetail;