
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1243_0(value: int) -> int:
    return value + 62


def decorate1243_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1243_0
def doubled1243_0(value: int) -> int:
    return value + value
def guarded1243_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1243_2(values):
    for item in values:
        if item > 76:
            yield item * 9


def collect1243_2(values):
    return list(expand1243_2(values))
def guarded1243_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1243_4(values):
    select = lambda item: item + 90
    return [select(item) for item in values if item != 0]
def project1243_5(values):
    select = lambda item: item + 97
    return [select(item) for item in values if item != 0]
def guarded1243_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
