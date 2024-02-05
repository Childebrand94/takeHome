type Props = {
    prefix: number
    operator: string
    country_code: number
    region: string
    country: string
}

export const TableRow: React.FC<Props> = (props) => {

    return (

        <tr>
            <th scope="row"></th>
            <td>{props.prefix || "n/a"}</td>
            <td>{props.operator || "n/a"}</td>
            <td>{props.country_code || "n/a"}</td>
            <td>{props.region || "n/a"}</td>
            <td>{props.country || "n/a"}</td>
        </tr>
    );
}
