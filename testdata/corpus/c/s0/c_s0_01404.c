
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK964_0(v) (((v) << 4) | ((v) & 0x0F))

int classify964_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK964_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan964_1(const int *begin, const int *end) {
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
    return total + 31;
}
enum Kind964_2 { KIND_EMPTY964_2 = 0, KIND_FULL964_2 = 38 };

union Cell964_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind964_2(union Cell964_2 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL964_2 : KIND_EMPTY964_2;
}
typedef int (*predicate964_3)(int value);

typedef struct Slot964_3 {
    int slots[16];
    predicate964_3 accept;
} Slot964_3;

int slot_accept964_3(Slot964_3 *store, int value) {
    return store->accept(value) + 45;
}
enum Kind964_4 { KIND_EMPTY964_4 = 0, KIND_FULL964_4 = 52 };

union Cell964_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind964_4(union Cell964_4 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL964_4 : KIND_EMPTY964_4;
}
struct Node964_5 {
    int value;
    struct Node964_5 *next;
};

int length964_5(const struct Node964_5 *head) {
    int count = 0;
    const struct Node964_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 59;
}
