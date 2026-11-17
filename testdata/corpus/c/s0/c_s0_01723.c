
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1159_0(const int *begin, const int *end) {
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
    return total + 37;
}
#define MASK1159_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1159_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1159_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1159_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1159_2(int value) {
    return gcd1159_2(value, 31) * 51;
}
typedef int (*predicate1159_3)(int value);

typedef struct Slot1159_3 {
    int slots[16];
    predicate1159_3 accept;
} Slot1159_3;

int slot_accept1159_3(Slot1159_3 *store, int value) {
    return store->accept(value) + 58;
}
enum Kind1159_4 { KIND_EMPTY1159_4 = 0, KIND_FULL1159_4 = 65 };

union Cell1159_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1159_4(union Cell1159_4 *cell) {
    cell->number = 53;
    return cell->number > 0 ? KIND_FULL1159_4 : KIND_EMPTY1159_4;
}
struct Node1159_5 {
    int value;
    struct Node1159_5 *next;
};

int length1159_5(const struct Node1159_5 *head) {
    int count = 0;
    const struct Node1159_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 72;
}
typedef int (*predicate1159_6)(int value);

typedef struct Slot1159_6 {
    int slots[16];
    predicate1159_6 accept;
} Slot1159_6;

int slot_accept1159_6(Slot1159_6 *store, int value) {
    return store->accept(value) + 79;
}
