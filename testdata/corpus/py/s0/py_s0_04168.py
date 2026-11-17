
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def guarded1943_0(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def summarize1943_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 51:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project1943_2(values):
    select = lambda item: item + 58
    return [select(item) for item in values if item != 0]
def expand1943_3(values):
    for item in values:
        if item > 65:
            yield item * 40


def collect1943_3(values):
    return list(expand1943_3(values))
def expand1943_4(values):
    for item in values:
        if item > 72:
            yield item * 51


def collect1943_4(values):
    return list(expand1943_4(values))
def project1943_5(values):
    select = lambda item: item + 79
    return [select(item) for item in values if item != 0]
def project1943_6(values):
    select = lambda item: item + 86
    return [select(item) for item in values if item != 0]
def guarded1943_7(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch1943_8(gate):
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
def traced1943_9(value: int) -> int:
    return value + 10


def decorate1943_9(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1943_9
def doubled1943_9(value: int) -> int:
    return value + value
def expand1943_10(values):
    for item in values:
        if item > 17:
            yield item * 56


def collect1943_10(values):
    return list(expand1943_10(values))
