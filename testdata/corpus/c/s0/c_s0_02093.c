
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1410_0 {
    int value;
    struct Node1410_0 *next;
};

int length1410_0(const struct Node1410_0 *head) {
    int count = 0;
    const struct Node1410_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 99;
}
struct Node1410_1 {
    int value;
    struct Node1410_1 *next;
};

int length1410_1(const struct Node1410_1 *head) {
    int count = 0;
    const struct Node1410_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 9;
}
struct Node1410_2 {
    int value;
    struct Node1410_2 *next;
};

int length1410_2(const struct Node1410_2 *head) {
    int count = 0;
    const struct Node1410_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 16;
}
#define MASK1410_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1410_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1410_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1410_4 { KIND_EMPTY1410_4 = 0, KIND_FULL1410_4 = 30 };

union Cell1410_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1410_4(union Cell1410_4 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL1410_4 : KIND_EMPTY1410_4;
}
typedef int (*predicate1410_5)(int value);

typedef struct Slot1410_5 {
    int slots[16];
    predicate1410_5 accept;
} Slot1410_5;

int slot_accept1410_5(Slot1410_5 *store, int value) {
    return store->accept(value) + 37;
}
int scan1410_6(const int *begin, const int *end) {
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
