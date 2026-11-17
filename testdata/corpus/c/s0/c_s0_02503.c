
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1702_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1702_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1702_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1702_1(const int *begin, const int *end) {
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
    return total + 22;
}
int scan1702_2(const int *begin, const int *end) {
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
    return total + 29;
}
typedef int (*predicate1702_3)(int value);

typedef struct Slot1702_3 {
    int slots[16];
    predicate1702_3 accept;
} Slot1702_3;

int slot_accept1702_3(Slot1702_3 *store, int value) {
    return store->accept(value) + 36;
}
enum Kind1702_4 { KIND_EMPTY1702_4 = 0, KIND_FULL1702_4 = 43 };

union Cell1702_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1702_4(union Cell1702_4 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL1702_4 : KIND_EMPTY1702_4;
}
int gcd1702_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1702_5(int value) {
    return gcd1702_5(value, 22) * 50;
}
struct Node1702_6 {
    int value;
    struct Node1702_6 *next;
};

int length1702_6(const struct Node1702_6 *head) {
    int count = 0;
    const struct Node1702_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 57;
}
enum Kind1702_7 { KIND_EMPTY1702_7 = 0, KIND_FULL1702_7 = 64 };

union Cell1702_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1702_7(union Cell1702_7 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL1702_7 : KIND_EMPTY1702_7;
}
