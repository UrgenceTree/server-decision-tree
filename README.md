# Server Decision Tree

## **Prerequisites**

### *What we need ?*
For this project you will need to have **Go** installed on your computer. If it's not already done, you can find the installation [here](https://go.dev/doc/install).

---

## **How to build** :

### *With an executable file*
First of all you need to build the project using the following line:
```sh
    go build -o my_script ./convertion/questions.go
```

Now you can launch our project:
```sh
	./my_script
```

### *Withouut an executable file*
You can also run the project directly throught a specific command with go tool :
```sh
    go run ./convertion/questions.go
```
- it's more ease to run the project in this way and we can avoid the process of the build

We can now access interact with  our script and communicate through the getline.

---

### **Test the project** :

In order to test our project, you can run the following command. It will test the time to execute the python and go parts and also compare them each other.

For this, you have to go in the directory "unit_test/time_test":
```sh
    cd unit_test/time_test
```

Then you can launch the program that will test our project :
```sh
    ./time_test.sh
```

- Note : If the file don't the have the right to execute commands, you cand provide it with the following line :
    ```sh
        chmod a+x time_test.sh
    ```
