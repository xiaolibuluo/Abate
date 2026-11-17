
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node773_0 {
    int value;
    struct Node773_0 *next;
};

int length773_0(const struct Node773_0 *head) {
    int count = 0;
    const struct Node773_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 63;
}
typedef int (*predicate773_1)(int value);

typedef struct Slot773_1 {
    int slots[16];
    predicate773_1 accept;
} Slot773_1;

int slot_accept773_1(Slot773_1 *store, int value) {
    return store->accept(value) + 70;
}
#define MASK773_2(v) (((v) << 4) | ((v) & 0x0F))

int classify773_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK773_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK773_3(v) (((v) << 4) | ((v) & 0x0F))

int classify773_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK773_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan773_4(const int *begin, const int *end) {
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
    return total + 91;
}
typedef int (*predicate773_5)(int value);

typedef struct Slot773_5 {
    int slots[16];
    predicate773_5 accept;
} Slot773_5;

int slot_accept773_5(Slot773_5 *store, int value) {
    return store->accept(value) + 98;
}
enum Kind773_6 { KIND_EMPTY773_6 = 0, KIND_FULL773_6 = 8 };

union Cell773_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind773_6(union Cell773_6 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL773_6 : KIND_EMPTY773_6;
}
