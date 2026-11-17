
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced376_0(value: int) -> int:
    return value + 43


def decorate376_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate376_0
def doubled376_0(value: int) -> int:
    return value + value
def traced376_1(value: int) -> int:
    return value + 50


def decorate376_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate376_1
def doubled376_1(value: int) -> int:
    return value + value
def traced376_2(value: int) -> int:
    return value + 57


def decorate376_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate376_2
def doubled376_2(value: int) -> int:
    return value + value
def guarded376_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project376_4(values):
    select = lambda item: item + 71
    return [select(item) for item in values if item != 0]
def guarded376_5(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced376_6(value: int) -> int:
    return value + 85


def decorate376_6(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate376_6
def doubled376_6(value: int) -> int:
    return value + value
def expand376_7(values):
    for item in values:
        if item > 92:
            yield item * 34


def collect376_7(values):
    return list(expand376_7(values))
