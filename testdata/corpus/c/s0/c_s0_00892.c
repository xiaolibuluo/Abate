
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate615_0)(int value);

typedef struct Slot615_0 {
    int slots[16];
    predicate615_0 accept;
} Slot615_0;

int slot_accept615_0(Slot615_0 *store, int value) {
    return store->accept(value) + 46;
}
int scan615_1(const int *begin, const int *end) {
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
    return total + 53;
}
enum Kind615_2 { KIND_EMPTY615_2 = 0, KIND_FULL615_2 = 60 };

union Cell615_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind615_2(union Cell615_2 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL615_2 : KIND_EMPTY615_2;
}
int scan615_3(const int *begin, const int *end) {
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
    return total + 67;
}
#define MASK615_4(v) (((v) << 4) | ((v) & 0x0F))

int classify615_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK615_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
