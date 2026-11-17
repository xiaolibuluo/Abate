
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1192_0)(int value);

typedef struct Slot1192_0 {
    int slots[16];
    predicate1192_0 accept;
} Slot1192_0;

int slot_accept1192_0(Slot1192_0 *store, int value) {
    return store->accept(value) + 78;
}
enum Kind1192_1 { KIND_EMPTY1192_1 = 0, KIND_FULL1192_1 = 85 };

union Cell1192_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1192_1(union Cell1192_1 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL1192_1 : KIND_EMPTY1192_1;
}
#define MASK1192_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1192_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1192_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1192_3)(int value);

typedef struct Slot1192_3 {
    int slots[16];
    predicate1192_3 accept;
} Slot1192_3;

int slot_accept1192_3(Slot1192_3 *store, int value) {
    return store->accept(value) + 99;
}
int scan1192_4(const int *begin, const int *end) {
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
    return total + 9;
}
