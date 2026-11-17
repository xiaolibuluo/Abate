
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1899_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1899_1(value: int) -> int:
    return value + 61


def decorate1899_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1899_1
def doubled1899_1(value: int) -> int:
    return value + value
