import {useEffect} from "react";
import {newConfig} from "@/model/config";
import {newSession} from "@/model/session";
import ReplyScreen from "@/screens/ReplyScreen.tsx";
import FeedbackScreen from "@/screens/FeedbackScreen.tsx";
import DialPadScreen from "@/screens/DialPadScreen.tsx";
import LoadingScreen from "@/screens/LoadingScreen.tsx";
import ConfigProvider from "@/model/config/provider.tsx";
import TerminalScreen from "@/screens/TerminalScreen.tsx";
import ScreenContainer from "@/components/ScreenContainer.tsx";
import SessionProvider, {useSession} from "@/model/session/provider.tsx";

export default function App() {
    return (
        <ConfigProvider config={newConfig()}>
            <SessionProvider session={newSession()}>
                <Screen/>
            </SessionProvider>
        </ConfigProvider>
    )
}


const Screen = () => {
    const {session} = useSession();
    useEffect(() => {
        console.log("session", session);
    }, [session]);
    const renderScreen = () => {
        switch (session.screen) {
            case 'dial-pad':
                return <DialPadScreen/>
            case 'loading':
                return <LoadingScreen/>
            case 'reply':
                return <ReplyScreen/>
            case 'feedback':
                return <FeedbackScreen/>
            case 'terminal':
                return <TerminalScreen/>
            default:
                return session.screen
        }
    }
    return (
        <ScreenContainer>
            {renderScreen()}
        </ScreenContainer>
    )
}