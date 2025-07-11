import {uuid4} from "@monime-lab/twater2/uuid";

export type Session = {
    id: string
    screen?: 'dial-pad' | 'loading' | 'feedback' | 'reply' | 'terminal'
    inputs?: {
        reply?: string
        ussdCode?: string
    },
    outputs?: {
        message?: string
    }
}

export const newSession = (screen?: Session['screen']): Session => {
    return {
        id: uuid4(),
        screen: screen || 'dial-pad',
    } as Session
}