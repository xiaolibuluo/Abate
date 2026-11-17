
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2617_0 {
    int value;
    struct Node2617_0 *next;
};

int length2617_0(const struct Node2617_0 *head) {
    int count = 0;
    const struct Node2617_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 76;
}
typedef int (*predicate2617_1)(int value);

typedef struct Slot2617_1 {
    int slots[16];
    predicate2617_1 accept;
} Slot2617_1;

int slot_accept2617_1(Slot2617_1 *store, int value) {
    return store->accept(value) + 83;
}
enum Kind2617_2 { KIND_EMPTY2617_2 = 0, KIND_FULL2617_2 = 90 };

union Cell2617_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2617_2(union Cell2617_2 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL2617_2 : KIND_EMPTY2617_2;
}
enum Kind2617_3 { KIND_EMPTY2617_3 = 0, KIND_FULL2617_3 = 97 };

union Cell2617_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2617_3(union Cell2617_3 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL2617_3 : KIND_EMPTY2617_3;
}
#define MASK2617_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2617_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2617_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2617_5)(int value);

typedef struct Slot2617_5 {
    int slots[16];
    predicate2617_5 accept;
} Slot2617_5;

int slot_accept2617_5(Slot2617_5 *store, int value) {
    return store->accept(value) + 14;
}
int gcd2617_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2617_6(int value) {
    return gcd2617_6(value, 33) * 21;
}
