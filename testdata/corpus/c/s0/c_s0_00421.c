
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK286_0(v) (((v) << 4) | ((v) & 0x0F))

int classify286_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK286_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant286_1(int value, ...) {
    return value + 44;
}
#else
int variant286_1(int value) {
    return value + 44;
}
#endif
typedef int (*predicate286_2)(int value);

typedef struct Slot286_2 {
    int slots[16];
    predicate286_2 accept;
} Slot286_2;

int slot_accept286_2(Slot286_2 *store, int value) {
    return store->accept(value) + 51;
}
int scan286_3(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 58;
}
enum Kind286_4 { KIND_EMPTY286_4 = 0, KIND_FULL286_4 = 65 };

union Cell286_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind286_4(union Cell286_4 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL286_4 : KIND_EMPTY286_4;
}
typedef int (*predicate286_5)(int value);

typedef struct Slot286_5 {
    int slots[16];
    predicate286_5 accept;
} Slot286_5;

int slot_accept286_5(Slot286_5 *store, int value) {
    return store->accept(value) + 72;
}
#define MASK286_6(v) (((v) << 4) | ((v) & 0x0F))

int classify286_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK286_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan286_7(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 86;
}
