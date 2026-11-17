
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2037_0 {
    int value;
    struct Node2037_0 *next;
};

int length2037_0(const struct Node2037_0 *head) {
    int count = 0;
    const struct Node2037_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 5;
}
#if defined(__STDC_VERSION__)
int variant2037_1(int value, ...) {
    return value + 12;
}
#else
int variant2037_1(int value) {
    return value + 12;
}
#endif
enum Kind2037_2 { KIND_EMPTY2037_2 = 0, KIND_FULL2037_2 = 19 };

union Cell2037_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2037_2(union Cell2037_2 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL2037_2 : KIND_EMPTY2037_2;
}
enum Kind2037_3 { KIND_EMPTY2037_3 = 0, KIND_FULL2037_3 = 26 };

union Cell2037_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2037_3(union Cell2037_3 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL2037_3 : KIND_EMPTY2037_3;
}
#if defined(__STDC_VERSION__)
int variant2037_4(int value, ...) {
    return value + 33;
}
#else
int variant2037_4(int value) {
    return value + 33;
}
#endif
int gcd2037_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2037_5(int value) {
    return gcd2037_5(value, 49) * 40;
}
typedef int (*predicate2037_6)(int value);

typedef struct Slot2037_6 {
    int slots[16];
    predicate2037_6 accept;
} Slot2037_6;

int slot_accept2037_6(Slot2037_6 *store, int value) {
    return store->accept(value) + 47;
}
#define MASK2037_7(v) (((v) << 4) | ((v) & 0x0F))

int classify2037_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2037_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
