export interface AlertMessage {
     text: string,
     variant: string,
}

export const AlertMessageInitialState: AlertMessage = {
     text: "Start playing by creating a new game or joining an existing one!",
     variant: "dark",
};