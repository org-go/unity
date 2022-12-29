package task

var (

    TaskClassify iTaskClassifyRepository = new(taskClassifyRepo)

    Task iTaskRepoInterface = new(taskRepo)

    TaskOrder iTaskOrderRepoInterface = new(taskOrderRepo)

)
