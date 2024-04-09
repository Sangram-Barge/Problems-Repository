
function createTable(data) {
  var list = document.getElementById("problems")
  data.forEach((problem) => {
    var tdId = document.createElement("td")
    tdId.innerText = problem.id
    var tdProblem = document.createElement("td")
    tdProblem.innerText = problem.problem
    var tdPlatform = document.createElement("td")
    tdPlatform.innerText = problem.platform
    var tdDescription = document.createElement("td")
    tdDescription.innerText = problem.desc
    var tdIntiution = document.createElement("td")
    tdIntiution.innerText = problem.intiution

    var tdLink = document.createElement("td")
    tdLink.innerHTML = `<a href=${problem.link} target="_blank">Link</a>`

    var tr = document.createElement("tr")
    tr.appendChild(tdId)
    tr.appendChild(tdProblem)
    tr.appendChild(tdPlatform)
    tr.appendChild(tdDescription)
    tr.appendChild(tdIntiution)
    tr.appendChild(tdLink)

    document.getElementById("problems")
      .appendChild(tr)
  });

}
function populateProblems() {
  clear()
  fetch(`http://localhost:9090/problems`)
    .then(response => response.json())
    .then(data => {
      createTable(data)
    })
}
function populateKeywordProblems() {
  clear()
  let key = document.getElementById("fKey").value
  fetch(`http://localhost:9090/problems/find?keyword=${key}`)
    .then(response => response.json())
    .then(data => {
      createTable(data)
    })
}
function clear() {
  let pr = document.getElementById("problems")
  pr.innerHTML = ""
}
async function insertProblem() {
  let problem = {
    "problem": document.getElementById("prob").value,
    "platform": document.getElementById("plat").value,
    "desc": document.getElementById("desc").value,
    "intiution": document.getElementById("intiu").value,
    "link": document.getElementById("link").value,
  }
  await fetch("http://localhost:9090/problems/add",
    {
      method: "POST",
      mode: "no-cors",
      body: JSON.stringify(problem),
      headers: {
        "Content-Type": "application/json",
      }
    })
  window.location.reload()
}
async function deleteProblem() {
  let id = document.getElementById("dId").value
  await fetch(`http://localhost:9090/problem?id=${id}`,
    {
      method: "DELETE",
    })
  window.location.reload()
}
populateProblems()

