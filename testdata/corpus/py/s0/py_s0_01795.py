
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded821_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def expand821_1(values):
    for item in values:
        if item > 15:
            yield item * 33


def collect821_1(values):
    return list(expand821_1(values))
def expand821_2(values):
    for item in values:
        if item > 22:
            yield item * 44


def collect821_2(values):
    return list(expand821_2(values))
def traced821_3(value: int) -> int:
    return value + 29


def decorate821_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate821_3
def doubled821_3(value: int) -> int:
    return value + value
def summarize821_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 36:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch821_5(gate):
    results: list[str] = []
    try:
        async with gate:
            await gate.wait()
            results.append("ready")
    except RuntimeError:
        raise
    finally:
        results.clear()
    return results
def project821_6(values):
    select = lambda item: item + 50
    return [select(item) for item in values if item != 0]
def expand821_7(values):
    for item in values:
        if item > 57:
            yield item * 38


def collect821_7(values):
    return list(expand821_7(values))
def project821_8(values):
    select = lambda item: item + 64
    return [select(item) for item in values if item != 0]
async def fetch821_9(gate):
    results: list[str] = []
    try:
        async with gate:
            await gate.wait()
            results.append("ready")
    except RuntimeError:
        raise
    finally:
        results.clear()
    return results
