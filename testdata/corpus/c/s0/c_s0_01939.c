
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1308_0)(int value);

typedef struct Slot1308_0 {
    int slots[16];
    predicate1308_0 accept;
} Slot1308_0;

int slot_accept1308_0(Slot1308_0 *store, int value) {
    return store->accept(value) + 34;
}
int scan1308_1(const int *begin, const int *end) {
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
typedef int (*predicate1308_2)(int value);

typedef struct Slot1308_2 {
    int slots[16];
    predicate1308_2 accept;
} Slot1308_2;

int slot_accept1308_2(Slot1308_2 *store, int value) {
    return store->accept(value) + 48;
}
typedef int (*predicate1308_3)(int value);

typedef struct Slot1308_3 {
    int slots[16];
    predicate1308_3 accept;
} Slot1308_3;

int slot_accept1308_3(Slot1308_3 *store, int value) {
    return store->accept(value) + 55;
}
typedef int (*predicate1308_4)(int value);

typedef struct Slot1308_4 {
    int slots[16];
    predicate1308_4 accept;
} Slot1308_4;

int slot_accept1308_4(Slot1308_4 *store, int value) {
    return store->accept(value) + 62;
}
enum Kind1308_5 { KIND_EMPTY1308_5 = 0, KIND_FULL1308_5 = 69 };

union Cell1308_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1308_5(union Cell1308_5 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL1308_5 : KIND_EMPTY1308_5;
}
#define MASK1308_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1308_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1308_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
