var cy = window.cy = cytoscape({
        container: document.getElementById('cy'),

        boxSelectionEnabled: false,

        style: [{
                selector: 'node',
                css: {
                        'content': 'data(label)',
                        'text-valign': 'center',
                        'text-halign': 'center',
                        'shape': 'data(shape)',
                        'padding': 'data(padding)',
                        'background-color': 'white',
                        'border-style': 'data(borderStyle)',
                        'border-color': 'black',
                        'border-width': 0.9,
                        'font-size': 'data(fontSize)',
                        'text-margin-y': 'data(textMarginY)',
                        'font-weight': 'data(fontWeight)',
                        'width': 'data(width)'
                        }
                },
                {
                        selector: ':parent',
                        css: {
                                'text-valign': 'top',
                                'text-halign': 'center',
                                'border-style': 'data(borderStyle)',
                                'border-color': 'data(borderColor)',
                                'background-color': 'white'
                        }
                },
                {
                        selector: 'edge',
                        css: {
                                'text-wrap': 'wrap',
                                'curve-style': 'unbundled-bezier',
                                'target-arrow-shape': 'triangle',
                                'target-arrow-color': 'black',
                                'width': 1,
                                'line-style': 'data(lineStyle)'
                        }
                },
                {
                        selector: "edge[label]",
                        css: {
                                "label": "data(label)",
                                "text-rotation": "autorotate",
                                'font-weight': 'data(fontWeight)',
                                'font-size': 'data(fontSize)',
                                'text-valign': 'top',
                        }
                },
        ],

        elements: {
                nodes: [
                        { data: { id: 'a',fontSize: 12,   textMarginY: 0, padding: 35, label: "Node App", parent: 'b' }, position: { x: 215, y: 85 } },
                        { data: { id: 'b',fontSize: 12,   textMarginY: -20, padding: 35, label: "Web Service", borderStyle: "dashed", borderColor: "black" } },
                        { data: { id: 'c',fontSize: 11,   textMarginY: 0, padding: 30, label: "External API", borderStyle: "dashed"} },
                        { data: { id: 'd',fontSize: 12,   textMarginY: 0, padding: 35,  label: "Browser", width: 100, shape: "rectangle" }, position: { x: 215, y: 175 } },
                        { data: { id: 'e',fontSize: 9,   textMarginY: 0, padding: 30, label: "Static front-end files", borderStyle: "dashed"} },
                ],
                edges: [
                        { data: { id: 'ad', source: 'a', target: 'd', label: "Web Response", fontSize: 11, lineStyle: '' } },
                        { data: { id: 'da', source: 'd', target: 'a',label: "Web Request",  fontSize: 11, lineStyle: ''} },
                        { data: { id: 'cd', source: 'c', target: 'd', label: 'API Request', fontSize: 11, lineStyle: 'dashed' } },
                        { data: { id: 'dc', source: 'd', target: 'c', label: 'API Response', fontSize: 11, lineStyle: 'dashed' } },
                        { data: { id: 'de', source: 'd', target: 'e', label: '', fontSize: 11, lineStyle: '' } },
                        { data: { id: 'ed', source: 'e', target: 'd', label: '', fontSize: 11, lineStyle: '' } },
                        { data: { id: 'ac', source: 'a', target: 'c', label: 'API Request', fontSize: 11, lineStyle: 'dashed' } },
                        { data: { id: 'ca', source: 'c', target: 'a', label: 'API Response', fontSize: 11, lineStyle: 'dashed' } },

                ]
        },

        layout: {
                name: 'preset',
                padding: 150
        }
});