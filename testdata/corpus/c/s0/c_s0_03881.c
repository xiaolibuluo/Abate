
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2622_0(const int *begin, const int *end) {
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
    return total + 44;
}
struct Node2622_1 {
    int value;
    struct Node2622_1 *next;
};

int length2622_1(const struct Node2622_1 *head) {
    int count = 0;
    const struct Node2622_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 51;
}
int gcd2622_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2622_2(int value) {
    return gcd2622_2(value, 24) * 58;
}
enum Kind2622_3 { KIND_EMPTY2622_3 = 0, KIND_FULL2622_3 = 65 };

union Cell2622_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2622_3(union Cell2622_3 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL2622_3 : KIND_EMPTY2622_3;
}
typedef int (*predicate2622_4)(int value);

typedef struct Slot2622_4 {
    int slots[16];
    predicate2622_4 accept;
} Slot2622_4;

int slot_accept2622_4(Slot2622_4 *store, int value) {
    return store->accept(value) + 72;
}
#define MASK2622_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2622_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2622_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2622_6)(int value);

typedef struct Slot2622_6 {
    int slots[16];
    predicate2622_6 accept;
} Slot2622_6;

int slot_accept2622_6(Slot2622_6 *store, int value) {
    return store->accept(value) + 86;
}
