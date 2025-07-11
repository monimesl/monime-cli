import React, {createContext, PropsWithChildren, useContext, useEffect, useRef, useState} from "react";
import {Config} from "@/model/config/index.tsx";
import {PartialDeep} from "@monime-lab/twater2/@types/data";
import {typedDeepMergeNew} from "@monime-lab/twater2/deepMerge";
import {deepEquals} from "@monime-lab/twater2/deepEqual";

const ConfigContext = createContext<{
    config: Config;
    setConfig: (checkout: PartialDeep<Config>) => void;
}>({
    config: {network: {id: 'm17'}} as Config,
    setConfig: () => {
        return
    },
});
export default function ConfigProvider(props: PropsWithChildren & {
    config: Config;
}) {
    const {config, children} = props
    const [internalConfig, setInternalConfig] = useState<Config>(config);
    const ref = useRef(internalConfig)
    useEffect(() => {
        setInternalConfig(config)
    }, [config])

    const setInternalData2 = React.useMemo(() => {
        return (src: PartialDeep<Config>) => {
            const result = typedDeepMergeNew(ref.current, {
                dontConcatDeepArray: true,
            }, src)
            if (!deepEquals(result, ref.current)) {
                setInternalConfig(result as Config);
                ref.current = result as Config
            }
        };
    }, []);
    return (
        <ConfigContext.Provider value={{config: internalConfig, setConfig: setInternalData2}}>
            {children}
        </ConfigContext.Provider>
    )
}
export const useConfig = () => {
    return useContext(ConfigContext);
};