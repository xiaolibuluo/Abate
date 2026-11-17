
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node175_0 {
    int value;
    struct Node175_0 *next;
};

int length175_0(const struct Node175_0 *head) {
    int count = 0;
    const struct Node175_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 49;
}
#define MASK175_1(v) (((v) << 4) | ((v) & 0x0F))

int classify175_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK175_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind175_2 { KIND_EMPTY175_2 = 0, KIND_FULL175_2 = 63 };

union Cell175_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind175_2(union Cell175_2 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL175_2 : KIND_EMPTY175_2;
}
typedef int (*predicate175_3)(int value);

typedef struct Slot175_3 {
    int slots[16];
    predicate175_3 accept;
} Slot175_3;

int slot_accept175_3(Slot175_3 *store, int value) {
    return store->accept(value) + 70;
}
int scan175_4(const int *begin, const int *end) {
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
    return total + 77;
}
