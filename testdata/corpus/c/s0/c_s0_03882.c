
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2623_0 {
    int value;
    struct Node2623_0 *next;
};

int length2623_0(const struct Node2623_0 *head) {
    int count = 0;
    const struct Node2623_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 57;
}
#define MASK2623_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2623_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2623_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2623_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2623_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2623_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2623_3 { KIND_EMPTY2623_3 = 0, KIND_FULL2623_3 = 78 };

union Cell2623_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2623_3(union Cell2623_3 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL2623_3 : KIND_EMPTY2623_3;
}
typedef int (*predicate2623_4)(int value);

typedef struct Slot2623_4 {
    int slots[16];
    predicate2623_4 accept;
} Slot2623_4;

int slot_accept2623_4(Slot2623_4 *store, int value) {
    return store->accept(value) + 85;
}
typedef int (*predicate2623_5)(int value);

typedef struct Slot2623_5 {
    int slots[16];
    predicate2623_5 accept;
} Slot2623_5;

int slot_accept2623_5(Slot2623_5 *store, int value) {
    return store->accept(value) + 92;
}
int gcd2623_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2623_6(int value) {
    return gcd2623_6(value, 14) * 99;
}
enum Kind2623_7 { KIND_EMPTY2623_7 = 0, KIND_FULL2623_7 = 9 };

union Cell2623_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2623_7(union Cell2623_7 *cell) {
    cell->number = 25;
    return cell->number > 0 ? KIND_FULL2623_7 : KIND_EMPTY2623_7;
}
