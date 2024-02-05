export type Resp = {
    prefixInfo: PhoneNumberInfo
    message: string
}

export type PhoneNumberInfo = {
    prefix: number;
    operator: string;
    country_code: number;
    region: string;
    country: string;
}
