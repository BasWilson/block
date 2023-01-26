import { useState } from "react";
import { RiArrowUpLine, RiCloseLine, RiUploadLine } from "react-icons/ri";

export const Status = () => {

    const [status, setStatus] = useState<"running" | "stopped" | "applying">("applying");

    if (status === "stopped") {
        return (
            <div className={'flex items-center px-2 py-2 text-sm text-red-500 font-medium rounded-md'}
            >
                <RiCloseLine
                    className={'text-red-500 mr-3 h-6 w-6'}
                    aria-hidden="true"
                />
                Stopped
            </div>
        )
    } else if (status === "applying") {
        return (
            <div className={'flex items-center px-2 py-2 text-sm text-orange-500 font-medium rounded-md'}
            >
                <RiUploadLine
                    className={'text-orange-500 mr-3 h-6 w-6'}
                    aria-hidden="true"
                />
                Applying changes
            </div>
        )
    } else {
        return (
            <div className={'flex items-center px-2 py-2 text-sm text-green-500 font-medium rounded-md'}
            >
                <RiArrowUpLine
                    className={'text-green-500 mr-3 h-6 w-6'}
                    aria-hidden="true"
                />
                Running
            </div>
        )
    }

}