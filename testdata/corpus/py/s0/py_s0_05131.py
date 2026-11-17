
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2390_0(values):
    select = lambda item: item + 35
    return [select(item) for item in values if item != 0]
def traced2390_1(value: int) -> int:
    return value + 42


def decorate2390_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2390_1
def doubled2390_1(value: int) -> int:
    return value + value
def traced2390_2(value: int) -> int:
    return value + 49


def decorate2390_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2390_2
def doubled2390_2(value: int) -> int:
    return value + value
