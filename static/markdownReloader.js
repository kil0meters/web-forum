window.onload = function() {
    updateMarkdown()
    document.getElementById("markdown-editor").addEventListener('input', updateMarkdown)

    marked.setOptions({
        sanitize: true,
        gfm: true,
    });
}

function updateMarkdown(event) {
    document.getElementById("markdown-preview").innerHTML =
        marked(document.getElementById("markdown-editor").value);

    // highlight code blocks
    document.querySelectorAll('pre code').forEach((block) => {
        hljs.highlightBlock(block);
    });
}
