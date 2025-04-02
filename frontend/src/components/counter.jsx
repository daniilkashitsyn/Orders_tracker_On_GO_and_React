import React, {useState} from "react";
import ReactDOM from "react-dom/client";

const Counter = () => {
    const [count, setCount] = useState(0)
    const [tags, setTags] = useState(['tag1', 'tag2', 'tag3'])

    const formatCount = () => {
        return count === 0 ? "empty" : count
    };

    const getBadgeClasses = () => {
        let classes = "badge m-2 ";
        classes += count === 0 ? "bg-warning" : "bg-primary";
        return classes;
    };

    const handleIncrement = () => {
        setCount((prevState => prevState + 1))
    }

    const handleDecrement = () => {
        setCount((prevState => prevState - 1))
    }

    const clearCount = () => {
        setCount((prevState => 0))
    }

    const handleTagChange = (id) => {
        setTags(prevState => prevState.filter((tag) => tag !== id))
    }

    const renderTags = () => {
        return tags.length === 0 ? "NoTags" : (
            tags.map(tag => (
                <li
                    key={tag}
                    className='btn btn-primary btn-sm m-2'
                    onClick={() => handleTagChange(tag)}
                >{tag}
                </li>
            ))
        )
    }

    return tags.length !== 0 ? <ul>
        {renderTags()}
    </ul> : (
        <React.Fragment>

            <span className={getBadgeClasses()}>{formatCount()}</span>
            <button
                className='btn btn-primary btn-sm m-2'
                onClick={handleIncrement}
            >+
            </button>
            <button
                className='btn btn-primary btn-sm m-2'
                onClick={handleDecrement}
            >-
            </button>
            <button
                className='btn btn-danger btn-sm m-2'
                onClick={clearCount}
            >Clear
            </button>
        </React.Fragment>
    );
}

export default Counter