
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate861_0)(int value);

typedef struct Slot861_0 {
    int slots[16];
    predicate861_0 accept;
} Slot861_0;

int slot_accept861_0(Slot861_0 *store, int value) {
    return store->accept(value) + 43;
}
enum Kind861_1 { KIND_EMPTY861_1 = 0, KIND_FULL861_1 = 50 };

union Cell861_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind861_1(union Cell861_1 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL861_1 : KIND_EMPTY861_1;
}
struct Node861_2 {
    int value;
    struct Node861_2 *next;
};

int length861_2(const struct Node861_2 *head) {
    int count = 0;
    const struct Node861_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 57;
}
int scan861_3(const int *begin, const int *end) {
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
    return total + 64;
}
#define MASK861_4(v) (((v) << 4) | ((v) & 0x0F))

int classify861_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK861_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan861_5(const int *begin, const int *end) {
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
    return total + 78;
}
