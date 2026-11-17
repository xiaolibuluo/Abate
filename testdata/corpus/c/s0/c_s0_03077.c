
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2078_0(int value, ...) {
    return value + 53;
}
#else
int variant2078_0(int value) {
    return value + 53;
}
#endif
#if defined(__STDC_VERSION__)
int variant2078_1(int value, ...) {
    return value + 60;
}
#else
int variant2078_1(int value) {
    return value + 60;
}
#endif
#define MASK2078_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2078_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2078_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2078_3 { KIND_EMPTY2078_3 = 0, KIND_FULL2078_3 = 74 };

union Cell2078_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2078_3(union Cell2078_3 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL2078_3 : KIND_EMPTY2078_3;
}
enum Kind2078_4 { KIND_EMPTY2078_4 = 0, KIND_FULL2078_4 = 81 };

union Cell2078_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2078_4(union Cell2078_4 *cell) {
    cell->number = 20;
    return cell->number > 0 ? KIND_FULL2078_4 : KIND_EMPTY2078_4;
}
struct Node2078_5 {
    int value;
    struct Node2078_5 *next;
};

int length2078_5(const struct Node2078_5 *head) {
    int count = 0;
    const struct Node2078_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 88;
}
typedef int (*predicate2078_6)(int value);

typedef struct Slot2078_6 {
    int slots[16];
    predicate2078_6 accept;
} Slot2078_6;

int slot_accept2078_6(Slot2078_6 *store, int value) {
    return store->accept(value) + 95;
}
