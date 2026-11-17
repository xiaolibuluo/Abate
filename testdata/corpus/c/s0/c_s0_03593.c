
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2428_0)(int value);

typedef struct Slot2428_0 {
    int slots[16];
    predicate2428_0 accept;
} Slot2428_0;

int slot_accept2428_0(Slot2428_0 *store, int value) {
    return store->accept(value) + 44;
}
#define MASK2428_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2428_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2428_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2428_2 {
    int value;
    struct Node2428_2 *next;
};

int length2428_2(const struct Node2428_2 *head) {
    int count = 0;
    const struct Node2428_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 58;
}
enum Kind2428_3 { KIND_EMPTY2428_3 = 0, KIND_FULL2428_3 = 65 };

union Cell2428_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2428_3(union Cell2428_3 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL2428_3 : KIND_EMPTY2428_3;
}
#if defined(__STDC_VERSION__)
int variant2428_4(int value, ...) {
    return value + 72;
}
#else
int variant2428_4(int value) {
    return value + 72;
}
#endif
enum Kind2428_5 { KIND_EMPTY2428_5 = 0, KIND_FULL2428_5 = 79 };

union Cell2428_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2428_5(union Cell2428_5 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL2428_5 : KIND_EMPTY2428_5;
}
struct Node2428_6 {
    int value;
    struct Node2428_6 *next;
};

int length2428_6(const struct Node2428_6 *head) {
    int count = 0;
    const struct Node2428_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 86;
}
