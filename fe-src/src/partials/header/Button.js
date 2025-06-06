


function MyButton({v1,v2}) {

    function handleClick(name) {

        const kat = localStorage.getItem("myCat");
        localStorage.setItem("myCat", "joke"+kat);
        const cat = localStorage.getItem("myCat");

        alert('You clicked me!'+cat);

    }

    return (
        <button className="button-h" onClick={handleClick}>
            I'm a button {v1}
        </button>
    );
}

export default MyButton;
