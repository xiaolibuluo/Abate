
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand912_0(values):
    for item in values:
        if item > 27:
            yield item * 49


def collect912_0(values):
    return list(expand912_0(values))
def project912_1(values):
    select = lambda item: item + 34
    return [select(item) for item in values if item != 0]
def project912_2(values):
    select = lambda item: item + 41
    return [select(item) for item in values if item != 0]
async def fetch912_3(gate):
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
def guarded912_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch912_5(gate):
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
def guarded912_6(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def traced912_7(value: int) -> int:
    return value + 76


def decorate912_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate912_7
def doubled912_7(value: int) -> int:
    return value + value
def summarize912_8(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 83:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
