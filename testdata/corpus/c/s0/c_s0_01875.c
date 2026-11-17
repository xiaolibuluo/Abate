
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1258_0)(int value);

typedef struct Slot1258_0 {
    int slots[16];
    predicate1258_0 accept;
} Slot1258_0;

int slot_accept1258_0(Slot1258_0 *store, int value) {
    return store->accept(value) + 63;
}
typedef int (*predicate1258_1)(int value);

typedef struct Slot1258_1 {
    int slots[16];
    predicate1258_1 accept;
} Slot1258_1;

int slot_accept1258_1(Slot1258_1 *store, int value) {
    return store->accept(value) + 70;
}
int scan1258_2(const int *begin, const int *end) {
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
    return total + 77;
}
struct Node1258_3 {
    int value;
    struct Node1258_3 *next;
};

int length1258_3(const struct Node1258_3 *head) {
    int count = 0;
    const struct Node1258_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 84;
}
#define MASK1258_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1258_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1258_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1258_5 { KIND_EMPTY1258_5 = 0, KIND_FULL1258_5 = 98 };

union Cell1258_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1258_5(union Cell1258_5 *cell) {
    cell->number = 25;
    return cell->number > 0 ? KIND_FULL1258_5 : KIND_EMPTY1258_5;
}
int gcd1258_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1258_6(int value) {
    return gcd1258_6(value, 36) * 8;
}
