import { useState } from "react";
import { Form } from "../components/Form";
import { TableRow } from "../components/TableRow";
import { Resp } from "../types/types";
import DOMPurify from 'dompurify';

const Home = () => {
    const [resp, setResp] = useState<Resp | null>(null);

    const createMarkup = (htmlString: string) => {
        const cleanHTML = DOMPurify.sanitize(htmlString)
        return { __html: cleanHTML };
    };
    return (
        <div className="bg-white p-5 rounded">
            <h3> Enter a Phone Number & Message </h3>
            <Form setResp={setResp} />
            <table className="table border rounded my-2">
                <thead>
                    <tr>
                        <th></th>
                        <th>Prefix</th>
                        <th>Operator</th>
                        <th>Country Code</th>
                        <th>Region</th>
                        <th>Country</th>
                    </tr>
                </thead>
                <tbody>
                    {resp ? (
                        <TableRow
                            prefix={resp.prefixInfo.prefix}
                            operator={resp.prefixInfo.operator}
                            country_code={resp.prefixInfo.country_code}
                            region={resp.prefixInfo.region}
                            country={resp.prefixInfo.country}
                        />
                    ) : (
                        <tr>
                            <td colSpan={5}>&nbsp;</td>
                        </tr>
                    )}
                </tbody>
            </table>
            {resp && (
                <div className="bg-white border rounded p-3 ">
                    <div>
                        <div dangerouslySetInnerHTML={createMarkup(resp.message)} />
                    </div>
                </div>
            )}
        </div>
    );
};

export default Home;

