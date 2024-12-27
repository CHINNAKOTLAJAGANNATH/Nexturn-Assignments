document.getElementById('addTaskBtn').addEventListener('click', function() {
    const taskInput = document.getElementById('taskInput');
    const taskValue = taskInput.value.trim();

    if (taskValue !== "") {
        
        const newTask = document.createElement('li');
        newTask.innerHTML = `${taskValue} <button class="removeBtn">Remove</button>`;

       
        document.getElementById('taskList').appendChild(newTask);

      
        taskInput.value = "";
        newTask.querySelector('.removeBtn').addEventListener('click', function() {
        newTask.remove();
        });
    }
});
