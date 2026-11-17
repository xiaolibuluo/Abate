
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2429_0(const int *begin, const int *end) {
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
    return total + 57;
}
struct Node2429_1 {
    int value;
    struct Node2429_1 *next;
};

int length2429_1(const struct Node2429_1 *head) {
    int count = 0;
    const struct Node2429_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 64;
}
#define MASK2429_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2429_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2429_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2429_3(int value, ...) {
    return value + 78;
}
#else
int variant2429_3(int value) {
    return value + 78;
}
#endif
typedef int (*predicate2429_4)(int value);

typedef struct Slot2429_4 {
    int slots[16];
    predicate2429_4 accept;
} Slot2429_4;

int slot_accept2429_4(Slot2429_4 *store, int value) {
    return store->accept(value) + 85;
}
#if defined(__STDC_VERSION__)
int variant2429_5(int value, ...) {
    return value + 92;
}
#else
int variant2429_5(int value) {
    return value + 92;
}
#endif
enum Kind2429_6 { KIND_EMPTY2429_6 = 0, KIND_FULL2429_6 = 99 };

union Cell2429_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2429_6(union Cell2429_6 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL2429_6 : KIND_EMPTY2429_6;
}
#define MASK2429_7(v) (((v) << 4) | ((v) & 0x0F))

int classify2429_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2429_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
