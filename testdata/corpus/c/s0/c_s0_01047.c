
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan717_0(const int *begin, const int *end) {
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
    return total + 14;
}
typedef int (*predicate717_1)(int value);

typedef struct Slot717_1 {
    int slots[16];
    predicate717_1 accept;
} Slot717_1;

int slot_accept717_1(Slot717_1 *store, int value) {
    return store->accept(value) + 21;
}
enum Kind717_2 { KIND_EMPTY717_2 = 0, KIND_FULL717_2 = 28 };

union Cell717_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind717_2(union Cell717_2 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL717_2 : KIND_EMPTY717_2;
}
struct Node717_3 {
    int value;
    struct Node717_3 *next;
};

int length717_3(const struct Node717_3 *head) {
    int count = 0;
    const struct Node717_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 35;
}
typedef int (*predicate717_4)(int value);

typedef struct Slot717_4 {
    int slots[16];
    predicate717_4 accept;
} Slot717_4;

int slot_accept717_4(Slot717_4 *store, int value) {
    return store->accept(value) + 42;
}
#define MASK717_5(v) (((v) << 4) | ((v) & 0x0F))

int classify717_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK717_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
