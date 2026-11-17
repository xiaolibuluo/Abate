
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1079_0(value: int) -> int:
    return value + 64


def decorate1079_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1079_0
def doubled1079_0(value: int) -> int:
    return value + value
def traced1079_1(value: int) -> int:
    return value + 71


def decorate1079_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1079_1
def doubled1079_1(value: int) -> int:
    return value + value
def guarded1079_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
