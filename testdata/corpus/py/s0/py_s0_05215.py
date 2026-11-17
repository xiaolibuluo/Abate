
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2427_0(gate):
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
def guarded2427_1(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
def guarded2427_2(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
async def fetch2427_3(gate):
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
def project2427_4(values):
    select = lambda item: item + 59
    return [select(item) for item in values if item != 0]
def expand2427_5(values):
    for item in values:
        if item > 66:
            yield item * 34


def collect2427_5(values):
    return list(expand2427_5(values))
def summarize2427_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 73:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand2427_7(values):
    for item in values:
        if item > 80:
            yield item * 56


def collect2427_7(values):
    return list(expand2427_7(values))
def traced2427_8(value: int) -> int:
    return value + 87


def decorate2427_8(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2427_8
def doubled2427_8(value: int) -> int:
    return value + value
