import React, {createContext, PropsWithChildren, useContext, useEffect, useRef, useState} from "react";
import {Session} from "@/model/session";
import {PartialDeep} from "@monime-lab/twater2/@types/data";
import {typedDeepMergeNew} from "@monime-lab/twater2/deepMerge";
import {deepEquals} from "@monime-lab/twater2/deepEqual";

const SessionContext = createContext<{
    session: Session;
    setSession: (checkout: PartialDeep<Session & { replace?: true }>) => void;
}>({
    session: {} as Session,
    setSession: () => {
        return
    },
});
export default function SessionProvider(props: PropsWithChildren & {
    session: Session;
}) {
    const {session, children} = props
    const [internalSession, setInternalSession] = useState<Session>(session);
    const ref = useRef(internalSession)
    useEffect(() => {
        setInternalSession(session)
    }, [session])

    const setInternalData2 = React.useMemo(() => {
        return (src0: PartialDeep<Session & { replace?: true }>) => {
            const {replace, ...src} = src0
            if (replace) {
                setInternalSession(src as Session);
                ref.current = src as Session
            } else {
                const result = typedDeepMergeNew(ref.current, {
                    dontConcatDeepArray: true,
                }, src)
                if (!deepEquals(result, ref.current)) {
                    setInternalSession(result as Session);
                    ref.current = result as Session
                }
            }
        };
    }, []);
    return (
        <SessionContext.Provider value={{session: internalSession, setSession: setInternalData2}}>
            {children}
        </SessionContext.Provider>
    )
}
export const useSession = () => {
    return useContext(SessionContext);
};