
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced1383_0(value: int) -> int:
    return value + 39


def decorate1383_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate1383_0
def doubled1383_0(value: int) -> int:
    return value + value
def project1383_1(values):
    select = lambda item: item + 46
    return [select(item) for item in values if item != 0]
def summarize1383_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 53:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1383_3(values):
    for item in values:
        if item > 60:
            yield item * 24


def collect1383_3(values):
    return list(expand1383_3(values))
def project1383_4(values):
    select = lambda item: item + 67
    return [select(item) for item in values if item != 0]
async def fetch1383_5(gate):
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
