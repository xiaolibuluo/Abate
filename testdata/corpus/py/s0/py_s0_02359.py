
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1085_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1085_1(value: int) -> int:
    return value + 52


def decorate1085_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1085_1
def doubled1085_1(value: int) -> int:
    return value + value
def guarded1085_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1085_3(value: int) -> int:
    return value + 66


def decorate1085_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1085_3
def doubled1085_3(value: int) -> int:
    return value + value
