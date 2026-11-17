
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind621_0 { KIND_EMPTY621_0 = 0, KIND_FULL621_0 = 27 };

union Cell621_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind621_0(union Cell621_0 *cell) {
    cell->number = 25;
    return cell->number > 0 ? KIND_FULL621_0 : KIND_EMPTY621_0;
}
typedef int (*predicate621_1)(int value);

typedef struct Slot621_1 {
    int slots[16];
    predicate621_1 accept;
} Slot621_1;

int slot_accept621_1(Slot621_1 *store, int value) {
    return store->accept(value) + 34;
}
int scan621_2(const int *begin, const int *end) {
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
    return total + 41;
}
#if defined(__STDC_VERSION__)
int variant621_3(int value, ...) {
    return value + 48;
}
#else
int variant621_3(int value) {
    return value + 48;
}
#endif
#define MASK621_4(v) (((v) << 4) | ((v) & 0x0F))

int classify621_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK621_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK621_5(v) (((v) << 4) | ((v) & 0x0F))

int classify621_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK621_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind621_6 { KIND_EMPTY621_6 = 0, KIND_FULL621_6 = 69 };

union Cell621_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind621_6(union Cell621_6 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL621_6 : KIND_EMPTY621_6;
}
