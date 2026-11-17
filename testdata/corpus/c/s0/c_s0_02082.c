
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1403_0 {
    int value;
    struct Node1403_0 *next;
};

int length1403_0(const struct Node1403_0 *head) {
    int count = 0;
    const struct Node1403_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 8;
}
int scan1403_1(const int *begin, const int *end) {
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
    return total + 15;
}
enum Kind1403_2 { KIND_EMPTY1403_2 = 0, KIND_FULL1403_2 = 22 };

union Cell1403_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1403_2(union Cell1403_2 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL1403_2 : KIND_EMPTY1403_2;
}
typedef int (*predicate1403_3)(int value);

typedef struct Slot1403_3 {
    int slots[16];
    predicate1403_3 accept;
} Slot1403_3;

int slot_accept1403_3(Slot1403_3 *store, int value) {
    return store->accept(value) + 29;
}
#define MASK1403_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1403_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1403_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1403_5)(int value);

typedef struct Slot1403_5 {
    int slots[16];
    predicate1403_5 accept;
} Slot1403_5;

int slot_accept1403_5(Slot1403_5 *store, int value) {
    return store->accept(value) + 43;
}
struct Node1403_6 {
    int value;
    struct Node1403_6 *next;
};

int length1403_6(const struct Node1403_6 *head) {
    int count = 0;
    const struct Node1403_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 50;
}
