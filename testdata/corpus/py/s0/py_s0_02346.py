
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1078_0(value: int) -> int:
    return value + 51


def decorate1078_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1078_0
def doubled1078_0(value: int) -> int:
    return value + value
def project1078_1(values):
    select = lambda item: item + 58
    return [select(item) for item in values if item != 0]
def traced1078_2(value: int) -> int:
    return value + 65


def decorate1078_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1078_2
def doubled1078_2(value: int) -> int:
    return value + value
