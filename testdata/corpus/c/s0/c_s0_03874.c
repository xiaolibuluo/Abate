
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2620_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2620_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2620_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2620_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2620_1(int value) {
    return gcd2620_1(value, 60) * 25;
}
typedef int (*predicate2620_2)(int value);

typedef struct Slot2620_2 {
    int slots[16];
    predicate2620_2 accept;
} Slot2620_2;

int slot_accept2620_2(Slot2620_2 *store, int value) {
    return store->accept(value) + 32;
}
struct Node2620_3 {
    int value;
    struct Node2620_3 *next;
};

int length2620_3(const struct Node2620_3 *head) {
    int count = 0;
    const struct Node2620_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}
typedef int (*predicate2620_4)(int value);

typedef struct Slot2620_4 {
    int slots[16];
    predicate2620_4 accept;
} Slot2620_4;

int slot_accept2620_4(Slot2620_4 *store, int value) {
    return store->accept(value) + 46;
}
enum Kind2620_5 { KIND_EMPTY2620_5 = 0, KIND_FULL2620_5 = 53 };

union Cell2620_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2620_5(union Cell2620_5 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL2620_5 : KIND_EMPTY2620_5;
}
int gcd2620_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2620_6(int value) {
    return gcd2620_6(value, 54) * 60;
}
