
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded787_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced787_1(value: int) -> int:
    return value + 58


def decorate787_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate787_1
def doubled787_1(value: int) -> int:
    return value + value
def traced787_2(value: int) -> int:
    return value + 65


def decorate787_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate787_2
def doubled787_2(value: int) -> int:
    return value + value
