
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1507_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1507_0(int value) {
    return gcd1507_0(value, 66) * 99;
}
struct Node1507_1 {
    int value;
    struct Node1507_1 *next;
};

int length1507_1(const struct Node1507_1 *head) {
    int count = 0;
    const struct Node1507_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 9;
}
enum Kind1507_2 { KIND_EMPTY1507_2 = 0, KIND_FULL1507_2 = 16 };

union Cell1507_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1507_2(union Cell1507_2 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL1507_2 : KIND_EMPTY1507_2;
}
#define MASK1507_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1507_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1507_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1507_4(const int *begin, const int *end) {
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
    return total + 30;
}
typedef int (*predicate1507_5)(int value);

typedef struct Slot1507_5 {
    int slots[16];
    predicate1507_5 accept;
} Slot1507_5;

int slot_accept1507_5(Slot1507_5 *store, int value) {
    return store->accept(value) + 37;
}
enum Kind1507_6 { KIND_EMPTY1507_6 = 0, KIND_FULL1507_6 = 44 };

union Cell1507_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1507_6(union Cell1507_6 *cell) {
    cell->number = 10;
    return cell->number > 0 ? KIND_FULL1507_6 : KIND_EMPTY1507_6;
}
struct Node1507_7 {
    int value;
    struct Node1507_7 *next;
};

int length1507_7(const struct Node1507_7 *head) {
    int count = 0;
    const struct Node1507_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 51;
}
