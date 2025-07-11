import {useState} from "react";
import {Session} from "@/model/session";
import {ussdgateway} from "../../wailsjs/go/models.ts";
import {SendReply} from '../../wailsjs/go/ussdgateway/Gateway'


type SendReplyResponse = ussdgateway.ReplyResponse;

export const useSendReply = (session: Session) => {
    const [replyResponse, setReplyResponse] = useState<SendReplyResponse>()
    const [error, setError] = useState<Error>()
    SendReply({
        sessionId: session.id,
        replyMessage: session.inputs?.reply || ""
    }).then((response) => {
        setReplyResponse(response)
    }).catch(error => {
       setError(error)
    })
    return {reply: replyResponse, error}
}