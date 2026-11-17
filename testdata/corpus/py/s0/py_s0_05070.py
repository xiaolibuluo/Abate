
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize2359_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 20:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def traced2359_1(value: int) -> int:
    return value + 27


def decorate2359_1(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2359_1
def doubled2359_1(value: int) -> int:
    return value + value
async def fetch2359_2(gate):
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
def traced2359_3(value: int) -> int:
    return value + 41


def decorate2359_3(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate2359_3
def doubled2359_3(value: int) -> int:
    return value + value
def project2359_4(values):
    select = lambda item: item + 48
    return [select(item) for item in values if item != 0]
