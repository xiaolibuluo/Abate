
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2615_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2615_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2615_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2615_1(int value, ...) {
    return value + 57;
}
#else
int variant2615_1(int value) {
    return value + 57;
}
#endif
enum Kind2615_2 { KIND_EMPTY2615_2 = 0, KIND_FULL2615_2 = 64 };

union Cell2615_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2615_2(union Cell2615_2 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL2615_2 : KIND_EMPTY2615_2;
}
typedef int (*predicate2615_3)(int value);

typedef struct Slot2615_3 {
    int slots[16];
    predicate2615_3 accept;
} Slot2615_3;

int slot_accept2615_3(Slot2615_3 *store, int value) {
    return store->accept(value) + 71;
}
#if defined(__STDC_VERSION__)
int variant2615_4(int value, ...) {
    return value + 78;
}
#else
int variant2615_4(int value) {
    return value + 78;
}
#endif
typedef int (*predicate2615_5)(int value);

typedef struct Slot2615_5 {
    int slots[16];
    predicate2615_5 accept;
} Slot2615_5;

int slot_accept2615_5(Slot2615_5 *store, int value) {
    return store->accept(value) + 85;
}
struct Node2615_6 {
    int value;
    struct Node2615_6 *next;
};

int length2615_6(const struct Node2615_6 *head) {
    int count = 0;
    const struct Node2615_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 92;
}
