
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2329_0)(int value);

typedef struct Slot2329_0 {
    int slots[16];
    predicate2329_0 accept;
} Slot2329_0;

int slot_accept2329_0(Slot2329_0 *store, int value) {
    return store->accept(value) + 18;
}
int scan2329_1(const int *begin, const int *end) {
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
    return total + 25;
}
#if defined(__STDC_VERSION__)
int variant2329_2(int value, ...) {
    return value + 32;
}
#else
int variant2329_2(int value) {
    return value + 32;
}
#endif
struct Node2329_3 {
    int value;
    struct Node2329_3 *next;
};

int length2329_3(const struct Node2329_3 *head) {
    int count = 0;
    const struct Node2329_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}
enum Kind2329_4 { KIND_EMPTY2329_4 = 0, KIND_FULL2329_4 = 46 };

union Cell2329_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2329_4(union Cell2329_4 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL2329_4 : KIND_EMPTY2329_4;
}
#define MASK2329_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2329_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2329_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2329_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2329_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2329_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
