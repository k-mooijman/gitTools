import { useDispatch } from 'react-redux'


function useTest(){

    console.log("test");
    const dispatch = useDispatch()
    dispatch({ type: 'repos/repoAdded', payload: "FromTest 001" })
    dispatch({ type: 'repos/repoAdded', payload: "FromTest 002" })


}

export default useTest;
