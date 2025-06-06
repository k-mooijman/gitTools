import MyButton from './Button'


function Header() {
    return (
        <header className="App-header">

            <table>
                <tbody>
                <tr className="tc">
                    <td className="tr" colSpan={5}>
                        ---- Header ---
                    </td>
                </tr>
                <tr className="tc">
                    <td className="tr">
                        <MyButton v1={"hooi1"}/>
                    </td>
                    <td className="tr">
                        <MyButton/>
                    </td>
                    <td className="tr">
                        <MyButton/>
                    </td>
                    <td className="tr">
                        <MyButton/>
                    </td>
                    <td className="tr">
                        <MyButton/>
                    </td>
                </tr>
                </tbody>
            </table>

        </header>
    );
}


export default Header;
