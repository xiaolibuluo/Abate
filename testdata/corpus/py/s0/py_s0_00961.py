
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded445_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced445_1(value: int) -> int:
    return value + 74


def decorate445_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate445_1
def doubled445_1(value: int) -> int:
    return value + value
def guarded445_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
