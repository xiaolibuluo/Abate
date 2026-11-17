
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2588_0(value: int) -> int:
    return value + 87


def decorate2588_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2588_0
def doubled2588_0(value: int) -> int:
    return value + value
def guarded2588_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced2588_2(value: int) -> int:
    return value + 101


def decorate2588_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2588_2
def doubled2588_2(value: int) -> int:
    return value + value
