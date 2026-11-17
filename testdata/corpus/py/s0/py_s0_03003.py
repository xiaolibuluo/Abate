
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1375_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 32}
def project1375_1(values):
    select = lambda item: item + 39
    return [select(item) for item in values if item != 0]
def guarded1375_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand1375_3(values):
    for item in values:
        if item > 53:
            yield item * 29


def collect1375_3(values):
    return list(expand1375_3(values))
