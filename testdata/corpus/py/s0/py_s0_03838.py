
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1771_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded1771_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced1771_2(value: int) -> int:
    return value + 53


def decorate1771_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1771_2
def doubled1771_2(value: int) -> int:
    return value + value
def expand1771_3(values):
    for item in values:
        if item > 60:
            yield item * 56


def collect1771_3(values):
    return list(expand1771_3(values))
def guarded1771_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
