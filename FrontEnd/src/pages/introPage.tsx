import { Bounce } from "gsap/all"
import { gsap } from "../gsap"
import { useEffect, useRef } from "react"



const Welcome = () => {
    useEffect(() => {
        const el = titleRef.current
        gsap.fromTo(el,{rotation:0},{rotation:360,duration:2,ease:Bounce.easeIn,scrollTrigger:{
            trigger:el,
        }})
    },[true])

    const titleRef = useRef(null)

    return(
        <div className="w-full min-h-screen bg-base flex flex-col justify-center items-center font-belanosima">
            <h1 className="card-title text-[40px] font-normal text-primary" ref={titleRef}>Hello, stranger</h1>
        </div>
    )
}

export default Welcome