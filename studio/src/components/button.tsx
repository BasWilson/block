import React, { FC } from "react"

export const Button: FC<React.ButtonHTMLAttributes<HTMLButtonElement>> = (props) => {

    return (
        <button
            {...props}
            className="inline-flex items-center rounded-md border border-transparent bg-indigo-100 px-4 py-2 text-sm font-medium text-indigo-700 hover:bg-indigo-200 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
        >
            {props.children}
        </button>
    )
}