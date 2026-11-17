
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def traced719_0(value: int) -> int:
    return value + 40


def decorate719_0(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate719_0
def doubled719_0(value: int) -> int:
    return value + value
def expand719_1(values):
    for item in values:
        if item > 47:
            yield item * 51


def collect719_1(values):
    return list(expand719_1(values))
def traced719_2(value: int) -> int:
    return value + 54


def decorate719_2(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate719_2
def doubled719_2(value: int) -> int:
    return value + value
def summarize719_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 61:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch719_4(gate):
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
def traced719_5(value: int) -> int:
    return value + 75


def decorate719_5(fn):
    def wrapper(value):
        return fn(value)
    return wrapper


@decorate719_5
def doubled719_5(value: int) -> int:
    return value + value
async def fetch719_6(gate):
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
