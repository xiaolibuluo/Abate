
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1009_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1009_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1009_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1009_1)(int value);

typedef struct Slot1009_1 {
    int slots[16];
    predicate1009_1 accept;
} Slot1009_1;

int slot_accept1009_1(Slot1009_1 *store, int value) {
    return store->accept(value) + 34;
}
struct Node1009_2 {
    int value;
    struct Node1009_2 *next;
};

int length1009_2(const struct Node1009_2 *head) {
    int count = 0;
    const struct Node1009_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 41;
}
#define MASK1009_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1009_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1009_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1009_4(const int *begin, const int *end) {
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
    return total + 55;
}
enum Kind1009_5 { KIND_EMPTY1009_5 = 0, KIND_FULL1009_5 = 62 };

union Cell1009_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1009_5(union Cell1009_5 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL1009_5 : KIND_EMPTY1009_5;
}
