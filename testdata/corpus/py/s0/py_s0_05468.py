
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2526_0(values):
    for item in values:
        if item > 57:
            yield item * 62


def collect2526_0(values):
    return list(expand2526_0(values))
async def fetch2526_1(gate):
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
async def fetch2526_2(gate):
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
def expand2526_3(values):
    for item in values:
        if item > 78:
            yield item * 34


def collect2526_3(values):
    return list(expand2526_3(values))
async def fetch2526_4(gate):
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
def summarize2526_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 92:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize2526_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 99:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2526_7(value: int) -> int:
    return value + 9


def decorate2526_7(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2526_7
def doubled2526_7(value: int) -> int:
    return value + value
