const name = document.getElementById("name")
const askButton = document.getElementById("ask-button")
const askTime = document.getElementById("ask-time")

askButton.addEventListener("click", function () {
   askTime.textContent = "new button clicked"
   let data = {
      username: name.value,
   };
   fetch("/post", {
      headers: {
         'Accept': 'application/json',
         'Content-Type': 'application/json'
      },
      method: "POST",
      body: JSON.stringify(data)
   }).then((response) => {
      response.text().then(function (data) {
         let result = JSON.parse(data);
         console.log(result)
      });
   }).catch((error) => {
      console.log(error)
   });
})