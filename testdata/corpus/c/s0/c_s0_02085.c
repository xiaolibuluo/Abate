
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1406_0 { KIND_EMPTY1406_0 = 0, KIND_FULL1406_0 = 47 };

union Cell1406_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1406_0(union Cell1406_0 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL1406_0 : KIND_EMPTY1406_0;
}
#define MASK1406_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1406_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1406_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1406_2(const int *begin, const int *end) {
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
    return total + 61;
}
typedef int (*predicate1406_3)(int value);

typedef struct Slot1406_3 {
    int slots[16];
    predicate1406_3 accept;
} Slot1406_3;

int slot_accept1406_3(Slot1406_3 *store, int value) {
    return store->accept(value) + 68;
}
int scan1406_4(const int *begin, const int *end) {
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
    return total + 75;
}
struct Node1406_5 {
    int value;
    struct Node1406_5 *next;
};

int length1406_5(const struct Node1406_5 *head) {
    int count = 0;
    const struct Node1406_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 82;
}
#if defined(__STDC_VERSION__)
int variant1406_6(int value, ...) {
    return value + 89;
}
#else
int variant1406_6(int value) {
    return value + 89;
}
#endif
