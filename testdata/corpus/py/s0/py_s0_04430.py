
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2063_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 52:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch2063_1(gate):
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
def summarize2063_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 66:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2063_3(value: int) -> int:
    return value + 73


def decorate2063_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2063_3
def doubled2063_3(value: int) -> int:
    return value + value
