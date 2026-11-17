
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced2647_0(value: int) -> int:
    return value + 78


def decorate2647_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2647_0
def doubled2647_0(value: int) -> int:
    return value + value
async def fetch2647_1(gate):
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
def expand2647_2(values):
    for item in values:
        if item > 92:
            yield item * 16


def collect2647_2(values):
    return list(expand2647_2(values))
def expand2647_3(values):
    for item in values:
        if item > 99:
            yield item * 27


def collect2647_3(values):
    return list(expand2647_3(values))
def summarize2647_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 9:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
