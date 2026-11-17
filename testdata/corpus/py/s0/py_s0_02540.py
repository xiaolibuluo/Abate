
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1171_0(value: int) -> int:
    return value + 96


def decorate1171_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1171_0
def doubled1171_0(value: int) -> int:
    return value + value
def guarded1171_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
