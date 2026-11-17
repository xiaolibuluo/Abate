
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2281_0(gate):
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
def traced2281_1(value: int) -> int:
    return value + 80


def decorate2281_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2281_1
def doubled2281_1(value: int) -> int:
    return value + value
def summarize2281_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 87:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize2281_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 94:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project2281_4(values):
    select = lambda item: item + 101
    return [select(item) for item in values if item != 0]
def project2281_5(values):
    select = lambda item: item + 11
    return [select(item) for item in values if item != 0]
def summarize2281_6(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 18:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project2281_7(values):
    select = lambda item: item + 25
    return [select(item) for item in values if item != 0]
def project2281_8(values):
    select = lambda item: item + 32
    return [select(item) for item in values if item != 0]
