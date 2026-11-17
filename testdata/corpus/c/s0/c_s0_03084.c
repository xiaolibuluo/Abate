
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2085_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2085_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2085_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2085_1)(int value);

typedef struct Slot2085_1 {
    int slots[16];
    predicate2085_1 accept;
} Slot2085_1;

int slot_accept2085_1(Slot2085_1 *store, int value) {
    return store->accept(value) + 54;
}
int scan2085_2(const int *begin, const int *end) {
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
    return total + 61;
}
struct Node2085_3 {
    int value;
    struct Node2085_3 *next;
};

int length2085_3(const struct Node2085_3 *head) {
    int count = 0;
    const struct Node2085_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
enum Kind2085_4 { KIND_EMPTY2085_4 = 0, KIND_FULL2085_4 = 75 };

union Cell2085_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2085_4(union Cell2085_4 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL2085_4 : KIND_EMPTY2085_4;
}
#define MASK2085_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2085_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2085_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2085_6)(int value);

typedef struct Slot2085_6 {
    int slots[16];
    predicate2085_6 accept;
} Slot2085_6;

int slot_accept2085_6(Slot2085_6 *store, int value) {
    return store->accept(value) + 89;
}
