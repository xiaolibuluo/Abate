
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2618_0 {
    int value;
    struct Node2618_0 *next;
};

int length2618_0(const struct Node2618_0 *head) {
    int count = 0;
    const struct Node2618_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 89;
}
typedef int (*predicate2618_1)(int value);

typedef struct Slot2618_1 {
    int slots[16];
    predicate2618_1 accept;
} Slot2618_1;

int slot_accept2618_1(Slot2618_1 *store, int value) {
    return store->accept(value) + 96;
}
#if defined(__STDC_VERSION__)
int variant2618_2(int value, ...) {
    return value + 6;
}
#else
int variant2618_2(int value) {
    return value + 6;
}
#endif
#define MASK2618_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2618_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2618_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2618_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2618_4(int value) {
    return gcd2618_4(value, 18) * 20;
}
enum Kind2618_5 { KIND_EMPTY2618_5 = 0, KIND_FULL2618_5 = 27 };

union Cell2618_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2618_5(union Cell2618_5 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL2618_5 : KIND_EMPTY2618_5;
}
struct Node2618_6 {
    int value;
    struct Node2618_6 *next;
};

int length2618_6(const struct Node2618_6 *head) {
    int count = 0;
    const struct Node2618_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 34;
}
