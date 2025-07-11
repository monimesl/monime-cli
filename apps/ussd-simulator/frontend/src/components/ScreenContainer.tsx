import {PropsWithChildren} from "react";
import Frame from "@/components/Frame.tsx";
import {useSession} from "@/model/session/provider.tsx";

export default function ScreenContainer(props: PropsWithChildren) {
    const {children} = props;
    const {session} = useSession();
    return (
        <Frame darkBackground={session.screen != 'dial-pad'}>
            {children}
        </Frame>
    )
}