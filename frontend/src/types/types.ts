export type Resp = {
    prfixInfo: PhoneNumberInfo
    message: string

}

export type PhoneNumberInfo = {
    operator: string;
    countryCode: number;
    region: string;
    country: string;
}
