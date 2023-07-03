import SignUpPage from "./signUpPage";


const SignUpPageContainer = () => {
    function sleep(ms:number){
        return new Promise(resolve => setTimeout(resolve, ms));
    }
    function mockFunction(username: string){
        const mockAnswer = sleep(2000).then(() => false)
        return mockAnswer
    }

    return(
        <SignUpPage checkIfUsernameIsTaken={mockFunction}/>
    )
}

export default SignUpPageContainer