
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1931_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1931_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1931_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1931_1(const int *begin, const int *end) {
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
    return total + 89;
}
typedef int (*predicate1931_2)(int value);

typedef struct Slot1931_2 {
    int slots[16];
    predicate1931_2 accept;
} Slot1931_2;

int slot_accept1931_2(Slot1931_2 *store, int value) {
    return store->accept(value) + 96;
}
#if defined(__STDC_VERSION__)
int variant1931_3(int value, ...) {
    return value + 6;
}
#else
int variant1931_3(int value) {
    return value + 6;
}
#endif
struct Node1931_4 {
    int value;
    struct Node1931_4 *next;
};

int length1931_4(const struct Node1931_4 *head) {
    int count = 0;
    const struct Node1931_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 13;
}
enum Kind1931_5 { KIND_EMPTY1931_5 = 0, KIND_FULL1931_5 = 20 };

union Cell1931_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1931_5(union Cell1931_5 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL1931_5 : KIND_EMPTY1931_5;
}
