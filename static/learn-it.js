let LearnIt = {
  editor: {
    textAreaElement: null,
    originalValue: "",
    codeMirror: null,
  },
  canGo: false,
  setup: () => {
    LearnIt.setupEditor()
    LearnIt.parseLesson()
    LearnIt.emit("ready")
  },
  parseLesson: () => {
    let lessonView = document.getElementById("LessonView")
    let preTags = Array.prototype.slice.call(lessonView.getElementsByTagName("pre"), 0)
    for (preTag of preTags) {
      let divTag = document.createElement("div")
      divTag.className = "codeline cm-s-base16-dark"
      CodeMirror.runMode(preTag.innerText, "diff", divTag)
      preTag.parentNode.replaceChild(divTag, preTag)
    }
    // Now search for code tags
    let codeTags = Array.prototype.slice.call(lessonView.getElementsByTagName("code"), 0)
    for (codeTag of codeTags) {
      let spanTag = document.createElement("span")
      spanTag.className = "cm-s-base16-dark"
      CodeMirror.runMode(codeTag.innerText, "go", spanTag)
      codeTag.parentNode.replaceChild(spanTag, codeTag)
    }
  },
  setupEditor: () => {
    LearnIt.editor.textArea = document.getElementById("Editor")
    if (!LearnIt.editor.textArea) return
    LearnIt.editor.originalValue = LearnIt.editor.textArea.value
    LearnIt.editor.codeMirror = CodeMirror.fromTextArea(LearnIt.editor.textArea, {
        theme: "base16-dark",
        lineNumbers: true,
        lineWrapping: true,
        indentUnit: 2,
        tabSize: 2,
        indentWithTabs: true,
        readOnly: true,
        mode: "go"
    })
    LearnIt.editor.codeMirror.on("change", function(cm, change) {
        LearnIt.editor.textArea.value = cm.getValue()
    })
  }
}
LearnIt.go = function() {
  LearnIt.canGo = true
}
LearnIt._listeners = {}
LearnIt.emit = function(evt, payload) {
  if (!LearnIt._listeners[evt]) { return }
  for (let i in LearnIt._listeners[evt]) {
    LearnIt._listeners[evt][i](payload)
  }
}
LearnIt.on = function(evt, cb) {
  if (!LearnIt._listeners[evt]) {
    LearnIt._listeners[evt] = []
  }
  LearnIt._listeners[evt].push(cb)
}
LearnIt.off = function(evt, cb) {
  if (!LearnIt._listeners[evt]) { return }
  for (let i = 0; i < LearnIt._listeners[evt].length; i++) {
    if (LearnIt._listeners[evt][i] == cb) {
      LearnIt._listeners[evt].splice(i, 1)
    }
  }
}

window.addEventListener("DOMContentLoaded", () => {
  LearnIt.setup()
})
