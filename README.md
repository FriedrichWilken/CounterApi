This is just a simple rest API to count things. Data storage is in memory and will be lost after closing the application.

## Get required packages

To get the required packages run

```bash
make get
```

## Run the tests

To run the tests execute from the folder "counter" run:

```bash
make test
```

## Run the API

To run the API execute from the folder "counter" run:

```bash
make run
```

## Using the API

There are three endpoints:

#### Get the current value

Use the ```GET``` method at ```http://localhost:8080/counter``` to fetch the current counter value:

```json
{
    "counter": 17
}
```

#### Increase the Value

Use the ```PUT``` method at ```http://localhost:8080/counter/increase``` to increase the counter value. If you send no body the value will be increased by 1. If you send a body with a value, the counter will be increased by the given value:

```json
{
	"increaseBy":5
}
```

as a response you will get the new current value.

#### Decrease the Value

Use the ```PUT``` method at ```http://localhost:8080/counter/decrease``` to decrease the counter value. If you send no body the value will be decreased by 1. If you send a body with a value, the counter will be increased by the given value:

```json
{
	"decreaseBy":6
}
```

as a response you will get the new current value.



​	

