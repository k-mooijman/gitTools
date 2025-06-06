import {id} from "../../params";


function Body() {
  console.log("-- body --");
    window.addEventListener("storage", (event) => {
        console.log("----");
    });

    onstorage = (event) => {
      console.log("---");
      console.log(window.localStorage.getItem("myCat"));
  };


  return (
      <div className="App-body">
          <p>
              ------------{id}-------------<br/>
              Edit <code>src/App.js</code> and save to reload.<br/>
              -------------------------<br/>
          </p>
      </div>
  );
}

export default Body;
