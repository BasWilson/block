import Editor, { Monaco } from "@monaco-editor/react";
import monaco from "monaco-editor";
import { useRef } from "react";
import { Button } from "../components/button";

export const ConfigPage = () => {

    const monacoRef = useRef<Monaco | null>(null);
    const editorRef = useRef<monaco.editor.IStandaloneCodeEditor | null>(null);

    function handleEditorWillMount(monaco: Monaco) {
    }

    function handleEditorDidMount(editor: monaco.editor.IStandaloneCodeEditor, monaco: Monaco) {
        monacoRef.current = monaco;
        editorRef.current = editor;
        setTimeout(() => {
            editor.getAction("editor.action.formatDocument").run();
        }, 50)
    }

    return (
        <>
            <div className="relative z-0 flex flex-1 overflow-hidden">
                <main className="relative z-0 flex-1 overflow-y-auto focus:outline-none">
                    {/* Start main area*/}
                    <div className="absolute inset-0 py-6">
                        <Editor
                            beforeMount={handleEditorWillMount}
                            onMount={handleEditorDidMount}
                            height={"100vh"} defaultLanguage="json" options={{ formatOnPaste: true, formatOnType: true, minimap: { enabled: false } }} defaultValue={JSON.stringify({
                                "pull_mode": "image",
                                "image": {
                                    "tag": "nginxdemos/hello"
                                },
                                "repo": {
                                    "url": "",
                                    "build_cmd": "",
                                    "run_cmd": ""
                                },
                                "settings": {
                                    "type": "webservice",
                                    "port": 80,
                                    "name": "testservice"
                                },
                                "loadmanager": {
                                    "machineType": "low",
                                    "maxInstances": 50,
                                    "minInstances": 1,
                                    "unitsPerInstance": 1,
                                    "autoScaleDown": false,
                                    "autoScaleUp": true,
                                    "terminationScript": "echo 'bye bye'"
                                },
                                "variables": {
                                    "MONGO_URL": "mongodb://mongo:27017/test"
                                }
                            })} />
                    </div>
                    {/* End main area */}
                </main>


                <aside className="relative hidden w-96 flex-shrink-0 overflow-y-auto border-r border-gray-200 xl:flex xl:flex-col">
                    {/* Start secondary column (hidden on smaller screens) */}
                    <div className="absolute inset-0 py-6 px-4 sm:px-6 lg:px-8">
                        <Button>Apply config</Button>
                    </div>
                    {/* End secondary column */}
                </aside>
            </div>
        </>
    )
}
