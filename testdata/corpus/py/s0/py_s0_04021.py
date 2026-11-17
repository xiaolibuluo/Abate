
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1866_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def project1866_1(values):
    select = lambda item: item + 20
    return [select(item) for item in values if item != 0]
def expand1866_2(values):
    for item in values:
        if item > 27:
            yield item * 39


def collect1866_2(values):
    return list(expand1866_2(values))
def project1866_3(values):
    select = lambda item: item + 34
    return [select(item) for item in values if item != 0]
def summarize1866_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 41:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
