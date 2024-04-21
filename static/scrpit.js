document.getElementById('uploadForm').addEventListener('submit', function(event) {
    event.preventDefault();
    let formData = new FormData();
    formData.append('videoFile', document.getElementById('videoFile').files[0]);

    fetch('/upload', {
        method: 'POST',
        body: formData
    }).then(response => {
        // Handle response from the server
        console.log("Upload succesfull");
    }).catch(error => {
        // Handle error
        console.log("Upload error");
    });
});
