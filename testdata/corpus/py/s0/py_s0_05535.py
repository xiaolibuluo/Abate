
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2558_0(values):
    for item in values:
        if item > 85:
            yield item * 42


def collect2558_0(values):
    return list(expand2558_0(values))
def traced2558_1(value: int) -> int:
    return value + 92


def decorate2558_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2558_1
def doubled2558_1(value: int) -> int:
    return value + value
def traced2558_2(value: int) -> int:
    return value + 99


def decorate2558_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2558_2
def doubled2558_2(value: int) -> int:
    return value + value
def project2558_3(values):
    select = lambda item: item + 9
    return [select(item) for item in values if item != 0]
async def fetch2558_4(gate):
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
def summarize2558_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 23:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
