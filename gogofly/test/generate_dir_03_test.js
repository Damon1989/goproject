"use strict";
(() => {
    let iRootNode = {};
    let stSeparator = "\\";
    const gstPaths = [];
    const loadJson = () => {
        return new Promise((resolve, reject) => {
            let xhr = new XMLHttpRequest();
            xhr.open("GET", "dir.json", true);
            xhr.onreadystatechange = () => {
                if (xhr.readyState === 4) {
                    if (xhr.status === 200) {
                        resolve(xhr.responseText);
                    }
                    else {
                        reject(new Error("Failed to load JSON file"));
                    }
                }
            };
            xhr.send();
        });
    };
    const parseNode = (iNode, stParentDir) => {
        if (iNode.text) {
            createDir(iNode, stParentDir);
        }
        if (stParentDir) {
            stParentDir += stSeparator;
        }
        iNode.text && (stParentDir += iNode.text);
        if (!iNode.children) {
            return;
        }
        for (const childNode of iNode.children || []) {
            parseNode(childNode, stParentDir);
        }
    };
    const createDir = (iNode, stParentDir) => {
        console.log("createDir");
        let path = iNode.text;
        if (stParentDir) {
            path = stParentDir + stSeparator + path;
        }
        gstPaths.push(path);
    };
    const generateBatFile = () => {
        let stBatContent = "@echo off\n";
        for (const path of gstPaths) {
            stBatContent += `mkdir "${path}"\n`;
        }
        console.log(stBatContent);
        const domA = document.createElement("a");
        domA.setAttribute("href", "data:text/plain;charset=utf-8," + encodeURIComponent(stBatContent));
        domA.setAttribute("download", "create_dirs.bat");
        domA.style.display = "none";
        document.body.appendChild(domA);
        domA.click();
        document.body.removeChild(domA);
    };
    loadJson().then((data) => {
        console.log("JSON loaded successfully");
        console.log(data);
        iRootNode = JSON.parse(data);
        parseNode(iRootNode, "");
        console.log(gstPaths);
        generateBatFile();
    }).catch((error) => {
        console.error("Error loading JSON:", error);
    });
    console.log("generate_dir_03_test");
})();
