
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2169_0(int value, ...) {
    return value + 72;
}
#else
int variant2169_0(int value) {
    return value + 72;
}
#endif
#if defined(__STDC_VERSION__)
int variant2169_1(int value, ...) {
    return value + 79;
}
#else
int variant2169_1(int value) {
    return value + 79;
}
#endif
struct Node2169_2 {
    int value;
    struct Node2169_2 *next;
};

int length2169_2(const struct Node2169_2 *head) {
    int count = 0;
    const struct Node2169_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 86;
}
typedef int (*predicate2169_3)(int value);

typedef struct Slot2169_3 {
    int slots[16];
    predicate2169_3 accept;
} Slot2169_3;

int slot_accept2169_3(Slot2169_3 *store, int value) {
    return store->accept(value) + 93;
}
#define MASK2169_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2169_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2169_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2169_5 { KIND_EMPTY2169_5 = 0, KIND_FULL2169_5 = 10 };

union Cell2169_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2169_5(union Cell2169_5 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL2169_5 : KIND_EMPTY2169_5;
}
