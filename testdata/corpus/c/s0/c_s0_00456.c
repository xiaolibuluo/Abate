
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate311_0)(int value);

typedef struct Slot311_0 {
    int slots[16];
    predicate311_0 accept;
} Slot311_0;

int slot_accept311_0(Slot311_0 *store, int value) {
    return store->accept(value) + 71;
}
#define MASK311_1(v) (((v) << 4) | ((v) & 0x0F))

int classify311_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK311_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node311_2 {
    int value;
    struct Node311_2 *next;
};

int length311_2(const struct Node311_2 *head) {
    int count = 0;
    const struct Node311_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 85;
}
int scan311_3(const int *begin, const int *end) {
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
    return total + 92;
}
