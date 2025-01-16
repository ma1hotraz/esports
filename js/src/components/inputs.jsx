import SearchIcon from "@rsuite/icons/Search";
import { debounce } from "lodash";
import { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import { Input, InputGroup } from "rsuite";
const styles = {
    // marginLeft: '3rem'
};

export const CustomInput = ({ ...props }) => (
    <Input {...props} style={styles} />
);

export const CustomInputGroup = ({ placeholder, ...props }) => (
    <InputGroup {...props} style={styles}>
        <Input placeholder={placeholder} />
        <InputGroup.Addon>
            <SearchIcon />
        </InputGroup.Addon>
    </InputGroup>
);

export const CustomInputGroupWidthButton = ({
    placeholder,
    onSearchClick,
    ...props
}) => {
    const [inputValue, setInputValue] = useState("");
    const location = useLocation();
    useEffect(() => {
        setInputValue("");
        onSearchClick("");
    }, [location]);

    const handleInputChange = (event) => {
        setInputValue(() => event);
        delayedSearch(event);
    };

    const delayedSearch = debounce((term) => {
        onSearchClick(term);
    }, 500);

    return (
        <InputGroup {...props} inside style={styles}>
            <Input
                placeholder={placeholder}
                value={inputValue}
                onChange={handleInputChange}
            />
            <InputGroup.Button onClick={() => onSearchClick(inputValue)}>
                <SearchIcon />
            </InputGroup.Button>
        </InputGroup>
    );
};
