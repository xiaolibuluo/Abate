
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1355_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1355_0(int value) {
    return gcd1355_0(value, 39) * 63;
}
#define MASK1355_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1355_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1355_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1355_2 {
    int value;
    struct Node1355_2 *next;
};

int length1355_2(const struct Node1355_2 *head) {
    int count = 0;
    const struct Node1355_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 77;
}
int scan1355_3(const int *begin, const int *end) {
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
    return total + 84;
}
int gcd1355_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1355_4(int value) {
    return gcd1355_4(value, 22) * 91;
}
enum Kind1355_5 { KIND_EMPTY1355_5 = 0, KIND_FULL1355_5 = 98 };

union Cell1355_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1355_5(union Cell1355_5 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL1355_5 : KIND_EMPTY1355_5;
}
typedef int (*predicate1355_6)(int value);

typedef struct Slot1355_6 {
    int slots[16];
    predicate1355_6 accept;
} Slot1355_6;

int slot_accept1355_6(Slot1355_6 *store, int value) {
    return store->accept(value) + 8;
}
