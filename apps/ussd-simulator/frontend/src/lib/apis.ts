import {useEffect, useState} from "react";
import {Config} from "@/model/config";
import {Session} from "@/model/session";
import {ussdgateway} from "../../wailsjs/go/models.ts";
import {Exchange} from '../../wailsjs/go/ussdgateway/Gateway'
import {memoize} from "@monime-lab/twater2/memoize";


type ExchangeResponse = ussdgateway.ExchangeResponse;

const exchange = memoize((request: ussdgateway.ExchangeRequest & {
    idempotency?: string;
}) => {
    return Exchange(request);
}, (newArgs, lastArgs)=> {
    if(newArgs.length !== lastArgs.length) {
        return false;
    }
    return JSON.stringify(newArgs[0]) == JSON.stringify(lastArgs[0])
})

export const useExchange = (config: Config, session: Session) => {
    const [exchangeResponse, setExchangeResponse] = useState<ExchangeResponse>()
    const [error, setError] = useState<Error>()
    useEffect(() => {
        exchange({
            sessionId: session.id ?? '',
            idempotency: session.idempotency,
            networkId: config.network?.id ?? '',
            replyData: session.inputs?.reply ?? '',
            initialUssdCode: session.inputs?.ussdCode ?? '',
        }).then((response) => {
            setExchangeResponse(response)
        }).catch(error => {
            setError(error)
        })
    }, [config.network?.id, session.id, session.inputs?.reply, session.inputs?.ussdCode])
    return {response: exchangeResponse, error}
}