let blogs = []

// alert project page
function formAlert() {
  let title = document.getElementById("title").value;
  let startDate = document.getElementById("startDate").value;
  let endDate = document.getElementById("endDate").value;
  let projectDescription = document.getElementById("projectDescription").value;
  let reactJS = document.getElementById("reactJS").value;
  let nodeJS = document.getElementById("nodeJS").value;
  let nextJS = document.getElementById("nextJS").value;
  let typeScript = document.getElementById("typeScript").value;
  let imageProject = document.getElementById("imageProject").value;
  
  if(title == "") {
      return alert("Input the title of project, please!");
  } else if(startDate == "") {
      return alert("Input start of Project's date, please!");
  } else if(endDate == "") {
      return alert("Input end of Project's date, please!");
  } else if(projectDescription == "") {
      return alert("Input description, please!");
  } else if( reactJS.checked == "" || nodeJS.checked == "" || nextJS.checked == "" || typeScript.checked == "") {
      return alert("Choose technology, please!");
  } else if(imageProject == "") {
      return alert("Inpu an image, please!");
  }
};

// event time
function getBlog(event){
    event.preventDefault()

    function getDistanceTime(){
      let diff = new Date(endDate) - new Date(startDate);

      let days = Math.floor(diff / (24 * 60 * 60 * 1000));
      let months = Math.floor(days / 30);
      let years = Math.floor(months / 12);
      let remainingDays = days % 30;
      let remainingMonths = months % 12;
      let daysAffix = `$`
    
      if (years > 0 && remainingMonths > 0 && remainingDays > 0) {
          return `${years} Years ${remainingMonths} Months ${remainingDays} Days`;
      } else if (years > 0 && remainingMonths > 0 && remainingDays == 0){
          return `${years} Years ${remainingMonths} Months`;
      } else if (years > 0 && remainingMonths == 0 && remainingDays == 0){
          return `${years} Years`;
      } else if (years > 0 && remainingMonths == 0 && remainingDays > 0){
          return `${years} Years ${remainingDays} Days`;
      } else if (years == 0 && remainingMonths > 0 && remainingDays > 0){
          return `${remainingMonths} Months ${remainingDays} Days`;
      } else if (years == 0 && remainingMonths > 0 && remainingDays == 0){
          return `${remainingMonths} Months`;
      } else if (years == 0 && remainingMonths == 0 && remainingDays > 0){
          return `${remainingDays} Days`;
      } 
  }

// input form project
    let title = document.getElementById("title").value
    let startDate = document.getElementById("startDate").value
    let endDate = document.getElementById("endDate").value
    let projectDescription = document.getElementById("projectDescription").value
    let imageProject = document.getElementById("imageProject").files
    let duration = getDistanceTime()

    const nodeJS = `<img src="images/nodejs.png" alt="nodejs">`;
    const nextJS = `<img src="images/reactjs.png" alt="reactjs">`;
    const reactJS = `<img src="images/reactjs.png" alt="reactjs">`;
    const typeScript = `<img src="images/typescript.png" alt="typescript">`;
    
    let nodeJsIMG = document.getElementById("nodeJS").checked ? nodeJS : "";
    let nextJsIMG = document.getElementById("nextJS").checked ? nextJS : "";
    let typeScriptIMG = document.getElementById("typeScript").checked ? typeScript : "";
    let reactJsIMG = document.getElementById("reactJS").checked ? reactJS : "";
    
    imageURL = URL.createObjectURL(imageProject[0])

    let projects = {
        title,
        duration,
        projectDescription,
        nodeJsIMG,
        nextJsIMG,
        typeScriptIMG,
        reactJsIMG,
        imageURL,
        author: "Nafiisan N. Achmad",
    }

    blogs.push(projects);    
    renderBlog();
}

function renderBlog(){
  document.getElementById("project").innerHTML = "";

  for(let i = 0; i < blogs.length; i++) {
    document.getElementById("project").innerHTML += `
      <div class="card rounded-4 border-0 shadow-sm ppc" style="width: 16rem;">
      <img src="${blogs[i].imageURL}" class="card-img-top rounded-top-4"
      style="object-fit: cover;"/>
      <h2 href="project-detail.html" class="card-title ppc-title"
        style="text-decoration: none;" target="_blank">${blogs[i].title}
      </h2>
      <div class="card-body">
        <div class="detailDate">
        ${blogs[i].duration} 
        </div>
        <div class="author">
          <p>${blogs[i].author}</p>
        </div> 
        <p class="card-text lh-sm" style="font-size:small; overflow: hidden;" >
          ${blogs[i].projectDescription}
        </p>
      </div>
      <div class="technology">
      ${blogs[i].nodeJsIMG}
      ${blogs[i].reactJsIMG}
      ${blogs[i].nextJsIMG}
      ${blogs[i].typeScriptIMG}
      </div>
      <div class="d-flex flex-row gap-3">
          <button class="btn rounded-pill btn-outline-secondary btn-sm w-50">Edit</button>
          <button class="btn rounded-pill btn-outline-danger btn-sm w-50">Delete</button>
      </div>
    </div>
    `
  }
    
}


