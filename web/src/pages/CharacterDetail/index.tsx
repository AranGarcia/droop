import React, { useEffect, useState } from "react";
import { useParams } from "react-router";

import { characterType } from "../../types/character.types";
import { toTitleCase } from "../../shared/strings";

function getSignedNumber(value: number): string {
    var sign
    // negatives are already printed with sign 
    switch (true) {
        case value > 0:
            sign = "+";
            break;
        default:
            sign = "";
    }
    return `${sign}${value}`
}

function modifier(attribute: number): string {
    return getSignedNumber(Math.floor((attribute / 2) - 5))
}

const Attribute = ({ name, attribute }) => {
    return <div>
        <b>{name}: </b> {attribute} ({modifier(attribute)})
    </div>
}

const CharacterDetail = () => {
    let { id } = useParams();
    const [character, setCharacter] = useState<characterType>({ id: "", class: "", level: 0, name: "" });
    useEffect(() => {
        fetch(`http://localhost:8080/character/${id}`)
            .then((res) => {
                return res.json()
            })
            .then((data) => {
                setCharacter(data);
            });
    }, []);
    return (
        <div>
            <h1>{character.name}</h1>
            <h2> {toTitleCase(character.class)} level {character.level}</h2>
            <Attribute name="Strength" attribute={character.strength} />
            <Attribute name="Dexterity" attribute={character.dexterity} />
            <Attribute name="Constitution" attribute={character.constitution} />
            <Attribute name="Intelligence" attribute={character.intelligence} />
            <Attribute name="Wisdom" attribute={character.wisdom} />
            <Attribute name="Charisma" attribute={character.charisma} />
        </div>
    )
}

export default CharacterDetail;