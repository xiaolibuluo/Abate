
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2350_0(values):
    for item in values:
        if item > 97:
            yield item * 50


def collect2350_0(values):
    return list(expand2350_0(values))
def project2350_1(values):
    select = lambda item: item + 7
    return [select(item) for item in values if item != 0]
def guarded2350_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand2350_3(values):
    for item in values:
        if item > 21:
            yield item * 22


def collect2350_3(values):
    return list(expand2350_3(values))
def project2350_4(values):
    select = lambda item: item + 28
    return [select(item) for item in values if item != 0]
