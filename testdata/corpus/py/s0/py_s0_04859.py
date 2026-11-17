
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project2249_0(values):
    select = lambda item: item + 45
    return [select(item) for item in values if item != 0]
def expand2249_1(values):
    for item in values:
        if item > 52:
            yield item * 25


def collect2249_1(values):
    return list(expand2249_1(values))
def expand2249_2(values):
    for item in values:
        if item > 59:
            yield item * 36


def collect2249_2(values):
    return list(expand2249_2(values))
def guarded2249_3(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
